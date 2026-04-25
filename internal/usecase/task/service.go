package task

import (
	"context"
	"fmt"
	"strings"
	"time"
	"strconv"

	taskdomain "example.com/taskservice/internal/domain/task"
)


type Service struct {
	repo Repository
	now  func() time.Time
}


func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
		now:  func() time.Time { return time.Now().UTC() },
	}
}


func (s *Service) Create(ctx context.Context, input CreateInput) (*taskdomain.Task, error) {
	normalized, err := validateCreateInput(input)
	if err != nil {
		return nil, err
	}

	if err := validateRecurrence(
		normalized.RecurrenceType,
		normalized.RecurrenceValue,
	); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidInput, err)
	}

	model := &taskdomain.Task{
		Title:            normalized.Title,
		Description:      normalized.Description,
		Status:           normalized.Status,
		RecurrenceType:   normalized.RecurrenceType,
		RecurrenceValue:  normalized.RecurrenceValue,
	}

	now := s.now()
	model.CreatedAt = now
	model.UpdatedAt = now

	return s.repo.Create(ctx, model)
}


func (s *Service) GetByID(ctx context.Context, id int64) (*taskdomain.Task, error) {
	if id <= 0 {
		return nil, fmt.Errorf("%w: id must be positive", ErrInvalidInput)
	}

	return s.repo.GetByID(ctx, id)
}


func (s *Service) Update(ctx context.Context, id int64, input UpdateInput) (*taskdomain.Task, error) {
	if id <= 0 {
		return nil, fmt.Errorf("%w: id must be positive", ErrInvalidInput)
	}

	normalized, err := validateUpdateInput(input)
	if err != nil {
		return nil, err
	}

	if err := validateRecurrence(
		normalized.RecurrenceType,
		normalized.RecurrenceValue,
	); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidInput, err)
	}

	model := &taskdomain.Task{
		ID:               id,
		Title:            normalized.Title,
		Description:      normalized.Description,
		Status:           normalized.Status,
		RecurrenceType:   normalized.RecurrenceType,
		RecurrenceValue:  normalized.RecurrenceValue,
		UpdatedAt:        s.now(),
	}

	updated, err := s.repo.Update(ctx, model)
	if err != nil {
		return nil, err
	}

	return updated, nil
}


func (s *Service) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return fmt.Errorf("%w: id must be positive", ErrInvalidInput)
	}

	return s.repo.Delete(ctx, id)
}


func (s *Service) List(ctx context.Context) ([]taskdomain.Task, error) {
	return s.repo.List(ctx)
}


func validateCreateInput(input CreateInput) (CreateInput, error) {
	input.Title = strings.TrimSpace(input.Title)
	input.Description = strings.TrimSpace(input.Description)

	if input.Title == "" {
		return CreateInput{}, fmt.Errorf("%w: title is required", ErrInvalidInput)
	}

	if input.Status == "" {
		input.Status = taskdomain.StatusNew
	}

	if !input.Status.Valid() {
		return CreateInput{}, fmt.Errorf("%w: invalid status", ErrInvalidInput)
	}

	return input, nil
}


func validateUpdateInput(input UpdateInput) (UpdateInput, error) {
	input.Title = strings.TrimSpace(input.Title)
	input.Description = strings.TrimSpace(input.Description)

	if input.Title == "" {
		return UpdateInput{}, fmt.Errorf("%w: title is required", ErrInvalidInput)
	}

	if !input.Status.Valid() {
		return UpdateInput{}, fmt.Errorf("%w: invalid status", ErrInvalidInput)
	}

	return input, nil
}


func validateRecurrence(
    recurrenceType taskdomain.RecurrenceType,
    recurrenceValue string,
) error {

    switch recurrenceType {

    case "none":
        return nil

    case "daily":
        interval, err := strconv.Atoi(recurrenceValue)
        if err != nil || interval <= 0 {
            return fmt.Errorf("daily interval must be positive")
        }

    case "monthly":
        day, err := strconv.Atoi(recurrenceValue)
        if err != nil || day < 1 || day > 30 {
            return fmt.Errorf("monthly day must be between 1 and 30")
        }

    case "specific_dates":
    	if recurrenceValue == "" {
        	return fmt.Errorf("specific dates cannot be empty")
    }

    case "even_days", "odd_days":
        return nil

    default:
        return fmt.Errorf("invalid recurrence type")
    }

    return nil
}