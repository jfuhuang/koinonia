package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// User information
	Username    string `json:"username" gorm:"uniqueIndex;not null"`
	Email       string `json:"email" gorm:"uniqueIndex;not null"`
	Password    string `json:"-" gorm:"not null"` // Never include in JSON responses
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Avatar      string `json:"avatar"`       // URL to profile picture
	Bio         string `json:"bio"`          // Short biography
	TotalPoints int    `json:"total_points"` // Accumulated points from quests

	// User role and status
	Role      string `json:"role" gorm:"default:user"`        // "user" or "admin"
	IsActive  bool   `json:"is_active" gorm:"default:true"`   // Account status
	LastLogin *time.Time `json:"last_login"`                  // Track last login

	// Relationships
	Submissions []Submission `json:"submissions,omitempty" gorm:"foreignKey:UserID"`
}

// QuestType represents different types of quests
type QuestType string

const (
	QuestTypeScripture   QuestType = "scripture"   // Scripture memory quests
	QuestTypeSideQuest   QuestType = "side_quest"  // Photo-based campus challenges
	QuestTypeTrivia      QuestType = "trivia"      // Bible trivia questions
	QuestTypeEncouragement QuestType = "encouragement" // Encouraging others
)

// Quest represents a quest/challenge that users can complete
type Quest struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// Quest information
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description" gorm:"type:text"`
	Type        QuestType `json:"type" gorm:"not null"`
	Points      int       `json:"points" gorm:"not null"` // Points awarded for completion
	Difficulty  string    `json:"difficulty"`             // "easy", "medium", "hard"

	// Quest content (varies by type)
	ScriptureReference string `json:"scripture_reference,omitempty"` // For scripture quests
	ScriptureText      string `json:"scripture_text,omitempty" gorm:"type:text"` // The verse to memorize
	TriviaQuestion     string `json:"trivia_question,omitempty" gorm:"type:text"` // For trivia quests
	TriviaOptions      string `json:"trivia_options,omitempty" gorm:"type:text"`  // JSON array of options
	CorrectAnswer      string `json:"-"`                            // Hidden from JSON responses
	
	// Quest status and metadata
	IsActive    bool      `json:"is_active" gorm:"default:true"`
	StartDate   *time.Time `json:"start_date"`   // When quest becomes available
	EndDate     *time.Time `json:"end_date"`     // When quest expires
	MaxSubmissions int    `json:"max_submissions"` // 0 = unlimited submissions

	// Relationships
	Submissions []Submission `json:"submissions,omitempty" gorm:"foreignKey:QuestID"`
}

// SubmissionStatus represents the status of a quest submission
type SubmissionStatus string

const (
	SubmissionStatusPending  SubmissionStatus = "pending"
	SubmissionStatusApproved SubmissionStatus = "approved"
	SubmissionStatusRejected SubmissionStatus = "rejected"
)

// Submission represents a user's submission for a quest
type Submission struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// Foreign keys
	UserID  uint `json:"user_id" gorm:"not null"`
	QuestID uint `json:"quest_id" gorm:"not null"`

	// Submission content
	Content     string `json:"content" gorm:"type:text"`      // Text response/answer
	MediaURL    string `json:"media_url"`                     // URL to uploaded photo/video
	MediaType   string `json:"media_type"`                    // "image", "video", "audio"
	
	// Submission metadata
	Status       SubmissionStatus `json:"status" gorm:"default:pending"`
	PointsAwarded int             `json:"points_awarded"`              // Points given (may differ from quest points)
	AdminNotes   string          `json:"admin_notes" gorm:"type:text"` // Admin feedback
	ReviewedAt   *time.Time      `json:"reviewed_at"`                  // When admin reviewed
	ReviewedByID *uint           `json:"reviewed_by_id"`               // Admin who reviewed

	// Relationships
	User       User  `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Quest      Quest `json:"quest,omitempty" gorm:"foreignKey:QuestID"`
	ReviewedBy *User `json:"reviewed_by,omitempty" gorm:"foreignKey:ReviewedByID"`
}

// LeaderboardEntry represents a user's position on the leaderboard
type LeaderboardEntry struct {
	Rank        int    `json:"rank"`
	UserID      uint   `json:"user_id"`
	Username    string `json:"username"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Avatar      string `json:"avatar"`
	TotalPoints int    `json:"total_points"`
	QuestsCompleted int `json:"quests_completed"`
}
