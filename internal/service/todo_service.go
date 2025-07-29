package service

import (
	"github.com/iskhakmuhamad/todo-api/internal/domain"
	"github.com/iskhakmuhamad/todo-api/internal/repository"
)

type TodoService interface {
	Create(userID uint, req domain.CreateTodoRequest) (*domain.Todo, error)
	GetAll(userID uint, filter domain.TodoFilter) ([]domain.Todo, int64, error)
	GetByID(id, userID uint) (*domain.Todo, error)
	Update(id, userID uint, req domain.UpdateTodoRequest) (*domain.Todo, error)
	Delete(id, userID uint) error
}

type todoService struct {
	todoRepo repository.TodoRepository
}

func NewTodoService(todoRepo repository.TodoRepository) TodoService {
	return &todoService{todoRepo: todoRepo}
}

func (s *todoService) Create(userID uint, req domain.CreateTodoRequest) (*domain.Todo, error) {
	todo := &domain.Todo{
		UserID:      userID,
		Title:       req.Title,
		Description: req.Description,
		CategoryID:  req.CategoryID,
		Deadline:    req.Deadline,
		Priority:    req.Priority,
		Status:      domain.StatusTodo,
	}

	if todo.Priority == "" {
		todo.Priority = domain.PriorityMedium
	}

	if err := s.todoRepo.Create(todo); err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *todoService) GetAll(userID uint, filter domain.TodoFilter) ([]domain.Todo, int64, error) {
	if filter.Page == 0 {
		filter.Page = 1
	}
	if filter.Limit == 0 {
		filter.Limit = 10
	}

	return s.todoRepo.GetByUserID(userID, filter)
}

func (s *todoService) GetByID(id, userID uint) (*domain.Todo, error) {
	return s.todoRepo.GetByID(id, userID)
}

func (s *todoService) Update(id, userID uint, req domain.UpdateTodoRequest) (*domain.Todo, error) {
	todo, err := s.todoRepo.GetByID(id, userID)
	if err != nil {
		return nil, err
	}

	if req.Title != "" {
		todo.Title = req.Title
	}
	if req.Description != "" {
		todo.Description = req.Description
	}
	if req.CategoryID != nil {
		todo.CategoryID = req.CategoryID
	}
	if req.Deadline != nil {
		todo.Deadline = req.Deadline
	}
	if req.Priority != "" {
		todo.Priority = req.Priority
	}
	if req.Status != "" {
		todo.Status = req.Status
	}

	if err := s.todoRepo.Update(todo); err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *todoService) Delete(id, userID uint) error {
	return s.todoRepo.Delete(id, userID)
}
