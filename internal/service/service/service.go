package service

import (
	"context"
	"time"
)

type Repository interface {
	Atomic(ctx context.Context, fn func(ctx context.Context, repo Repository) error) error
	DefaultNotificationParams() NotificationParamsRepository
	Tasks() TaskRepository
	Events() EventRepository
	TgImages() TgImagesRepository
	PeriodicEvents() PeriodicEventsRepository
}

type Service struct {
	repo        Repository
	notifierJob *NotifierJob
	notifier    Notifier
}

type Config struct {
	CheckTasksPeriod time.Duration
}

func New(_ context.Context, repo Repository, notifier Notifier, cfg Config) *Service {
	s := &Service{
		repo:        repo,
		notifierJob: NewNotifierJob(repo, notifier, cfg),
		notifier:    notifier,
	}

	return s
}