package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Milagrosgzmn/devops_todo_go.git/internal/constants"
	"github.com/Milagrosgzmn/devops_todo_go.git/internal/models"
	"github.com/Milagrosgzmn/devops_todo_go.git/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() (*gin.Engine, *ItemHandler) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	repo := repository.NewMockRepository()
	handler := NewItemHandler(repo)
	return router, handler
}

func TestGetItems(t *testing.T) {
	router, handler := setupRouter()
	router.GET("/items", handler.GetItems)

	// Crear items de prueba
	handler.repo.Create(models.TodoItem{Title: "Task 1", State: constants.StatePending, Description: "Test 1"})
	handler.repo.Create(models.TodoItem{Title: "Task 2", State: constants.StateCompleted, Description: "Test 2"})

	req, _ := http.NewRequest("GET", "/items", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, constants.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, constants.ItemsObtenidos, response["message"])
	assert.NotNil(t, response["data"])
}

func TestGetItemByID_Success(t *testing.T) {
	router, handler := setupRouter()
	router.GET("/items/:id", handler.GetItemByID)

	// Crear item de prueba
	handler.repo.Create(models.TodoItem{Title: "Task 1", State: constants.StatePending, Description: "Test 1"})

	req, _ := http.NewRequest("GET", "/items/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, constants.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, constants.ItemObtenido, response["message"])
	assert.NotNil(t, response["data"])
}

func TestGetItemByID_NotFound(t *testing.T) {
	router, handler := setupRouter()
	router.GET("/items/:id", handler.GetItemByID)

	req, _ := http.NewRequest("GET", "/items/999", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, constants.StatusNotFound, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, constants.ItemNoEncontrado, response["error"])
}

func TestGetItemByID_InvalidID(t *testing.T) {
	router, handler := setupRouter()
	router.GET("/items/:id", handler.GetItemByID)

	req, _ := http.NewRequest("GET", "/items/invalid", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, constants.StatusBadRequest, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, constants.IDInvalido, response["error"])
}

func TestCreateItem_Success(t *testing.T) {
	router, handler := setupRouter()
	router.POST("/items", handler.CreateItem)

	item := models.TodoItem{
		Title:       "New Task",
		Description: "Test Description",
		State:       constants.StatePending,
	}
	jsonData, _ := json.Marshal(item)

	req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, constants.StatusCreated, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, constants.ItemCreado, response["message"])
	assert.NotNil(t, response["data"])
}

func TestCreateItem_MissingFields(t *testing.T) {
	router, handler := setupRouter()
	router.POST("/items", handler.CreateItem)

	item := models.TodoItem{
		Description: "Test Description",
	}
	jsonData, _ := json.Marshal(item)

	req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, constants.StatusBadRequest, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	// Ahora devuelve el mensaje específico del campo que falta
	assert.Contains(t, response["error"], "requerido")
}

func TestCreateItem_InvalidJSON(t *testing.T) {
	router, handler := setupRouter()
	router.POST("/items", handler.CreateItem)

	req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, constants.StatusBadRequest, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, constants.CuerpoInvalido, response["error"])
}

func TestUpdateItem_Success(t *testing.T) {
	router, handler := setupRouter()
	router.PUT("/items/:id", handler.UpdateItem)

	// Crear item inicial
	handler.repo.Create(models.TodoItem{Title: "Task 1", State: constants.StatePending, Description: "Test 1"})

	updatedItem := models.TodoItem{
		Title:       "Updated Task",
		Description: "Updated Description",
		State:       constants.StateCompleted,
	}
	jsonData, _ := json.Marshal(updatedItem)

	req, _ := http.NewRequest("PUT", "/items/1", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, constants.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, constants.ItemActualizado, response["message"])
}

