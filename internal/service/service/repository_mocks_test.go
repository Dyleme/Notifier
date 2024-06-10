package service_test

import (
	"context"

	"github.com/Dyleme/Notifier/internal/service/service"
)

type RepositoryMock struct {
	DefaultNotificationRepo service.NotificationParamsRepository
	TasksRepo               service.TaskRepository
	EventsRepo              service.BasicEventRepository
	TgImagesRepo            service.TgImagesRepository
	PeriodicEventsRepo      service.PeriodicEventsRepository
	NotificationsRepo       service.NotificationsRepository
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

func (r *RepositoryMock) Events() service.BasicEventRepository {
	return r.EventsRepo
}

func (r *RepositoryMock) TgImages() service.TgImagesRepository {
	return r.TgImagesRepo
}

func (r *RepositoryMock) PeriodicEvents() service.PeriodicEventsRepository {
	return r.PeriodicEventsRepo
}

func (r *RepositoryMock) Notifications() service.NotificationsRepository {
	return r.NotificationsRepo
}
