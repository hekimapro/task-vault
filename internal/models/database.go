package models

import (
	"time"

	"github.com/google/uuid"
)

type CommonFields struct {
	ID        uuid.UUID  `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

// user model
type User struct {
	CommonFields
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// task modal
type Task struct {
	CommonFields
	Title          string     `json:"title"`
	Description    *string    `json:"description"`
	DueDate        *time.Time `json:"due_date"`
	Status         string     `json:"status"`
	UserID         uuid.UUID  `json:"user_id"`
	CreatedByName  *string    `json:"created_by_name"`
	CreatedByEmail *string    `json:"created_by_email"`
}
