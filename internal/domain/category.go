package domain

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	UserID      uint           `json:"user_id" gorm:"not null;index"`
	Name        string         `json:"name" gorm:"not null"`
	Description string         `json:"description"`
	Color       string         `json:"color" gorm:"default:#3B82F6"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`

	// Relations
	User  User   `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Todos []Todo `json:"todos,omitempty" gorm:"foreignKey:CategoryID"`
}

type CreateCategoryRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Color       string `json:"color"`
}

type UpdateCategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
}
