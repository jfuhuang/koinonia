package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"koinonia-backend/models"
)

// Handler holds database connection and provides HTTP handlers
type Handler struct {
	db *gorm.DB
}

// New creates a new handler instance
func New(db *gorm.DB) *Handler {
	return &Handler{db: db}
}

// JWT secret key - in production, this should be in environment variables
var jwtSecret = []byte("your-secret-key-change-this-in-production")

// Claims represents JWT claims
type Claims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// Response structures
type AuthResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

// Auth Handlers

// Register creates a new user account
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username  string `json:"username"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	// Parse JSON request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if req.Username == "" || req.Email == "" || req.Password == "" {
		writeJSONError(w, "Username, email, and password are required", http.StatusBadRequest)
		return
	}

	// Check if user already exists
	var existingUser models.User
	if err := h.db.Where("username = ? OR email = ?", req.Username, req.Email).First(&existingUser).Error; err == nil {
		writeJSONError(w, "Username or email already exists", http.StatusConflict)
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		writeJSONError(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	// Create user
	user := models.User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  string(hashedPassword),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Role:      "user", // Default role
		IsActive:  true,
	}

	if err := h.db.Create(&user).Error; err != nil {
		writeJSONError(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Generate JWT token
	token, err := h.generateJWT(user.ID, user.Role)
	if err != nil {
		writeJSONError(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Return response
	response := AuthResponse{
		Token: token,
		User:  user,
	}
	writeJSON(w, response, http.StatusCreated)
}

// Login authenticates a user and returns a JWT token
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"` // Can be username or email
		Password string `json:"password"`
	}

	// Parse JSON request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if req.Username == "" || req.Password == "" {
		writeJSONError(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	// Find user by username or email
	var user models.User
	if err := h.db.Where("username = ? OR email = ?", req.Username, req.Username).First(&user).Error; err != nil {
		writeJSONError(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Check if user is active
	if !user.IsActive {
		writeJSONError(w, "Account is deactivated", http.StatusUnauthorized)
		return
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		writeJSONError(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Update last login
	now := time.Now()
	user.LastLogin = &now
	h.db.Save(&user)

	// Generate JWT token
	token, err := h.generateJWT(user.ID, user.Role)
	if err != nil {
		writeJSONError(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Return response
	response := AuthResponse{
		Token: token,
		User:  user,
	}
	writeJSON(w, response, http.StatusOK)
}

// User Profile Handlers

// GetProfile returns the current user's profile
func (h *Handler) GetProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)

	var user models.User
	if err := h.db.First(&user, userID).Error; err != nil {
		writeJSONError(w, "User not found", http.StatusNotFound)
		return
	}

	writeJSON(w, user, http.StatusOK)
}

// UpdateProfile updates the current user's profile
func (h *Handler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)

	var req struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Bio       string `json:"bio"`
		Avatar    string `json:"avatar"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Update user profile
	updates := map[string]interface{}{
		"first_name": req.FirstName,
		"last_name":  req.LastName,
		"bio":        req.Bio,
		"avatar":     req.Avatar,
	}

	if err := h.db.Model(&models.User{}).Where("id = ?", userID).Updates(updates).Error; err != nil {
		writeJSONError(w, "Failed to update profile", http.StatusInternalServerError)
		return
	}

	// Return updated user
	var user models.User
	h.db.First(&user, userID)
	writeJSON(w, user, http.StatusOK)
}

// Helper functions

// generateJWT creates a JWT token for a user
func (h *Handler) generateJWT(userID uint, role string) (string, error) {
	claims := Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 24 hours
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// writeJSON writes a JSON response
func writeJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// writeJSONError writes a JSON error response
func writeJSONError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}

// parseID parses an ID from URL parameter
func parseID(r *http.Request, param string) (uint, error) {
	idStr := chi.URLParam(r, param)
	id, err := strconv.ParseUint(idStr, 10, 32)
	return uint(id), err
}
