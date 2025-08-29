package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"koinonia-backend/models"
)

// Quest Handlers

// GetQuests returns all active quests
func (h *Handler) GetQuests(w http.ResponseWriter, r *http.Request) {
	// Query parameters for filtering
	questType := r.URL.Query().Get("type")
	difficulty := r.URL.Query().Get("difficulty")

	query := h.db.Where("is_active = ?", true)

	// Apply filters if provided
	if questType != "" {
		query = query.Where("type = ?", questType)
	}
	if difficulty != "" {
		query = query.Where("difficulty = ?", difficulty)
	}

	// Get quests ordered by creation date
	var quests []models.Quest
	if err := query.Order("created_at DESC").Find(&quests).Error; err != nil {
		writeJSONError(w, "Failed to fetch quests", http.StatusInternalServerError)
		return
	}

	writeJSON(w, quests, http.StatusOK)
}

// GetQuest returns a specific quest by ID
func (h *Handler) GetQuest(w http.ResponseWriter, r *http.Request) {
	questID, err := parseID(r, "id")
	if err != nil {
		writeJSONError(w, "Invalid quest ID", http.StatusBadRequest)
		return
	}

	var quest models.Quest
	if err := h.db.Where("id = ? AND is_active = ?", questID, true).First(&quest).Error; err != nil {
		writeJSONError(w, "Quest not found", http.StatusNotFound)
		return
	}

	writeJSON(w, quest, http.StatusOK)
}

// SubmitQuest allows a user to submit a quest completion
func (h *Handler) SubmitQuest(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)
	questID, err := parseID(r, "id")
	if err != nil {
		writeJSONError(w, "Invalid quest ID", http.StatusBadRequest)
		return
	}

	var req struct {
		Content   string `json:"content"`    // Text response/answer
		MediaURL  string `json:"media_url"`  // URL to uploaded media
		MediaType string `json:"media_type"` // "image", "video", "audio"
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Verify quest exists and is active
	var quest models.Quest
	if err := h.db.Where("id = ? AND is_active = ?", questID, true).First(&quest).Error; err != nil {
		writeJSONError(w, "Quest not found or inactive", http.StatusNotFound)
		return
	}

	// Check if user has already submitted this quest (if applicable)
	var existingSubmission models.Submission
	if err := h.db.Where("user_id = ? AND quest_id = ?", userID, questID).First(&existingSubmission).Error; err == nil {
		// User already submitted - check if quest allows multiple submissions
		if quest.MaxSubmissions == 1 {
			writeJSONError(w, "You have already submitted this quest", http.StatusConflict)
			return
		}
	}

	// Create submission
	submission := models.Submission{
		UserID:    userID,
		QuestID:   questID,
		Content:   req.Content,
		MediaURL:  req.MediaURL,
		MediaType: req.MediaType,
		Status:    models.SubmissionStatusPending,
	}

	if err := h.db.Create(&submission).Error; err != nil {
		writeJSONError(w, "Failed to create submission", http.StatusInternalServerError)
		return
	}

	// Load related data for response
	h.db.Preload("Quest").Preload("User").First(&submission, submission.ID)

	writeJSON(w, submission, http.StatusCreated)
}

// Admin Quest Handlers

// CreateQuest allows admins to create new quests
func (h *Handler) CreateQuest(w http.ResponseWriter, r *http.Request) {
	var req models.Quest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if req.Title == "" || req.Type == "" || req.Points <= 0 {
		writeJSONError(w, "Title, type, and points are required", http.StatusBadRequest)
		return
	}

	// Create quest
	if err := h.db.Create(&req).Error; err != nil {
		writeJSONError(w, "Failed to create quest", http.StatusInternalServerError)
		return
	}

	writeJSON(w, req, http.StatusCreated)
}

// UpdateQuest allows admins to update existing quests
func (h *Handler) UpdateQuest(w http.ResponseWriter, r *http.Request) {
	questID, err := parseID(r, "id")
	if err != nil {
		writeJSONError(w, "Invalid quest ID", http.StatusBadRequest)
		return
	}

	var req models.Quest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Update quest
	if err := h.db.Model(&models.Quest{}).Where("id = ?", questID).Updates(&req).Error; err != nil {
		writeJSONError(w, "Failed to update quest", http.StatusInternalServerError)
		return
	}

	// Return updated quest
	var quest models.Quest
	h.db.First(&quest, questID)
	writeJSON(w, quest, http.StatusOK)
}

// DeleteQuest allows admins to delete quests (soft delete)
func (h *Handler) DeleteQuest(w http.ResponseWriter, r *http.Request) {
	questID, err := parseID(r, "id")
	if err != nil {
		writeJSONError(w, "Invalid quest ID", http.StatusBadRequest)
		return
	}

	// Soft delete the quest
	if err := h.db.Delete(&models.Quest{}, questID).Error; err != nil {
		writeJSONError(w, "Failed to delete quest", http.StatusInternalServerError)
		return
	}

	writeJSON(w, MessageResponse{Message: "Quest deleted successfully"}, http.StatusOK)
}