func TestUpdateItem_MissingFields(t *testing.T) {
	router, handler := setupRouter()
	router.PUT("/items/:id", handler.UpdateItem)

	item := models.TodoItem{
		Description: "Only description",
	}
	jsonData, _ := json.Marshal(item)

	req, _ := http.NewRequest("PUT", "/items/1", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, constants.StatusBadRequest, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	// Ahora devuelve el mensaje específico del campo que falta
	assert.Contains(t, response["error"], "requerido")
}

func TestDeleteItem_Success(t *testing.T) {
	router, handler := setupRouter()
	router.DELETE("/items/:id", handler.DeleteItem)

	// Crear item de prueba
	handler.repo.Create(models.TodoItem{Title: "Task 1", State: constants.StatePending, Description: "Test 1"})

	req, _ := http.NewRequest("DELETE", "/items/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, constants.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, constants.ItemEliminado, response["message"])
}

// Tests de errores de base de datos

func TestGetItems_DatabaseError(t *testing.T) {
	router, handler := setupRouter()
	router.GET("/items", handler.GetItems)

	// Simular error de BD
	mockRepo := handler.repo.(*repository.MockRepository)
	mockRepo.SimulateError(true)

	req, _ := http.NewRequest("GET", "/items", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, constants.StatusInternalServerError, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, constants.ErrorBaseDatos, response["error"])
	assert.NotNil(t, response["details"])
}

func TestGetItemByID_DatabaseError(t *testing.T) {
	router, handler := setupRouter()
	router.GET("/items/:id", handler.GetItemByID)

	// Simular error de BD
	mockRepo := handler.repo.(*repository.MockRepository)
	mockRepo.SimulateError(true)

	req, _ := http.NewRequest("GET", "/items/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, constants.StatusInternalServerError, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, constants.ErrorBaseDatos, response["error"])
	assert.NotNil(t, response["details"])
}

func TestCreateItem_DatabaseError(t *testing.T) {
	router, handler := setupRouter()
	router.POST("/items", handler.CreateItem)

	// Simular error de BD
	mockRepo := handler.repo.(*repository.MockRepository)
	mockRepo.SimulateError(true)

	item := models.TodoItem{
		Title:       "New Task",
		Description: "Test Description",
		State:       constants.StatePending,
	}
	jsonData, _ := json.Marshal(item)

	req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, constants.StatusInternalServerError, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, constants.ErrorBaseDatos, response["error"])
	assert.NotNil(t, response["details"])
}

func TestUpdateItem_DatabaseError(t *testing.T) {
	router, handler := setupRouter()
	router.PUT("/items/:id", handler.UpdateItem)

	// Simular error de BD
	mockRepo := handler.repo.(*repository.MockRepository)
	mockRepo.SimulateError(true)

	updatedItem := models.TodoItem{
		Title:       "Updated Task",
		Description: "Updated Description",
		State:       constants.StateCompleted,
	}
	jsonData, _ := json.Marshal(updatedItem)

	req, _ := http.NewRequest("PUT", "/items/1", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, constants.StatusInternalServerError, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, constants.ErrorBaseDatos, response["error"])
	assert.NotNil(t, response["details"])
}

func TestDeleteItem_DatabaseError(t *testing.T) {
	router, handler := setupRouter()
	router.DELETE("/items/:id", handler.DeleteItem)

	// Simular error de BD
	mockRepo := handler.repo.(*repository.MockRepository)
	mockRepo.SimulateError(true)

	req, _ := http.NewRequest("DELETE", "/items/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, constants.StatusInternalServerError, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, constants.ErrorBaseDatos, response["error"])
	assert.NotNil(t, response["details"])
}

// Tests de validación de estados

func TestCreateItem_InvalidState(t *testing.T) {
	router, handler := setupRouter()
	router.POST("/items", handler.CreateItem)

	item := models.TodoItem{
		Title:       "New Task",
		Description: "Test Description",
		State:       "invalid_state",
	}
	jsonData, _ := json.Marshal(item)

	req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, constants.StatusBadRequest, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, constants.EstadoInvalido, response["error"])
}

func TestCreateItem_ValidStates(t *testing.T) {
	validStates := []string{constants.StatePending, constants.StateInProgress, constants.StateCompleted}

	for _, state := range validStates {
		t.Run("state_"+state, func(t *testing.T) {
			router, handler := setupRouter()
			router.POST("/items", handler.CreateItem)

			item := models.TodoItem{
				Title:       "New Task",
				Description: "Test Description",
				State:       state,
			}
			jsonData, _ := json.Marshal(item)

			req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, constants.StatusCreated, w.Code)

			var response map[string]interface{}
			json.Unmarshal(w.Body.Bytes(), &response)
			assert.Equal(t, constants.ItemCreado, response["message"])
		})
	}
}

func TestUpdateItem_InvalidState(t *testing.T) {
	router, handler := setupRouter()
	router.PUT("/items/:id", handler.UpdateItem)

	// Crear item inicial
	handler.repo.Create(models.TodoItem{Title: "Task 1", State: constants.StatePending, Description: "Test 1"})

	updatedItem := models.TodoItem{
		Title:       "Updated Task",
		Description: "Updated Description",
		State:       "done",
	}
	jsonData, _ := json.Marshal(updatedItem)

	req, _ := http.NewRequest("PUT", "/items/1", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, constants.StatusBadRequest, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, constants.EstadoInvalido, response["error"])
}
