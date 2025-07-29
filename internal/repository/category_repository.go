package repository

import (
	"github.com/iskhakmuhamad/todo-api/internal/domain"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *domain.Category) error
	GetByUserID(userID uint) ([]domain.Category, error)
	GetByID(id, userID uint) (*domain.Category, error)
	Update(category *domain.Category) error
	Delete(id, userID uint) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Create(category *domain.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) GetByUserID(userID uint) ([]domain.Category, error) {
	var categories []domain.Category
	err := r.db.Where("user_id = ?", userID).Find(&categories).Error
	return categories, err
}

func (r *categoryRepository) GetByID(id, userID uint) (*domain.Category, error) {
	var category domain.Category
	err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) Update(category *domain.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepository) Delete(id, userID uint) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&domain.Category{}).Error
}
