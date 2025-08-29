package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"koinonia-backend/handlers"
	"koinonia-backend/models"
)

func main() {
	// Database connection
	db, err := connectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate database tables
	err = db.AutoMigrate(&models.User{}, &models.Quest{}, &models.Submission{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Initialize handlers with database
	h := handlers.New(db)

	// Setup router
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Next.js dev server
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Routes
	r.Route("/api", func(r chi.Router) {
		// Auth routes
		r.Post("/auth/register", h.Register)
		r.Post("/auth/login", h.Login)

		// Protected routes (require authentication)
		r.Group(func(r chi.Router) {
			r.Use(h.AuthMiddleware) // JWT authentication middleware

			// User routes
			r.Get("/profile", h.GetProfile)
			r.Put("/profile", h.UpdateProfile)

			// Quest routes
			r.Get("/quests", h.GetQuests)
			r.Get("/quests/{id}", h.GetQuest)
			r.Post("/quests/{id}/submit", h.SubmitQuest)

			// Leaderboard
			r.Get("/leaderboard", h.GetLeaderboard)

			// Admin routes (require admin role)
			r.Group(func(r chi.Router) {
				r.Use(h.AdminMiddleware)
				r.Post("/quests", h.CreateQuest)
				r.Put("/quests/{id}", h.UpdateQuest)
				r.Delete("/quests/{id}", h.DeleteQuest)
				r.Get("/submissions", h.GetSubmissions)
				r.Put("/submissions/{id}/approve", h.ApproveSubmission)
				r.Put("/submissions/{id}/reject", h.RejectSubmission)
			})
		})
	})

	// Health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

// connectDB establishes connection to PostgreSQL database
func connectDB() (*gorm.DB, error) {
	// Database configuration from environment variables
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "password")
	dbname := getEnv("DB_NAME", "koinonia")

	// Connection string
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Connect to database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// getEnv gets environment variable with fallback
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
