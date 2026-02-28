package repository

import (
	"sync"

	"github.com/google/uuid"
	"github.com/junaid/todo-service/internal/model"
)

type MemoryRepository struct {
	mu    sync.RWMutex
	store map[string]model.Todo
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		store: make(map[string]model.Todo),
	}
}

func (r *MemoryRepository) GetAll() []model.Todo {
	r.mu.RLock()
	defer r.mu.RUnlock()

	todos := []model.Todo{}
	for _, v := range r.store {
		todos = append(todos, v)
	}
	return todos
}

func (r *MemoryRepository) Create(createTodoModel model.CreateTodoInput) model.Todo {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Convert input to model.Todo and generate ID
	todo := model.Todo{
		ID:        uuid.New().String(), // auto-generated
		Title:     createTodoModel.Title,
		Completed: createTodoModel.Completed,
	}

	r.store[todo.ID] = todo
	return todo
}

func (r *MemoryRepository) Update(id string, input model.CreateTodoInput) (model.Todo, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	todo, exists := r.store[id]
	if !exists {
		return model.Todo{}, false
	}

	// Only update fields that are non-nil
	todo.Title = input.Title
	todo.Completed = input.Completed

	r.store[id] = todo
	return todo, true
}

func (r *MemoryRepository) Delete(id string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, exists := r.store[id]
	if !exists {
		return false
	}
	delete(r.store, id)
	return true
}
