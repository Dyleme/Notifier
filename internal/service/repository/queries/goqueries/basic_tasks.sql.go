// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: basic_tasks.sql

package goqueries

import (
	"context"

	domains "github.com/Dyleme/Notifier/internal/domains"
	"github.com/jackc/pgx/v5/pgtype"
)

const addBasicTask = `-- name: AddBasicTask :one
INSERT INTO basic_tasks (
  user_id,
  text,
  start,
  description,
  notification_params
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5
)
RETURNING id, created_at, text, description, user_id, start, notification_params
`

type AddBasicTaskParams struct {
	UserID             int32                       `db:"user_id"`
	Text               string                      `db:"text"`
	Start              pgtype.Timestamptz          `db:"start"`
	Description        pgtype.Text                 `db:"description"`
	NotificationParams *domains.NotificationParams `db:"notification_params"`
}

func (q *Queries) AddBasicTask(ctx context.Context, db DBTX, arg AddBasicTaskParams) (BasicTask, error) {
	row := db.QueryRow(ctx, addBasicTask,
		arg.UserID,
		arg.Text,
		arg.Start,
		arg.Description,
		arg.NotificationParams,
	)
	var i BasicTask
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Text,
		&i.Description,
		&i.UserID,
		&i.Start,
		&i.NotificationParams,
	)
	return i, err
}

const countListBasicTasks = `-- name: CountListBasicTasks :one
SELECT COUNT(*)
FROM basic_tasks
WHERE user_id = $1
`

func (q *Queries) CountListBasicTasks(ctx context.Context, db DBTX, userID int32) (int64, error) {
	row := db.QueryRow(ctx, countListBasicTasks, userID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const deleteBasicTask = `-- name: DeleteBasicTask :many
DELETE
FROM basic_tasks
WHERE id = $1
RETURNING id, created_at, text, description, user_id, start, notification_params
`

func (q *Queries) DeleteBasicTask(ctx context.Context, db DBTX, id int32) ([]BasicTask, error) {
	rows, err := db.Query(ctx, deleteBasicTask, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []BasicTask
	for rows.Next() {
		var i BasicTask
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.Text,
			&i.Description,
			&i.UserID,
			&i.Start,
			&i.NotificationParams,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getBasicTask = `-- name: GetBasicTask :one
SELECT id, created_at, text, description, user_id, start, notification_params
FROM basic_tasks
WHERE id = $1
`

func (q *Queries) GetBasicTask(ctx context.Context, db DBTX, id int32) (BasicTask, error) {
	row := db.QueryRow(ctx, getBasicTask, id)
	var i BasicTask
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Text,
		&i.Description,
		&i.UserID,
		&i.Start,
		&i.NotificationParams,
	)
	return i, err
}

const listBasicTasks = `-- name: ListBasicTasks :many
SELECT id, created_at, text, description, user_id, start, notification_params
FROM basic_tasks
WHERE user_id = $1
ORDER BY id DESC
LIMIT $3 OFFSET $2
`

type ListBasicTasksParams struct {
	UserID int32 `db:"user_id"`
	Off    int32 `db:"off"`
	Lim    int32 `db:"lim"`
}

func (q *Queries) ListBasicTasks(ctx context.Context, db DBTX, arg ListBasicTasksParams) ([]BasicTask, error) {
	rows, err := db.Query(ctx, listBasicTasks, arg.UserID, arg.Off, arg.Lim)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []BasicTask
	for rows.Next() {
		var i BasicTask
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.Text,
			&i.Description,
			&i.UserID,
			&i.Start,
			&i.NotificationParams,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateBasicTask = `-- name: UpdateBasicTask :one
UPDATE basic_tasks
SET start       = $1,
    text        = $2,
    description = $3,
    notification_params = $4
WHERE id = $5
  AND user_id = $6
RETURNING id, created_at, text, description, user_id, start, notification_params
`

type UpdateBasicTaskParams struct {
	Start              pgtype.Timestamptz          `db:"start"`
	Text               string                      `db:"text"`
	Description        pgtype.Text                 `db:"description"`
	NotificationParams *domains.NotificationParams `db:"notification_params"`
	ID                 int32                       `db:"id"`
	UserID             int32                       `db:"user_id"`
}

func (q *Queries) UpdateBasicTask(ctx context.Context, db DBTX, arg UpdateBasicTaskParams) (BasicTask, error) {
	row := db.QueryRow(ctx, updateBasicTask,
		arg.Start,
		arg.Text,
		arg.Description,
		arg.NotificationParams,
		arg.ID,
		arg.UserID,
	)
	var i BasicTask
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Text,
		&i.Description,
		&i.UserID,
		&i.Start,
		&i.NotificationParams,
	)
	return i, err
}