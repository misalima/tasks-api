package postgres

import (
	"context"
	"fmt"
	"tasks-api/internal/core/domain"
	"tasks-api/internal/core/interfaces/irepository"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ irepository.TaskLoader = (*TaskRepository)(nil)

type TaskRepository struct {
	pool *pgxpool.Pool
}

func NewTaskRepository(pool *pgxpool.Pool) *TaskRepository {
	return &TaskRepository{pool: pool}
}

func (t *TaskRepository) InsertTask(ctx context.Context, task domain.Task) (*domain.Task, error){
	sql := "INSERT INTO tasks (title, description) VALUES ($1, $2) RETURNING id"

	err := t.pool.QueryRow(ctx, sql, 
		task.Title, 
		task.Description, 
	).Scan(&task.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to save task: %w", err)
	}

	return &task, nil

}

func (t *TaskRepository) FetchTasks(ctx context.Context) ([]domain.Task, error){
	panic("Not Implemented")
}
func (t *TaskRepository) GetTaskByID(ctx context.Context, id int) (*domain.Task, error){
	query := `
		SELECT id, title, description, status
		FROM tasks
		WHERE id = $1
	`
	row := t.pool.QueryRow(ctx, query, id)

	var task domain.Task

	err := row.Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Status,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, domain.ErrTaskNotFound
		}
		return nil, err
	}

	return &task, nil
}

func (t *TaskRepository) UpdateTask(ctx context.Context, id int) error{
	panic("Not Implemented")
}
func (t *TaskRepository) DeleteTask(ctx context.Context, id int) error{
	panic("Not Implemented")
}