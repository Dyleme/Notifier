package service

import (
	"context"
	"fmt"
	"time"

	"github.com/Dyleme/Notifier/internal/domains"
	"github.com/Dyleme/Notifier/pkg/serverrors"
)

//go:generate mockgen -destination=mocks/basic_tasks_mocks.go -package=mocks . BasicTaskRepository
type BasicTaskRepository interface {
	Add(ctx context.Context, task domains.BasicTask) (domains.BasicTask, error)
	List(ctx context.Context, userID int, params ListParams) ([]domains.BasicTask, error)
	Update(ctx context.Context, task domains.BasicTask) (domains.BasicTask, error)
	Delete(ctx context.Context, taskID int) error
	Get(ctx context.Context, taskID int) (domains.BasicTask, error)
}

func (s *Service) CreateTask(ctx context.Context, task domains.BasicTask) (domains.BasicTask, error) {
	op := "Service.CreateTask: %w"
	var createdTask domains.BasicTask

	err := s.tr.Do(ctx, func(ctx context.Context) error {
		var err error

		createdTask, err = s.repo.Tasks().Add(ctx, task)
		if err != nil {
			return fmt.Errorf("add task: %w", err)
		}

		event := createdTask.NewEvent()
		_, err = s.repo.Events().Add(ctx, event)
		if err != nil {
			return fmt.Errorf("add event: %w", err)
		}

		return nil
	})
	if err != nil {
		err = fmt.Errorf(op, err)
		logError(ctx, err)

		return domains.BasicTask{}, err
	}

	s.notifierJob.UpdateWithTime(ctx, createdTask.Start)

	return createdTask, nil
}

func (s *Service) GetTask(ctx context.Context, userID, taskID int) (domains.BasicTask, error) {
	op := "Service.GetTask: %w"
	tt, err := s.repo.Tasks().Get(ctx, taskID)
	if err != nil {
		err = fmt.Errorf(op, err)
		logError(ctx, err)

		return domains.BasicTask{}, err
	}

	if !tt.BelongsTo(userID) {
		return domains.BasicTask{}, fmt.Errorf("belongs to: %w", serverrors.NewNotFoundError(err, "task"))
	}

	return tt, nil
}

func (s *Service) ListTasks(ctx context.Context, userID int, listParams ListParams) ([]domains.BasicTask, error) {
	op := "Service.ListTasks: %w"
	tts, err := s.repo.Tasks().List(ctx, userID, listParams)
	if err != nil {
		err = fmt.Errorf(op, err)
		logError(ctx, err)

		return nil, err
	}

	return tts, nil
}

func (s *Service) UpdateBasicTask(ctx context.Context, params domains.BasicTask, userID int) (domains.BasicTask, error) {
	var task domains.BasicTask
	err := s.tr.Do(ctx, func(ctx context.Context) error {
		t, err := s.repo.Tasks().Get(ctx, params.ID)
		if err != nil {
			return fmt.Errorf("get task: %w", err)
		}

		if !t.BelongsTo(userID) {
			return serverrors.NewNotFoundError(err, "task")
		}

		t.Text = params.Text
		t.Description = params.Description
		t.Start = params.Start

		task, err = s.repo.Tasks().Update(ctx, t)
		if err != nil {
			return fmt.Errorf("update task: %w", err)
		}

		event, err := s.repo.Events().GetLatest(ctx, t.ID)
		if err != nil {
			return fmt.Errorf("get latest event: %w", err)
		}

		event.Text = params.Text
		event.Description = params.Description
		event.SendTime = params.Start

		err = s.repo.Events().Update(ctx, event)
		if err != nil {
			return fmt.Errorf("update event: %w", err)
		}

		return nil
	})
	if err != nil {
		err = fmt.Errorf("tr: %w", err)
		logError(ctx, err)

		return domains.BasicTask{}, err
	}

	s.notifierJob.UpdateWithTime(ctx, task.Start)

	return task, nil
}

func (s *Service) createNewEventForPeriodicTask(ctx context.Context, taskID, userID int) error {
	err := s.tr.Do(ctx, func(ctx context.Context) error {
		bt, err := s.repo.PeriodicTasks().Get(ctx, taskID)
		if err != nil {
			return fmt.Errorf("periodic tasks get[taskID=%v,userID=%v]: %w", taskID, userID, err)
		}
		if !bt.BelongsTo(userID) {
			return serverrors.NewBusinessLogicError("task does not belong to user")
		}

		nextNotif, err := bt.NewEvent(time.Now())
		if err != nil {
			return fmt.Errorf("new event: %w", serverrors.NewBusinessLogicError(err.Error()))
		}

		_, err = s.repo.Events().Add(ctx, nextNotif)
		if err != nil {
			return fmt.Errorf("periodic tasks add event: %w", err)
		}

		s.notifierJob.UpdateWithTime(ctx, nextNotif.SendTime)

		return nil
	})
	if err != nil {
		return fmt.Errorf("tr: %w", err)
	}

	return nil
}

func (s *Service) SetEventDoneStatus(ctx context.Context, eventID, userID int, done bool) error {
	err := s.tr.Do(ctx, func(ctx context.Context) error {
		event, err := s.repo.Events().Get(ctx, eventID)
		if err != nil {
			return fmt.Errorf("get event: %w", err)
		}
		event.Done = done
		err = s.repo.Events().Update(ctx, event)
		if err != nil {
			return fmt.Errorf("update event: %w", err)
		}
		err = s.notifier.Delete(ctx, event.ID)
		if err != nil {
			return fmt.Errorf("delete notifier event: %w", err)
		}
		switch event.TaskType {
		case domains.BasicTaskType:
		case domains.PeriodicTaskType:
			err := s.createNewEventForPeriodicTask(ctx, event.TaskID, userID)
			if err != nil {
				return fmt.Errorf("setTaskDoneStatusPeriodicTask: %w", err)
			}
		default:
			return fmt.Errorf("unknown taskType[%v]", event.TaskType)
		}

		err = s.notifier.Delete(ctx, eventID)
		if err != nil {
			return fmt.Errorf("delete[eventID=%v]: %w", eventID, err)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("tr: %w", err)
	}

	return nil
}

func (s *Service) DeleteBasicTask(ctx context.Context, userID, taskID int) error {
	err := s.tr.Do(ctx, func(ctx context.Context) error {
		task, err := s.repo.Tasks().Get(ctx, taskID)
		if err != nil {
			return fmt.Errorf("get task: %w", err)
		}

		if !task.BelongsTo(userID) {
			return serverrors.NewNotFoundError(err, "basic task")
		}

		event, err := s.repo.Events().GetLatest(ctx, taskID)
		if err != nil {
			return fmt.Errorf("get latest event: %w", err)
		}

		err = s.repo.Events().Delete(ctx, event.ID)
		if err != nil {
			return fmt.Errorf("delete event: %w", err)
		}

		err = s.repo.Tasks().Delete(ctx, taskID)
		if err != nil {
			return fmt.Errorf("delete basic task: %w", err)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("tr: %w", err)
	}

	return nil
}