package domain

import (
	"time"

	"gorm.io/gorm"
)

type Priority string
type Status string

const (
	PriorityLow    Priority = "low"
	PriorityMedium Priority = "medium"
	PriorityHigh   Priority = "high"

	StatusTodo Status = "todo"
	StatusDone Status = "done"
)

type Todo struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	UserID      uint           `json:"user_id" gorm:"not null;index"`
	CategoryID  *uint          `json:"category_id" gorm:"index"`
	Title       string         `json:"title" gorm:"not null"`
	Description string         `json:"description"`
	Deadline    *time.Time     `json:"deadline"`
	Priority    Priority       `json:"priority" gorm:"default:medium"`
	Status      Status         `json:"status" gorm:"default:todo"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`

	// Relations
	User     User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Category *Category `json:"category,omitempty" gorm:"foreignKey:CategoryID"`
}

type CreateTodoRequest struct {
	Title       string     `json:"title" validate:"required"`
	Description string     `json:"description"`
	CategoryID  *uint      `json:"category_id"`
	Deadline    *time.Time `json:"deadline"`
	Priority    Priority   `json:"priority" validate:"omitempty,oneof=low medium high"`
}

type UpdateTodoRequest struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	CategoryID  *uint      `json:"category_id"`
	Deadline    *time.Time `json:"deadline"`
	Priority    Priority   `json:"priority" validate:"omitempty,oneof=low medium high"`
	Status      Status     `json:"status" validate:"omitempty,oneof=todo done"`
}

type TodoFilter struct {
	Status     Status   `json:"status"`
	Priority   Priority `json:"priority"`
	CategoryID uint     `json:"category_id"`
	Keyword    string   `json:"keyword"`
	Page       int      `json:"page"`
	Limit      int      `json:"limit"`
}
