package api

import (
	database "github.com/afarid/todo/db/sqlc"
	"github.com/gin-gonic/gin"
	"time"
)

type todoRequest struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Username    string    `json:"username" binding:"required"`
	Deadline    time.Time `json:"deadline" binding:"required"`
}

type todoResponse struct {
	ID          int32     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Username    string    `json:"username"`
	Deadline    time.Time `json:"deadline" binding:"required"`
}

func (s *Server) createTodo(c *gin.Context) {
	var request todoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	result, err := s.dbClient.CreateTodo(c, database.CreateTodoParams{
		Name:        request.Name,
		Description: request.Description,
		Username:    request.Username,
		Deadline:    request.Deadline,
	})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, result)
}

type listTodoRequest struct {
	Username string `json:"username" binding:"required"`
}

func (s *Server) listTodos(c *gin.Context) {
	todos, err := s.dbClient.ListTodos(c.Request.Context(), "afarid")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, todos)
}
