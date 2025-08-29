package handlers

import (
	"net/http"

	"koinonia-backend/models"
)

// GetLeaderboard returns the top users ranked by total points
func (h *Handler) GetLeaderboard(w http.ResponseWriter, r *http.Request) {
	// Query parameters
	limit := 10 // Default limit
	if limitParam := r.URL.Query().Get("limit"); limitParam != "" {
		if parsedLimit, err := parseID(r, "limit"); err == nil && parsedLimit > 0 {
			limit = int(parsedLimit)
		}
	}

	// Get users with their quest completion counts
	var leaderboard []models.LeaderboardEntry

	// Complex query to get users with quest completion counts
	query := `
		SELECT 
			u.id as user_id,
			u.username,
			u.first_name,
			u.last_name,
			u.avatar,
			u.total_points,
			COALESCE(s.quests_completed, 0) as quests_completed,
			ROW_NUMBER() OVER (ORDER BY u.total_points DESC, u.created_at ASC) as rank
		FROM users u
		LEFT JOIN (
			SELECT 
				user_id, 
				COUNT(*) as quests_completed
			FROM submissions 
			WHERE status = 'approved'
			GROUP BY user_id
		) s ON u.id = s.user_id
		WHERE u.is_active = true
		ORDER BY u.total_points DESC, u.created_at ASC
		LIMIT ?
	`

	if err := h.db.Raw(query, limit).Scan(&leaderboard).Error; err != nil {
		writeJSONError(w, "Failed to fetch leaderboard", http.StatusInternalServerError)
		return
	}

	writeJSON(w, leaderboard, http.StatusOK)
}
