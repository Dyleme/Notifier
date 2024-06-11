package service

import (
	"context"
	"fmt"
	"time"

	"github.com/Dyleme/Notifier/internal/domains"
	"github.com/Dyleme/Notifier/pkg/serverrors"
)

//go:generate mockgen -destination=mocks/periodic_tasks_mocks.go -package=mocks . PeriodicTasksRepository
type PeriodicTasksRepository interface {
	Add(ctx context.Context, task domains.PeriodicTask) (domains.PeriodicTask, error)
	Get(ctx context.Context, taskID int) (domains.PeriodicTask, error)
	Update(ctx context.Context, task domains.PeriodicTask) (domains.PeriodicTask, error)
	Delete(ctx context.Context, taskID int) error
	List(ctx context.Context, userID int, listParams ListParams) ([]domains.PeriodicTask, error)
}

func (s *Service) AddPeriodicTask(ctx context.Context, perTask domains.PeriodicTask, userID int) (domains.PeriodicTask, error) {
	if !perTask.BelongsTo(userID) {
		return domains.PeriodicTask{}, fmt.Errorf("belongs to: %w", serverrors.NewBusinessLogicError("you are not allowed to add task to another user"))
	}

	var createdPerTask domains.PeriodicTask
	err := s.tr.Do(ctx, func(ctx context.Context) error {
		createdPerTask, err := s.repo.PeriodicTasks().Add(ctx, perTask)
		if err != nil {
			return fmt.Errorf("add periodic task: %w", err)
		}

		event, err := createdPerTask.NewEvent(time.Now())
		if err != nil {
			return fmt.Errorf("next event: %w", serverrors.NewBusinessLogicError(err.Error()))
		}

		_, err = s.repo.Events().Add(ctx, event)
		if err != nil {
			return fmt.Errorf("add event: %w", err)
		}

		return nil
	})
	if err != nil {
		return domains.PeriodicTask{}, fmt.Errorf("tr: %w", err)
	}

	return createdPerTask, nil
}

func (s *Service) GetPeriodicTask(ctx context.Context, taskID, userID int) (domains.PeriodicTask, error) {
	var pt domains.PeriodicTask
	err := s.tr.Do(ctx, func(ctx context.Context) error {
		var err error
		pt, err = s.repo.PeriodicTasks().Get(ctx, taskID)
		if err != nil {
			return fmt.Errorf("get[taskID=%v]: %w", taskID, err)
		}

		if !pt.BelongsTo(userID) {
			return serverrors.NewNotFoundError(err, "periodic task")
		}

		return nil
	})
	if err != nil {
		return domains.PeriodicTask{}, fmt.Errorf("tr: %w", err)
	}

	return pt, nil
}

type UpdatePeriodicTaskParams struct {
	ID             int
	Text           string
	Description    string
	UserID         int
	Start          time.Duration // Event time from beginning of day
	SmallestPeriod time.Duration
	BiggestPeriod  time.Duration
	EventParams    *domains.NotificationParams
}

func (s *Service) UpdatePeriodicTask(ctx context.Context, perTask domains.PeriodicTask, userID int) error {
	var updatedTask domains.PeriodicTask
	err := s.tr.Do(ctx, func(ctx context.Context) error {
		var pt domains.PeriodicTask
		pt, err := s.repo.PeriodicTasks().Get(ctx, perTask.ID)
		if err != nil {
			return fmt.Errorf("get[taskID=%v]: %w", perTask.ID, err)
		}
		if !pt.BelongsTo(userID) {
			return serverrors.NewNotFoundError(err, "periodic task")
		}

		updatedTask, err = s.repo.PeriodicTasks().Update(ctx, perTask)
		if err != nil {
			return fmt.Errorf("update: %w", err)
		}

		event, err := s.repo.Events().GetLatest(ctx, pt.ID)
		if err != nil {
			return fmt.Errorf("delete event: %w", err)
		}

		nextNotif, err := updatedTask.NewEvent(time.Now())
		if err != nil {
			return fmt.Errorf("next event: %w", serverrors.NewBusinessLogicError(err.Error()))
		}

		nextNotif.ID = event.ID

		err = s.repo.Events().Update(ctx, nextNotif)
		if err != nil {
			return fmt.Errorf("add event: %w", err)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("atomoic: %w", err)
	}

	return nil
}

func (s *Service) DeletePeriodicTask(ctx context.Context, taskID, userID int) error {
	op := "Service.DeletePeriodicTask: %w"

	err := s.tr.Do(ctx, func(ctx context.Context) error {
		task, err := s.repo.PeriodicTasks().Get(ctx, taskID)
		if err != nil {
			return fmt.Errorf("get current events: %w", err)
		}

		if task.BelongsTo(userID) {
			return serverrors.NewNotFoundError(err, "periodic task")
		}

		event, err := s.repo.Events().GetLatest(ctx, taskID)
		if err != nil {
			return fmt.Errorf("get latest event: %w", err)
		}

		err = s.repo.Events().Delete(ctx, event.ID)
		if err != nil {
			return fmt.Errorf("delete event: %w", err)
		}

		err = s.repo.PeriodicTasks().Delete(ctx, taskID)
		if err != nil {
			return fmt.Errorf("delete periodic task: %w", err)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf(op, err)
	}

	return nil
}

func (s *Service) ListPeriodicTasks(ctx context.Context, userID int, listParams ListParams) ([]domains.PeriodicTask, error) {
	tasks, err := s.repo.PeriodicTasks().List(ctx, userID, listParams)
	if err != nil {
		err = fmt.Errorf("list periodic tasks: %w", err)
		logError(ctx, err)

		return nil, err
	}

	return tasks, nil
}
