package service

import (
	"context"
	"github.com/ZakSlinin/gzg-id/model"
	"github.com/ZakSlinin/gzg-id/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type GZGIDService struct {
	repo repository.GZGIDRepository
}

func NewGZGIDService(repo repository.GZGIDRepository) *GZGIDService {
	return &GZGIDService{repo: repo}
}

func (s *GZGIDService) Create(ctx context.Context, createUserRequest *model.CreateUserRequest) (*model.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(createUserRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := &model.User{
		ID:           uuid.New(),
		Email:        createUserRequest.Email,
		PasswordHash: hash,
		Username:     createUserRequest.Username,
		Avatar:       createUserRequest.Avatar,
		FirstName:    createUserRequest.FirstName,
		Surname:      createUserRequest.Surname,
		IsVerified:   false,
		IsActive:     true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	createdUser, err := s.repo.Create(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
