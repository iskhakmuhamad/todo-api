package service

import (
	"github.com/iskhakmuhamad/todo-api/internal/domain"
	"github.com/iskhakmuhamad/todo-api/internal/repository"
)

type CategoryService interface {
	Create(userID uint, req domain.CreateCategoryRequest) (*domain.Category, error)
	GetAll(userID uint) ([]domain.Category, error)
	GetByID(id, userID uint) (*domain.Category, error)
	Update(id, userID uint, req domain.UpdateCategoryRequest) (*domain.Category, error)
	Delete(id, userID uint) error
}

type categoryService struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryService(categoryRepo repository.CategoryRepository) CategoryService {
	return &categoryService{categoryRepo: categoryRepo}
}

func (s *categoryService) Create(userID uint, req domain.CreateCategoryRequest) (*domain.Category, error) {
	category := &domain.Category{
		UserID:      userID,
		Name:        req.Name,
		Description: req.Description,
		Color:       req.Color,
	}

	if category.Color == "" {
		category.Color = "#3B82F6"
	}

	if err := s.categoryRepo.Create(category); err != nil {
		return nil, err
	}

	return category, nil
}

func (s *categoryService) GetAll(userID uint) ([]domain.Category, error) {
	return s.categoryRepo.GetByUserID(userID)
}

func (s *categoryService) GetByID(id, userID uint) (*domain.Category, error) {
	return s.categoryRepo.GetByID(id, userID)
}

func (s *categoryService) Update(id, userID uint, req domain.UpdateCategoryRequest) (*domain.Category, error) {
	category, err := s.categoryRepo.GetByID(id, userID)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		category.Name = req.Name
	}
	if req.Description != "" {
		category.Description = req.Description
	}
	if req.Color != "" {
		category.Color = req.Color
	}

	if err := s.categoryRepo.Update(category); err != nil {
		return nil, err
	}

	return category, nil
}

func (s *categoryService) Delete(id, userID uint) error {
	return s.categoryRepo.Delete(id, userID)
}
