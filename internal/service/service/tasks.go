package service

import (
	"context"
	"fmt"

	"github.com/Dyleme/Notifier/internal/domains"
)

//go:generate mockgen -destination=mocks/tasks_mocks.go -package=mocks . TaskRepository
type TaskRepository interface {
	Add(ctx context.Context, task domains.Task) (domains.Task, error)
	Get(ctx context.Context, taskID, userID int) (domains.Task, error)
	Delete(ctx context.Context, taskID, userID int) error
	Update(ctx context.Context, task domains.Task) error
	List(ctx context.Context, userID int, listParams ListParams) ([]domains.Task, error)
}

func (s *Service) AddTask(ctx context.Context, task domains.Task) (domains.Task, error) {
	op := "Service.AddTask: %w"
	createdTask, err := s.repo.Tasks().Add(ctx, task)
	if err != nil {
		err = fmt.Errorf(op, err)
		logError(ctx, err)

		return domains.Task{}, err
	}

	return createdTask, nil
}

func (s *Service) GetTask(ctx context.Context, taskID, userID int) (domains.Task, error) {
	op := "Service.GetTask: %w"
	task, err := s.repo.Tasks().Get(ctx, taskID, userID)
	if err != nil {
		err = fmt.Errorf(op, err)
		logError(ctx, err)

		return domains.Task{}, err
	}

	return task, nil
}

func (s *Service) UpdateTask(ctx context.Context, task domains.Task) error {
	op := "Service.UpdateTask: %w"
	err := s.repo.Tasks().Update(ctx, task)
	if err != nil {
		logError(ctx, fmt.Errorf(op, err))

		return fmt.Errorf(op, err)
	}

	return nil
}

func (s *Service) ListUserTasks(ctx context.Context, userID int, listParams ListParams) ([]domains.Task, error) {
	op := "Service.ListUserTasks: %w"
	tasks, err := s.repo.Tasks().List(ctx, userID, listParams)
	if err != nil {
		logError(ctx, fmt.Errorf(op, err))

		return nil, fmt.Errorf(op, err)
	}

	return tasks, nil
}

func (s *Service) DeleteTask(ctx context.Context, taskID, userID int) error {
	op := "Service.DeleteTask: %w"

	err := s.repo.Tasks().Delete(ctx, taskID, userID)
	if err != nil {
		logError(ctx, fmt.Errorf(op, err))

		return fmt.Errorf(op, err)
	}

	return nil
}
