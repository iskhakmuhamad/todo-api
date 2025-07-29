package repository

import (
	"github.com/iskhakmuhamad/todo-api/internal/domain"

	"gorm.io/gorm"
)

type TodoRepository interface {
	Create(todo *domain.Todo) error
	GetByUserID(userID uint, filter domain.TodoFilter) ([]domain.Todo, int64, error)
	GetByID(id, userID uint) (*domain.Todo, error)
	Update(todo *domain.Todo) error
	Delete(id, userID uint) error
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{db: db}
}

func (r *todoRepository) Create(todo *domain.Todo) error {
	return r.db.Create(todo).Error
}

func (r *todoRepository) GetByUserID(userID uint, filter domain.TodoFilter) ([]domain.Todo, int64, error) {
	var todos []domain.Todo
	var total int64

	query := r.db.Model(&domain.Todo{}).Where("user_id = ?", userID)

	// Apply filters
	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}
	if filter.Priority != "" {
		query = query.Where("priority = ?", filter.Priority)
	}
	if filter.CategoryID > 0 {
		query = query.Where("category_id = ?", filter.CategoryID)
	}
	if filter.Keyword != "" {
		query = query.Where("title ILIKE ? OR description ILIKE ?", "%"+filter.Keyword+"%", "%"+filter.Keyword+"%")
	}

	// Count total
	query.Count(&total)

	// Apply pagination
	if filter.Page > 0 && filter.Limit > 0 {
		offset := (filter.Page - 1) * filter.Limit
		query = query.Offset(offset).Limit(filter.Limit)
	}

	err := query.Preload("Category").Order("created_at DESC").Find(&todos).Error
	return todos, total, err
}

func (r *todoRepository) GetByID(id, userID uint) (*domain.Todo, error) {
	var todo domain.Todo
	err := r.db.Where("id = ? AND user_id = ?", id, userID).Preload("Category").First(&todo).Error
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *todoRepository) Update(todo *domain.Todo) error {
	return r.db.Save(todo).Error
}

func (r *todoRepository) Delete(id, userID uint) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&domain.Todo{}).Error
}
