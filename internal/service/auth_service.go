package service

import (
	"errors"

	"github.com/iskhakmuhamad/todo-api/internal/domain"
	"github.com/iskhakmuhamad/todo-api/internal/repository"
	"github.com/iskhakmuhamad/todo-api/pkg/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(req domain.RegisterRequest) (*domain.User, error)
	Login(req domain.LoginRequest) (*domain.LoginResponse, error)
}

type authService struct {
	userRepo  repository.UserRepository
	jwtSecret string
}

func NewAuthService(userRepo repository.UserRepository, jwtSecret string) AuthService {
	return &authService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

func (s *authService) Register(req domain.RegisterRequest) (*domain.User, error) {
	// Check if user already exists
	existingUser, _ := s.userRepo.GetByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Email:    req.Email,
		Username: req.Username,
		Password: string(hashedPassword),
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *authService) Login(req domain.LoginRequest) (*domain.LoginResponse, error) {
	user, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid credentials")
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	token, err := utils.GenerateJWT(user.ID, user.Email, s.jwtSecret, 24)
	if err != nil {
		return nil, err
	}

	return &domain.LoginResponse{
		Token: token,
		User:  *user,
	}, nil
}
