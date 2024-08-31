package repository

import "sync"

type Todo struct {
	ID          int32
	Title       string
	Description string
}

type TodoRepository struct {
	mu     sync.Mutex
	todos  map[int32]Todo
	nextID int32
}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{
		todos:  make(map[int32]Todo),
		nextID: 1,
	}
}

func (r *TodoRepository) Create(t Todo) Todo {
	r.mu.Lock()
	defer r.mu.Unlock()
	t.ID = r.nextID
	r.todos[t.ID] = t
	r.nextID++
	return t
}

func (r *TodoRepository) Get(id int32) (Todo, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	t, found := r.todos[id]
	return t, found
}

func (r *TodoRepository) Update(t Todo) (Todo, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, found := r.todos[t.ID]
	if !found {
		return Todo{}, false
	}
	r.todos[t.ID] = t
	return t, true
}

func (r *TodoRepository) Delete(id int32) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, found := r.todos[id]
	if !found {
		return false
	}
	delete(r.todos, id)
	return true
}

func (r *TodoRepository) List() []Todo {
	r.mu.Lock()
	defer r.mu.Unlock()
	todos := make([]Todo, 0, len(r.todos))
	for _, t := range r.todos {
		todos = append(todos, t)
	}
	return todos
}
