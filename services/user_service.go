package services

import (
	"context"
	"restapi-golang/interfaces"
	"restapi-golang/models"
)

type UserService struct {
	UserRepository interfaces.UserRepository
}

func NewUserService(repo interfaces.UserRepository) *UserService {
	return &UserService{
		UserRepository: repo,
	}
}

func (service *UserService) CheckUserByUsername(ctx context.Context, username string) (*models.User, error) {
	return service.UserRepository.CheckUserByUsername(ctx, username)
}
