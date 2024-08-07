package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/Dyleme/Notifier/internal/authorization/repository/queries/goqueries"
	"github.com/Dyleme/Notifier/internal/authorization/service"
	"github.com/Dyleme/Notifier/internal/domains"
	"github.com/Dyleme/Notifier/pkg/serverrors"
	"github.com/Dyleme/Notifier/pkg/sql/pgxconv"
)

func (r *Repository) Create(ctx context.Context, input service.CreateUserInput) (domains.User, error) {
	op := "Repository.Create: %w"
	user, err := r.q.AddUser(ctx, goqueries.AddUserParams{
		Email:        pgxconv.Text(input.Email),
		PasswordHash: pgxconv.Text(input.Password),
		TgID:         pgxconv.Int4(input.TGID),
	})
	if err != nil {
		if intersection, isUnique := uniqueError(err); isUnique {
			return domains.User{}, fmt.Errorf(op, serverrors.NewUniqueError(intersection, input.Email))
		}

		return domains.User{}, fmt.Errorf(op, serverrors.NewRepositoryError(err))
	}

	return domains.User{
		ID:             int(user.ID),
		Email:          pgxconv.String(user.Email),
		PasswordHash:   pgxconv.ByteSlice(user.PasswordHash),
		TGID:           pgxconv.Int(user.TgID),
		TimeZoneOffset: 0,
		IsTimeZoneDST:  false,
	}, nil
}

func uniqueError(err error) (string, bool) {
	var pgerr *pgconn.PgError
	if errors.As(err, &pgerr) {
		if pgerr.Code == pgerrcode.UniqueViolation {
			if strings.Contains(pgerr.Detail, "email") {
				return "email", true
			}

			return "", true
		}
	}

	return "", false
}

func (r *Repository) Get(ctx context.Context, email string, tgID *int) (domains.User, error) {
	op := "Repository.Get: %w"
	out, err := r.q.FindUser(ctx, goqueries.FindUserParams{
		Email: pgxconv.Text(email),
		TgID:  pgxconv.Int4(tgID),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domains.User{}, fmt.Errorf(op, serverrors.NewNotFoundError(err, "user"))
		}

		return domains.User{}, fmt.Errorf(op, serverrors.NewRepositoryError(err))
	}

	return domains.User{
		ID:             int(out.ID),
		Email:          pgxconv.String(out.Email),
		PasswordHash:   pgxconv.ByteSlice(out.PasswordHash),
		TGID:           pgxconv.Int(out.TgID),
		TimeZoneOffset: int(out.TimezoneOffset),
		IsTimeZoneDST:  out.TimezoneDst,
	}, nil
}

func (r *Repository) UpdateTime(ctx context.Context, id int, tzOffset domains.TimeZoneOffset, isDST bool) error {
	op := "Repository.UpdateWithTime: %w"

	err := r.q.UpdateTime(ctx, goqueries.UpdateTimeParams{
		TimezoneOffset: int32(tzOffset),
		IsDst:          isDST,
		ID:             int32(id),
	})
	if err != nil {
		return fmt.Errorf(op, serverrors.NewRepositoryError(err))
	}

	return nil
}
