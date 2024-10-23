package service

import (
	"pushNotifications/pkg/repository"
)

type AuthService interface {
	Login(username, password string) (string, error)
	ValidateToken(token string) (bool, string, error)
}

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) Login(username, password string) (string, error) {
	userID, err := s.repo.VerifyCredentials(username, password)
	if err != nil {
		return "", err
	}
	token, err := s.repo.GenerateToken(userID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *authService) ValidateToken(token string) (bool, string, error) {
	return s.repo.ValidateToken(token)
}
