package postgres

import (
	"context"
	"fmt"
	"project-service/internal/errorz"
	"project-service/internal/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

const (
	tasksTableName = "tasks"

	tasksIDColumn          = "id"
	tasksTitleColumn       = "title"
	tasksDescriptionColumn = "description"
	tasksStatusColumn      = "status"
	tasksProjectIDColumn   = "project_id"
	tasksExecutorIDColumn  = "executor_id"
	tasksDeadlineColumn    = "deadline"
	tasksCreatedAtColumn   = "created_at"
	tasksUpdatedAtColumn   = "updated_at"
)

type TaskRepository struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) *TaskRepository { return &TaskRepository{db: db} }

func (r *TaskRepository) Create(ctx context.Context, task *models.Task) (string, error) {
	const op = "Postgres.ProjectRepository.Create"

	builder := sq.Insert(tasksTableName).
		Columns(tasksTitleColumn, tasksDescriptionColumn, tasksStatusColumn, tasksProjectIDColumn, tasksExecutorIDColumn, tasksDeadlineColumn, tasksCreatedAtColumn).
		Values(task.Title, task.Description, task.Status, task.ProjectId, task.ExecutorId, task.Deadline, sq.Expr("NOW()")).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	var id string
	if err = r.db.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return id, nil
}

func (r *TaskRepository) GetByID(ctx context.Context, id string) (*models.Task, error) {
	const op = "Postgres.ProjectRepository.GetByID"

	builder := sq.Select(tasksIDColumn, tasksTitleColumn, tasksDescriptionColumn, tasksStatusColumn, tasksProjectIDColumn, tasksExecutorIDColumn, tasksDeadlineColumn, tasksCreatedAtColumn, tasksUpdatedAtColumn).
		From(tasksTableName).
		Where(sq.Eq{tasksIDColumn: id}).
		Limit(1).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var task models.Task
	if err = r.db.GetContext(ctx, &task, query, args...); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &task, nil
}

func (r *TaskRepository) GetAllByProjectID(ctx context.Context, projectID string) ([]*models.Task, error) {
	const op = "Postgres.ProjectRepository.GetAllByProjectID"

	builder := sq.Select(tasksIDColumn, tasksTitleColumn, tasksDescriptionColumn, tasksStatusColumn, tasksProjectIDColumn, tasksExecutorIDColumn, tasksDeadlineColumn, tasksCreatedAtColumn, tasksUpdatedAtColumn).
		From(tasksTableName).
		Where(sq.Eq{tasksProjectIDColumn: projectID}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var tasks []*models.Task
	if err = r.db.SelectContext(ctx, &tasks, query, args...); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return tasks, nil
}

func (r *TaskRepository) Update(ctx context.Context, task *models.Task) error {
	const op = "Postgres.ProjectRepository.Update"

	builder := sq.Update(tasksTableName).
		Set(tasksTitleColumn, task.Title).
		Set(tasksDescriptionColumn, task.Description).
		Set(tasksStatusColumn, task.Status).
		Set(tasksExecutorIDColumn, task.ExecutorId).
		Set(tasksDeadlineColumn, task.Deadline).
		Set(tasksUpdatedAtColumn, sq.Expr("NOW()")).
		Where(sq.Eq{tasksIDColumn: task.ID}).
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
		return fmt.Errorf("%s: %w", op, errorz.ErrTaskNotFound)
	}
	return nil
}

func (r *TaskRepository) Delete(ctx context.Context, id string) error {
	const op = "Postgres.ProjectRepository.Delete"

	builder := sq.Delete(tasksTableName).
		Where(sq.Eq{tasksIDColumn: id}).
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
		return fmt.Errorf("%s: %w", op, errorz.ErrTaskNotFound)
	}

	return nil
}
