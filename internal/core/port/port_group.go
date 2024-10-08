package port

import (
	"context"
	"samsamoohooh-go-api/internal/core/domain"
	"samsamoohooh-go-api/internal/core/dto"
)

type GroupRepository interface {
	Create(ctx context.Context, group *domain.Group) (*domain.Group, error)
	GetByID(ctx context.Context, id uint) (*domain.Group, error)
	GetUsersByID(ctx context.Context, id uint) ([]*domain.User, error)
	GetPostsByID(ctx context.Context, id uint) ([]domain.Post, error)
	GetTasksByID(ctx context.Context, id uint) ([]domain.Task, error)
	GetAll(ctx context.Context, skip, limit int) ([]domain.Group, error)
	Update(ctx context.Context, id uint, group *domain.Group) (*domain.Group, error)
	Delete(ctx context.Context, id uint) error
}

type GroupService interface {
	Create(ctx context.Context, group *dto.GroupCreateRequest, creatorID uint) (*dto.GroupCreateResponse, error)
	GetByID(ctx context.Context, id uint) (*dto.GroupGetByIDResponse, error)
	GetUsersByID(ctx context.Context, id uint) ([]*dto.GroupGetUsersByIDResponse, error)
}
