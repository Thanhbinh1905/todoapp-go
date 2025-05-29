package todo

import "context"

type Service interface {
	GetTodos(ctx context.Context) ([]Todo, error)
	CreateTodo(ctx context.Context, title string) (*Todo, error)
	UpdateTodo(ctx context.Context, id int, title string, completed bool) (*Todo, error)
	DeleteTodo(ctx context.Context, id int) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetTodos(ctx context.Context) ([]Todo, error) {
	return s.repo.GetAll(ctx)
}

func (s *service) CreateTodo(ctx context.Context, title string) (*Todo, error) {
	return s.repo.Create(ctx, title)
}

func (s *service) UpdateTodo(ctx context.Context, id int, title string, completed bool) (*Todo, error) {
	return s.repo.Update(ctx, id, title, completed)
}

func (s *service) DeleteTodo(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
