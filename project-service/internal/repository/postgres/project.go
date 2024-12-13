package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"project-service/internal/errorz"

	"github.com/jmoiron/sqlx"
	"project-service/internal/models"
)

const (
	projectsTableName = "projects"

	projectsIDColumn          = "id"
	projectsTitleColumn       = "title"
	projectsDescriptionColumn = "description"
	projectsUsersColumn       = "users"
	projectsAdminIDColumn     = "admin_id"
	projectsCreatedAtColumn   = "created_at"
)

type ProjectRepository struct {
	db *sqlx.DB
}

func NewProjectRepository(db *sqlx.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (r *ProjectRepository) Create(ctx context.Context, project *models.Project) (string, error) {
	const op = "Postgres.ProjectRepository.Create"

	builder := sq.Insert(projectsTableName).
		Columns(projectsTitleColumn, projectsDescriptionColumn, projectsUsersColumn, projectsAdminIDColumn).
		Values(project.Title, project.Description, project.Users, project.AdminID).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	var id string
	if err = r.db.QueryRowxContext(ctx, query, args...).Scan(&id); err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (r *ProjectRepository) GetByID(ctx context.Context, id string) (*models.Project, error) {
	const op = "Postgres.ProjectRepository.GetByID"

	builder := sq.Select(projectsIDColumn, projectsTitleColumn, projectsDescriptionColumn, projectsUsersColumn, projectsAdminIDColumn, projectsCreatedAtColumn).
		From(projectsTableName).
		Where(sq.Eq{projectsIDColumn: id}).
		PlaceholderFormat(sq.Dollar).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var project models.Project
	err = r.db.GetContext(ctx, &project, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, errorz.ErrProjectNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &project, nil
}

func (r *ProjectRepository) GetAllByUserID(ctx context.Context, userID string) ([]*models.Project, error) {
	const op = "Postgres.ProjectRepository.GetAllByUserID"

	builder := sq.Select(projectsIDColumn, projectsTitleColumn, projectsDescriptionColumn, projectsUsersColumn, projectsAdminIDColumn, projectsCreatedAtColumn).
		From(projectsTableName).
		Where(sq.Expr("? = ANY(user_ids)", userID)).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var projects []*models.Project
	err = r.db.SelectContext(ctx, &projects, query, args...)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return projects, nil
}

func (r *ProjectRepository) Update(ctx context.Context, project *models.Project) error {
	const op = "Postgres.ProjectRepository.Update"

	builder := sq.Update(projectsTableName).
		Set("title", project.Title).
		Set("description", project.Description).
		Set("users", project.Users).
		Set("admin_id", project.AdminID).
		Where(sq.Eq{"id": project.ID}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	result, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("%s: %w", op, errorz.ErrProjectNotFound)
	}

	return nil
}

func (r *ProjectRepository) Delete(ctx context.Context, id string) error {
	const op = "Postgres.ProjectRepository.Delete"

	builder := sq.Delete(projectsIDColumn).
		Where(sq.Eq{projectsIDColumn: id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	result, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("%s: %w", op, errorz.ErrProjectNotFound)
	}

	return nil
}
