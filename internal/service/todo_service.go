package service

import (
	"context"
	"grpc-crud-app/api/proto/todo"
	"grpc-crud-app/internal/repository"
)

type TodoService struct {
	todo.UnimplementedTodoServiceServer
	repo *repository.TodoRepository
}

func NewTodoService(repo *repository.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) CreateTodo(ctx context.Context, req *todo.Todo) (*todo.Todo, error) {
	t := repository.Todo{
		Title:       req.Title,
		Description: req.Description,
	}
	createdTodo := s.repo.Create(t)
	return &todo.Todo{
		Id:          createdTodo.ID,
		Title:       createdTodo.Title,
		Description: createdTodo.Description,
	}, nil
}

func (s *TodoService) GetTodo(ctx context.Context, req *todo.TodoId) (*todo.Todo, error) {
	todoItem, found := s.repo.Get(req.Id)
	if !found {
		return nil, nil // In production, you should return a proper error
	}
	return &todo.Todo{
		Id:          todoItem.ID,
		Title:       todoItem.Title,
		Description: todoItem.Description,
	}, nil
}

func (s *TodoService) UpdateTodo(ctx context.Context, req *todo.Todo) (*todo.Todo, error) {
	t := repository.Todo{
		ID:          req.Id,
		Title:       req.Title,
		Description: req.Description,
	}
	updatedTodo, found := s.repo.Update(t)
	if !found {
		return nil, nil // In production, you should return a proper error
	}
	return &todo.Todo{
		Id:          updatedTodo.ID,
		Title:       updatedTodo.Title,
		Description: updatedTodo.Description,
	}, nil
}

func (s *TodoService) DeleteTodo(ctx context.Context, req *todo.TodoId) (*todo.Empty, error) {
	deleted := s.repo.Delete(req.Id)
	if !deleted {
		return nil, nil // In production, you should return a proper error
	}
	return &todo.Empty{}, nil
}

func (s *TodoService) ListTodos(req *todo.Empty, stream todo.TodoService_ListTodosServer) error {
	todos := s.repo.List()
	for _, t := range todos {
		stream.Send(&todo.Todo{
			Id:          t.ID,
			Title:       t.Title,
			Description: t.Description,
		})
	}
	return nil
}
