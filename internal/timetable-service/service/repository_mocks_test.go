package service_test

import (
	"context"

	"github.com/Dyleme/Notifier/internal/timetable-service/service"
)

type RepositoryMock struct {
	DefaultNotificationRepo service.NotificationParamsRepository
	TasksRepo               service.TaskRepository
	EventsRepo              service.EventRepository
	TgImagesRepo            service.TgImagesRepository
}

func (r *RepositoryMock) Atomic(ctx context.Context, fn func(ctx context.Context, repo service.Repository) error) error {
	return fn(ctx, r)
}

func (r *RepositoryMock) DefaultNotificationParams() service.NotificationParamsRepository {
	return r.DefaultNotificationRepo
}

func (r *RepositoryMock) Tasks() service.TaskRepository {
	return r.TasksRepo
}

func (r *RepositoryMock) Events() service.EventRepository {
	return r.EventsRepo
}

func (r *RepositoryMock) TgImages() service.TgImagesRepository {
	return r.TgImagesRepo
}