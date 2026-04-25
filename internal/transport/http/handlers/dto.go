package handlers

import (
	"time"

	taskdomain "example.com/taskservice/internal/domain/task"
)

type taskMutationDTO struct {
    Title            string                    `json:"title"`
    Description      string                    `json:"description"`
    Status           taskdomain.Status         `json:"status"`
    RecurrenceType   taskdomain.RecurrenceType `json:"recurrence_type"`
    RecurrenceValue  string                    `json:"recurrence_value"`
}

type taskDTO struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Status taskdomain.Status `json:"status"`
	RecurrenceType string `json:"recurrence_type"`
	RecurrenceValue string `json:"recurrence_value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func newTaskDTO(task *taskdomain.Task) taskDTO {
	return taskDTO{
		ID: task.ID,
		Title: task.Title,
		Description: task.Description,
		Status: task.Status,
		RecurrenceType: string(task.RecurrenceType),
		RecurrenceValue: task.RecurrenceValue,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
}
