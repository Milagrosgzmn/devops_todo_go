package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/Milagrosgzmn/devops_todo_go.git/internal/models"
)

// MockRepository implementa IRepository usando un map en memoria para testing
type MockRepository struct {
	items        map[string]models.TodoItem
	nextID       int
	mu           sync.RWMutex
	simulateError bool
}

func NewMockRepository() *MockRepository {
	return &MockRepository{
		items:        make(map[string]models.TodoItem),
		nextID:       1,
		simulateError: false,
	}
}

// SimulateError activa la simulación de errores de base de datos
func (r *MockRepository) SimulateError(enable bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.simulateError = enable
}

func (r *MockRepository) GetAll() ([]models.TodoItem, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.simulateError {
		return nil, errors.New("simulated database error")
	}

	items := make([]models.TodoItem, 0, len(r.items))
	for _, item := range r.items {
		if item.DeletedAt == nil {
			items = append(items, item)
		}
	}
	return items, nil
}

func (r *MockRepository) Get(id int) (models.TodoItem, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.simulateError {
		return models.TodoItem{}, errors.New("simulated database error")
	}

	idStr := strconv.Itoa(id)
	item, exists := r.items[idStr]
	if !exists || item.DeletedAt != nil {
		return models.TodoItem{}, sql.ErrNoRows
	}
	return item, nil
}

func (r *MockRepository) Create(item models.TodoItem) (models.TodoItem, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.simulateError {
		return models.TodoItem{}, errors.New("simulated database error")
	}

	item.ID = fmt.Sprintf("%d", r.nextID)
	r.nextID++
	r.items[item.ID] = item
	return item, nil
}

func (r *MockRepository) Update(item models.TodoItem) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.simulateError {
		return errors.New("simulated database error")
	}

	if _, exists := r.items[item.ID]; !exists {
		return sql.ErrNoRows
	}
	r.items[item.ID] = item
	return nil
}

func (r *MockRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.simulateError {
		return errors.New("simulated database error")
	}

	item, exists := r.items[id]
	if !exists {
		return sql.ErrNoRows
	}
	deleted := "deleted"
	item.DeletedAt = &deleted
	r.items[id] = item
	return nil
}
