package presenter

import (
	"samsamoohooh-go-api/internal/domain"
	"time"
)

type TaskCreateResponse struct {
	ID        int       `json:"id"`
	Deadline  time.Time `json:"deadline"`
	Range     int       `json:"range"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewTaskCreateResponse(task *domain.Task) *TaskCreateResponse {
	return &TaskCreateResponse{
		ID:        task.ID,
		Deadline:  task.Deadline,
		Range:     task.Range,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
}

type TaskListResponse struct {
	ID        int       `json:"id"`
	Deadline  time.Time `json:"deadline"`
	Range     int       `json:"range"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewTaskListResponse(tasks []*domain.Task) []*TaskListResponse {
	var taskList []*TaskListResponse
	for _, task := range tasks {
		taskList = append(taskList, &TaskListResponse{
			ID:        task.ID,
			Deadline:  task.Deadline,
			Range:     task.Range,
			CreatedAt: task.CreatedAt,
			UpdatedAt: task.UpdatedAt,
		})
	}
	return taskList
}

type TaskGetByIDResponse struct {
	ID        int       `json:"id"`
	Deadline  time.Time `json:"deadline"`
	Range     int       `json:"range"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewTaskGetByIDResponse(task *domain.Task) *TaskGetByIDResponse {
	return &TaskGetByIDResponse{
		ID:        task.ID,
		Deadline:  task.Deadline,
		Range:     task.Range,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
}

type TaskGetTopicsByIDResponse struct {
	ID        int       `json:"id"`
	Topic     string    `json:"topic"`
	Feeling   string    `json:"feeling"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewTaskGetTopicsByIDResponse(topics []*domain.Topic) []*TaskGetTopicsByIDResponse {
	var topicList []*TaskGetTopicsByIDResponse
	for _, topic := range topics {
		topicList = append(topicList, &TaskGetTopicsByIDResponse{
			ID:        topic.ID,
			Topic:     topic.Topic,
			Feeling:   topic.Feeling,
			CreatedAt: topic.CreatedAt,
			UpdatedAt: topic.UpdatedAt,
		})
	}
	return topicList
}

type TaskUpdateResponse struct {
	ID        int       `json:"id"`
	Deadline  time.Time `json:"deadline"`
	Range     int       `json:"range"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewTaskUpdateResponse(task *domain.Task) *TaskUpdateResponse {
	return &TaskUpdateResponse{
		ID:        task.ID,
		Deadline:  task.Deadline,
		Range:     task.Range,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
}
