package handlers

import (
	"net/http"
	"strconv"
	"time"

	"koinonia-backend/models"
)

// Submission Handlers

// GetSubmissions returns all submissions (admin only)
func (h *Handler) GetSubmissions(w http.ResponseWriter, r *http.Request) {
	// Query parameters for filtering
	status := r.URL.Query().Get("status")
	questID := r.URL.Query().Get("quest_id")
	userID := r.URL.Query().Get("user_id")

	query := h.db.Preload("User").Preload("Quest").Preload("ReviewedBy")

	// Apply filters if provided
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if questID != "" {
		if id, err := strconv.ParseUint(questID, 10, 32); err == nil {
			query = query.Where("quest_id = ?", uint(id))
		}
	}
	if userID != "" {
		if id, err := strconv.ParseUint(userID, 10, 32); err == nil {
			query = query.Where("user_id = ?", uint(id))
		}
	}

	var submissions []models.Submission
	if err := query.Order("created_at DESC").Find(&submissions).Error; err != nil {
		writeJSONError(w, "Failed to fetch submissions", http.StatusInternalServerError)
		return
	}

	writeJSON(w, submissions, http.StatusOK)
}

// ApproveSubmission approves a quest submission and awards points
func (h *Handler) ApproveSubmission(w http.ResponseWriter, r *http.Request) {
	submissionID, err := parseID(r, "id")
	if err != nil {
		writeJSONError(w, "Invalid submission ID", http.StatusBadRequest)
		return
	}

	adminID := r.Context().Value("user_id").(uint)

	// Get submission with related quest
	var submission models.Submission
	if err := h.db.Preload("Quest").Preload("User").First(&submission, submissionID).Error; err != nil {
		writeJSONError(w, "Submission not found", http.StatusNotFound)
		return
	}

	// Check if already reviewed
	if submission.Status != models.SubmissionStatusPending {
		writeJSONError(w, "Submission already reviewed", http.StatusBadRequest)
		return
	}

	// Start database transaction
	tx := h.db.Begin()

	// Update submission status
	now := time.Now()
	updates := map[string]interface{}{
		"status":          models.SubmissionStatusApproved,
		"points_awarded":  submission.Quest.Points,
		"reviewed_at":     &now,
		"reviewed_by_id":  adminID,
	}

	if err := tx.Model(&submission).Updates(updates).Error; err != nil {
		tx.Rollback()
		writeJSONError(w, "Failed to update submission", http.StatusInternalServerError)
		return
	}

	// Award points to user
	if err := tx.Model(&models.User{}).Where("id = ?", submission.UserID).
		UpdateColumn("total_points", h.db.Raw("total_points + ?", submission.Quest.Points)).Error; err != nil {
		tx.Rollback()
		writeJSONError(w, "Failed to award points", http.StatusInternalServerError)
		return
	}

	// Commit transaction
	tx.Commit()

	// Return updated submission
	h.db.Preload("User").Preload("Quest").Preload("ReviewedBy").First(&submission, submissionID)
	writeJSON(w, submission, http.StatusOK)
}

// RejectSubmission rejects a quest submission
func (h *Handler) RejectSubmission(w http.ResponseWriter, r *http.Request) {
	submissionID, err := parseID(r, "id")
	if err != nil {
		writeJSONError(w, "Invalid submission ID", http.StatusBadRequest)
		return
	}

	adminID := r.Context().Value("user_id").(uint)

	// Parse admin notes from request body
	var req struct {
		AdminNotes string `json:"admin_notes"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	// Get submission
	var submission models.Submission
	if err := h.db.First(&submission, submissionID).Error; err != nil {
		writeJSONError(w, "Submission not found", http.StatusNotFound)
		return
	}

	// Check if already reviewed
	if submission.Status != models.SubmissionStatusPending {
		writeJSONError(w, "Submission already reviewed", http.StatusBadRequest)
		return
	}

	// Update submission status
	now := time.Now()
	updates := map[string]interface{}{
		"status":         models.SubmissionStatusRejected,
		"admin_notes":    req.AdminNotes,
		"reviewed_at":    &now,
		"reviewed_by_id": adminID,
	}

	if err := h.db.Model(&submission).Updates(updates).Error; err != nil {
		writeJSONError(w, "Failed to update submission", http.StatusInternalServerError)
		return
	}

	// Return updated submission
	h.db.Preload("User").Preload("Quest").Preload("ReviewedBy").First(&submission, submissionID)
	writeJSON(w, submission, http.StatusOK)
}
