package handlers

import (
	"net/http"

	"github.com/junaid/todo-service/internal/model"
	"github.com/junaid/todo-service/internal/repository"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	repo *repository.MemoryRepository
}

func NewTodoHandler(r *repository.MemoryRepository) *TodoHandler {
	return &TodoHandler{repo: r}
}

func (h *TodoHandler) RegisterRoutes(router *gin.Engine) {
	router.GET("/todos", h.GetAll)
	router.POST("/todos", h.Create)
	router.PUT("/todos/:id", h.Update)
	router.DELETE("/todos/:id", h.Delete)
	router.GET("/health", h.Health)
}

// GetAll godoc
// @Summary List all todos
// @Description Retrieve all todo items
// @Tags todos
// @Produce json
// @Success 200 {array} model.Todo
// @Router /todos [get]
func (h *TodoHandler) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, h.repo.GetAll())
}

// Create godoc
// @Summary Create a new todo
// @Description Add a new todo item
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body model.CreateTodoInput true "Todo item"
// @Success 201 {object} model.Todo
// @Failure 400 {object} map[string]string
// @Router /todos [post]
func (h *TodoHandler) Create(c *gin.Context) {
	var createTodoModel model.CreateTodoInput
	if err := c.ShouldBindJSON(&createTodoModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created := h.repo.Create(createTodoModel)
	c.JSON(http.StatusCreated, created)
}

// Update godoc
// @Summary Update a todo
// @Description Update an existing todo by ID
// @Tags todos
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Param todo body model.CreateTodoInput true "Updated todo data"
// @Success 200 {object} model.Todo
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /todos/{id} [put]
func (h *TodoHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var todo model.CreateTodoInput

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, ok := h.repo.Update(id, todo)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, updated)
}

// Delete godoc
// @Summary Delete a todo
// @Description Delete a todo item by ID
// @Tags todos
// @Param id path string true "Todo ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Router /todos/{id} [delete]
func (h *TodoHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if !h.repo.Delete(id) {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.Status(http.StatusNoContent)
}

// Health godoc
// @Summary Health check
// @Description Check API health status
// @Tags health
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func (h *TodoHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}
