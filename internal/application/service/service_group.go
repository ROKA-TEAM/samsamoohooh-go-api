package service

import (
	"context"
	"samsamoohooh-go-api/internal/application/domain"
)

var _ domain.GroupService = (*GroupService)(nil)

type GroupService struct {
	groupRepository domain.GroupRepository
	userService     domain.UserService
	taskService     domain.TaskService
}

func NewGroupService(
	groupRepository domain.GroupRepository,
	userService domain.UserService,
	taskService domain.TaskService,
) *GroupService {
	return &GroupService{
		groupRepository: groupRepository,
		userService:     userService,
		taskService:     taskService,
	}
}

func (s *GroupService) CreateGroup(ctx context.Context, userID int, group *domain.Group) (*domain.Group, error) {
	createdUser, err := s.groupRepository.CreateGroup(ctx, userID, group)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (s *GroupService) GetGroups(ctx context.Context, offset, limit int) ([]*domain.Group, error) {
	listGroup, err := s.groupRepository.GetGroups(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	return listGroup, nil
}

func (s *GroupService) GetByGroupID(ctx context.Context, id int) (*domain.Group, error) {
	gotGroup, err := s.groupRepository.GetByGroupID(ctx, id)
	if err != nil {
		return nil, err
	}

	return gotGroup, nil
}

func (s *GroupService) GetUsersByGroupID(ctx context.Context, id int, offset, limit int) ([]*domain.User, error) {
	listUser, err := s.groupRepository.GetUsersByGroupID(ctx, id, offset, limit)
	if err != nil {
		return nil, err
	}

	return listUser, nil
}

func (s *GroupService) GetPostsByGroupID(ctx context.Context, id int, offset, limit int) ([]*domain.Post, error) {
	listGroup, err := s.groupRepository.GetPostsByGroupID(ctx, id, offset, limit)
	if err != nil {
		return nil, err
	}

	return listGroup, nil
}

func (s *GroupService) GetTasksByGroupID(ctx context.Context, id int, offset, limit int) ([]*domain.Task, error) {
	listTask, err := s.groupRepository.GetTasksByGroupID(ctx, id, offset, limit)
	if err != nil {
		return nil, err
	}

	return listTask, nil
}

func (s *GroupService) UpdateGroup(ctx context.Context, id int, group *domain.Group) (*domain.Group, error) {
	updatedGroup, err := s.groupRepository.UpdateGroup(ctx, id, group)
	if err != nil {
		return nil, err
	}

	return updatedGroup, nil
}

func (s *GroupService) DeleteGroup(ctx context.Context, id int) error {
	err := s.groupRepository.DeleteGroup(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *GroupService) StartDiscussion(ctx context.Context, groupID, taskID int) (topics []string, userNames []string, err error) {

	topicsLen, err := s.taskService.GetTopicsLenByTaskID(ctx, taskID)
	if err != nil {
		return nil, nil, err
	}

	listTopic, err := s.taskService.GetTopicsByTaskID(ctx, taskID, 0, topicsLen)
	if err != nil {
		return nil, nil, err
	}

	usersLen, err := s.groupRepository.GetUsersLenByGroupID(ctx, groupID)
	if err != nil {
		return nil, nil, err
	}
	listUser, err := s.groupRepository.GetUsersByGroupID(ctx, groupID, 0, usersLen)
	if err != nil {
		return nil, nil, err
	}

	// bookmark
	gotTask, err := s.taskService.GetByTaskID(ctx, taskID)
	if err != nil {
		return nil, nil, err
	}

	// group bookmark 설정
	_, err = s.groupRepository.UpdateGroup(ctx, groupID, &domain.Group{
		Bookmark: gotTask.Range,
	})
	if err != nil {
		return nil, nil, err
	}

	for _, topic := range listTopic {
		topics = append(topics, topic.Topic)
	}

	for _, user := range listUser {
		userNames = append(userNames, user.Name)
	}

	return topics, userNames, nil
}
