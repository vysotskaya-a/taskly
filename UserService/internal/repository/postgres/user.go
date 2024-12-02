package postgres

import (
	sq "github.com/Masterminds/squirrel"
	"user-service/internal/errorz"

	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"user-service/internal/models"
)

const (
	usersTableName = "users"

	usersIDColumn           = "id"
	usersEmailColumn        = "email"
	usersPasswordHashColumn = "password_hash"
	usersGradeColumn        = "grade"
	usersCreatedAtColumn    = "created_at"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create создаёт нового пользователя.
func (r *UserRepository) Create(ctx context.Context, user *models.User) (string, error) {
	builder := sq.Insert(usersTableName).
		Columns(usersEmailColumn, usersPasswordHashColumn, usersGradeColumn).
		Values(user.Email, user.Password, user.Grade).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return "", err
	}

	var id string
	err = r.db.QueryRowContext(ctx, query, args...).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

// GetByID получает пользователя по его ID.
func (r *UserRepository) GetByID(ctx context.Context, id string) (*models.User, error) {
	builder := sq.Select(usersIDColumn, usersEmailColumn, usersPasswordHashColumn, usersGradeColumn, usersCreatedAtColumn).
		From(usersTableName).
		Where(sq.Eq{usersIDColumn: id}).
		Limit(1).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var user models.User
	err = r.db.GetContext(ctx, &user, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errorz.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

// GetByEmail получает пользователя по его почте.
func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	builder := sq.Select(usersIDColumn, usersEmailColumn, usersPasswordHashColumn, usersGradeColumn, usersCreatedAtColumn).
		From(usersTableName).
		Where(sq.Eq{usersEmailColumn: email}).
		Limit(1).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var user models.User
	err = r.db.GetContext(ctx, &user, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errorz.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

// Update обновляет данные пользователя.
func (r *UserRepository) Update(ctx context.Context, user *models.User) error {
	builder := sq.Update(usersTableName).
		Set(usersPasswordHashColumn, user.Password).
		Set(usersGradeColumn, user.Grade).
		Where(sq.Eq{usersIDColumn: user.ID}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	result, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errorz.ErrUserNotFound
	}

	return nil
}

// Delete удаляет пользователя по его ID.
func (r *UserRepository) Delete(ctx context.Context, id string) error {
	builder := sq.Delete(usersTableName).
		Where(sq.Eq{usersIDColumn: id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	result, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errorz.ErrUserNotFound
	}

	return nil
}
