package todo

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	GetAll(ctx context.Context) ([]Todo, error)
	Create(ctx context.Context, title string) (*Todo, error)
	Update(ctx context.Context, id int, title string, completed bool) (*Todo, error)
	Delete(ctx context.Context, id int) error
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{db: db}
}

func (r *repository) GetAll(ctx context.Context) ([]Todo, error) {
	rows, err := r.db.Query(ctx, "SELECT id, title, completed, created_at FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var t Todo
		if err := rows.Scan(&t.ID, &t.Title, &t.Completed, &t.CreatedAt); err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	return todos, nil
}

func (r *repository) Create(ctx context.Context, title string) (*Todo, error) {
	todo := &Todo{
		Title:     title,
		Completed: false,
	}

	err := r.db.QueryRow(
		ctx,
		"INSERT INTO todos (title, completed) VALUES ($1, $2) RETURNING id, created_at",
		todo.Title, todo.Completed,
	).Scan(&todo.ID, &todo.CreatedAt)

	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (r *repository) Update(ctx context.Context, id int, title string, completed bool) (*Todo, error) {
	todo := &Todo{
		ID:        id,
		Title:     title,
		Completed: completed,
	}

	err := r.db.QueryRow(ctx, "UPDATE todos SET title = $1, completed = $2 WHERE id = $3 RETURNING id, created_at", title, completed, id).Scan(&todo.ID, &todo.CreatedAt)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, "DELETE FROM todos WHERE id=$1", id)
	return err
}
