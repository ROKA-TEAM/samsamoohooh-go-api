package repository

import (
	"context"
	"samsamoohooh-go-api/internal/adapter/persistence/sql/database"
	"samsamoohooh-go-api/internal/adapter/persistence/sql/repository/utils"
	"samsamoohooh-go-api/internal/core/domain"
	"samsamoohooh-go-api/internal/core/port"
)

var _ port.GroupRepository = (*GroupRepository)(nil)

type GroupRepository struct {
	database *database.Database
}

func NewGroupRepository(database *database.Database) *GroupRepository {
	return &GroupRepository{
		database: database,
	}
}

func (r *GroupRepository) Create(ctx context.Context, group *domain.Group) (*domain.Group, error) {
	err := r.database.WithContext(ctx).Create(group).Error
	if err != nil {
		return nil, utils.Wrap(err)
	}

	return group, nil
}

func (r *GroupRepository) GetByID(ctx context.Context, id uint) (*domain.Group, error) {
	gruop := domain.Group{}
	err := r.database.First(&gruop, id).Error
	if err != nil {
		return nil, utils.Wrap(err)
	}

	return &gruop, nil
}

func (r *GroupRepository) GetUsersByID(ctx context.Context, id uint) ([]*domain.User, error) {
	group := domain.Group{}
	err := r.database.WithContext(ctx).Preload("Users").First(group, id).Error
	if err != nil {
		return nil, utils.Wrap(err)
	}

	return group.Users, nil
}

func (r *GroupRepository) GetPostsByID(ctx context.Context, id uint) ([]*domain.Post, error) {
	posts := []*domain.Post{}
	err := r.database.WithContext(ctx).Preload("Posts").First(&posts, id).Error
	if err != nil {
		return nil, utils.Wrap(err)
	}

	return posts, nil
}
func (r *GroupRepository) GetTasksByID(ctx context.Context, id uint) ([]*domain.Task, error) {
	tasks := []*domain.Task{}
	err := r.database.WithContext(ctx).Preload("Tasks").First(&tasks, id).Error
	if err != nil {
		return nil, utils.Wrap(err)
	}

	return tasks, nil
}

func (r *GroupRepository) GetAll(ctx context.Context, skip, limit int) ([]*domain.Group, error) {
	groups := []*domain.Group{}
	err := r.database.WithContext(ctx).Limit(limit).Offset((skip - 1) * limit).Find(&groups).Error
	if err != nil {
		return nil, utils.Wrap(err)
	}

	return groups, nil
}

func (r *GroupRepository) Update(ctx context.Context, id uint, group *domain.Group) (*domain.Group, error) {
	group.Model.ID = id
	err := r.database.WithContext(ctx).Save(group).Error
	if err != nil {
		return nil, utils.Wrap(err)
	}

	return group, nil
}

func (r *GroupRepository) Delete(ctx context.Context, id uint) error {
	err := r.database.WithContext(ctx).Delete(ctx, id).Error
	if err != nil {
		return utils.Wrap(err)
	}

	return nil
}
