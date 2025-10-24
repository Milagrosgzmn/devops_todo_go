package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/Milagrosgzmn/devops_todo_go.git/internal/constants"
	"github.com/Milagrosgzmn/devops_todo_go.git/internal/models"
	"github.com/go-sql-driver/mysql"
)

type ItemMySqlRepository struct {
	db *sql.DB
}

func NewItemMySqlRepository(db *sql.DB) *ItemMySqlRepository {
	return &ItemMySqlRepository{db: db}
}

func handleMySQLError(err error) error {
	if err == nil {
		return nil
	}

	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		// Error 1265: Data truncated (errores de ENUM principalmente)
		if mysqlErr.Number == 1265 {
			if strings.Contains(mysqlErr.Message, "state") {
				return errors.New(constants.EstadoInvalido)
			}
		}
	}

	return err
}

func (r *ItemMySqlRepository) GetAll() ([]models.TodoItem, error) {
	rows, err := r.db.Query("SELECT id, title, description, state, created_at, updated_at, deleted_at FROM todo_items WHERE deleted_at IS NULL")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.TodoItem
	for rows.Next() {
		var item models.TodoItem
		if err := rows.Scan(&item.ID, &item.Title, &item.Description, &item.State, &item.CreatedAt, &item.UpdatedAt, &item.DeletedAt); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *ItemMySqlRepository) Get(id int) (models.TodoItem, error) {
	var item models.TodoItem
	err := r.db.QueryRow("SELECT id, title, description, state, created_at, updated_at, deleted_at FROM todo_items WHERE id = ? AND deleted_at IS NULL", id).Scan(&item.ID, &item.Title, &item.Description, &item.State, &item.CreatedAt, &item.UpdatedAt, &item.DeletedAt)
	if err != nil {
		return models.TodoItem{}, err
	}
	return item, nil
}

func (r *ItemMySqlRepository) Create(item models.TodoItem) (models.TodoItem, error) {
	result, err := r.db.Exec("INSERT INTO todo_items (title, description, state) VALUES (?, ?, ?)", item.Title, item.Description, item.State)
	if err != nil {
		return models.TodoItem{}, handleMySQLError(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return models.TodoItem{}, err
	}
	item.ID = fmt.Sprintf("%d", id)
	return item, nil
}

func (r *ItemMySqlRepository) Update(item models.TodoItem) error {
	_, err := r.db.Exec("UPDATE todo_items SET title = ?, description = ?, state = ? WHERE id = ?", item.Title, item.Description, item.State, item.ID)
	return handleMySQLError(err)
}

func (r *ItemMySqlRepository) Delete(id string) error {
	_, err := r.db.Exec("UPDATE todo_items SET deleted_at = NOW() WHERE id = ?", id)
	return err
}
