package handlers

import (
	"database/sql"
	"strconv"

	"github.com/Milagrosgzmn/devops_todo_go.git/internal/constants"
	"github.com/Milagrosgzmn/devops_todo_go.git/internal/models"
	"github.com/Milagrosgzmn/devops_todo_go.git/internal/repository"
	"github.com/gin-gonic/gin"
)

type ItemHandler struct {
	repo repository.IRepository
}

func NewItemHandler(repo repository.IRepository) *ItemHandler {
	return &ItemHandler{
		repo: repo,
	}
}

func (h *ItemHandler) GetItems(c *gin.Context) {
	items, err := h.repo.GetAll()
	if err != nil {
		c.JSON(constants.StatusInternalServerError, gin.H{
			"error":   constants.ErrorBaseDatos,
			"details": err.Error(),
		})
		return
	}

	c.JSON(constants.StatusOK, gin.H{
		"message": constants.ItemsObtenidos,
		"data":    items,
	})
}

func (h *ItemHandler) GetItemByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(constants.StatusBadRequest, gin.H{
			"error": constants.IDInvalido,
		})
		return
	}

	item, err := h.repo.Get(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(constants.StatusNotFound, gin.H{
				"error": constants.ItemNoEncontrado,
			})
			return
		}
		c.JSON(constants.StatusInternalServerError, gin.H{
			"error":   constants.ErrorBaseDatos,
			"details": err.Error(),
		})
		return
	}

	c.JSON(constants.StatusOK, gin.H{
		"message": constants.ItemObtenido,
		"data":    item,
	})
}

func (h *ItemHandler) CreateItem(c *gin.Context) {
	var item models.TodoItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(constants.StatusBadRequest, gin.H{
			"error":   constants.CuerpoInvalido,
			"details": err.Error(),
		})
		return
	}

	// Validar el item usando el método Validate
	if err := item.Validate(); err != nil {
		c.JSON(constants.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	createdItem, err := h.repo.Create(item)
	if err != nil {
		c.JSON(constants.StatusInternalServerError, gin.H{
			"error":   constants.ErrorBaseDatos,
			"details": err.Error(),
		})
		return
	}

	c.JSON(constants.StatusCreated, gin.H{
		"message": constants.ItemCreado,
		"data":    createdItem,
	})
}

func (h *ItemHandler) UpdateItem(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(constants.StatusBadRequest, gin.H{
			"error": constants.IDInvalido,
		})
		return
	}

	var item models.TodoItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(constants.StatusBadRequest, gin.H{
			"error":   constants.CuerpoInvalido,
			"details": err.Error(),
		})
		return
	}

	item.ID = id

	// Validar el item usando el método Validate
	if err := item.Validate(); err != nil {
		c.JSON(constants.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.repo.Update(item)
	if err != nil {
		c.JSON(constants.StatusInternalServerError, gin.H{
			"error":   constants.ErrorBaseDatos,
			"details": err.Error(),
		})
		return
	}

	c.JSON(constants.StatusOK, gin.H{
		"message": constants.ItemActualizado,
		"data":    item,
	})
}

func (h *ItemHandler) DeleteItem(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(constants.StatusBadRequest, gin.H{
			"error": constants.IDInvalido,
		})
		return
	}

	err := h.repo.Delete(id)
	if err != nil {
		c.JSON(constants.StatusInternalServerError, gin.H{
			"error":   constants.ErrorBaseDatos,
			"details": err.Error(),
		})
		return
	}

	c.JSON(constants.StatusOK, gin.H{
		"message": constants.ItemEliminado,
	})
}

