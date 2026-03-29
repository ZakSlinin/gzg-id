package repository

import (
	"context"
	"errors"
	"github.com/ZakSlinin/gzg-id/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GzgIDRepository interface {
	Create(ctx context.Context, user *model.User) (*model.User, error)
	FindByID(ctx context.Context, id uuid.UUID) (*model.User, error)
	UpdateByID(ctx context.Context, id uuid.UUID, user *model.User) (*model.User, error)
	SoftDeleteByID(ctx context.Context, id uuid.UUID) error
}

type PostgresGZGIDRepository struct {
	db *gorm.DB
}

func NewPostgresGZGIDRepository(db *gorm.DB) *PostgresGZGIDRepository {
	return &PostgresGZGIDRepository{db: db}
}

func (r *PostgresGZGIDRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	err := r.db.WithContext(ctx).Create(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *PostgresGZGIDRepository) FindByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	user := &model.User{}

	response := r.db.WithContext(ctx).Where("id = ?", id).First(user)

	if response.Error != nil {
		if errors.Is(response.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, response.Error
	}

	return user, nil
}

func (r *PostgresGZGIDRepository) UpdateByID(ctx context.Context, id uuid.UUID, user *model.User) (*model.User, error) {
	response := r.db.WithContext(ctx).Where("id = ?", id).Updates(user)
	if response.Error != nil {
		return nil, response.Error
	}

	if response.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (r *PostgresGZGIDRepository) SoftDeleteByID(ctx context.Context, id uuid.UUID) error {
	response := r.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", id).Update("is_active", false)

	if response.Error != nil {
		if errors.Is(response.Error, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return response.Error
	}

	return nil
}
