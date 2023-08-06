package service

import (
	"context"
	"time"
)

type Repository interface {
	Atomic(ctx context.Context, fn func(ctx context.Context, repo Repository) error) error
	DefaultNotificationParams() NotificationParamsRepository
	Tasks() TaskRepository
	TimetableTasks() TimetableTaskRepository
}

type Service struct {
	repo            Repository
	notifier        Notifier
	checkTaskPeriod time.Duration
}

type Config struct {
	CheckTasksPeriod time.Duration
}

func New(repo Repository, notifier Notifier, cfg Config) *Service {
	return &Service{repo: repo, notifier: notifier, checkTaskPeriod: cfg.CheckTasksPeriod}
}
