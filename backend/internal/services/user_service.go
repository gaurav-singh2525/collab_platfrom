package services

import (
	"collab-code-platform/internal/repositories"
	"collab-code-platform/internal/auth"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo      *repositories.UserRepository
	jwtSecret string
}

func NewUserService(
	repo *repositories.UserRepository,
	jwtSecret string,
) *UserService {
	return &UserService{
		repo:      repo,
		jwtSecret: jwtSecret,
	}
}

func (s *UserService) HashPassword(
	password string,
) (string, error) {

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (s *UserService) CreateUser(
	email string,
	password string,
) error {

	exists, err := s.repo.EmailExists(email)

	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("email already exists")
	}

	hash, err := s.HashPassword(password)

	if err != nil {
		return err
	}

	return s.repo.CreateUser(
		email,
		hash,
	)
}

func (s *UserService) VerifyPassword(
	hash string,
	password string,
) error {

	return bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	)
}

func (s *UserService) Login(
	email string,
	password string,
) (string, error) {

	user, err := s.repo.GetUserByEmail(email)

	if err != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	err = s.VerifyPassword(
		user.PasswordHash,
		password,
	)

	if err != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	token, err := auth.GenerateToken(
		user.ID,
		user.Email,
		s.jwtSecret,
	)

	if err != nil {
		return "", err
	}

	return token, nil
}
