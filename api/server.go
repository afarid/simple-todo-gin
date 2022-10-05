package api

import (
	database "github.com/afarid/todo/db/sqlc"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	dbClient *database.Queries
}

func NewServer(dbClient *database.Queries) *Server {
	return &Server{
		dbClient: dbClient,
	}
}

func (s *Server) Run() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.POST("/todos", s.createTodo)
	r.GET("/todos", s.listTodos)
	r.Run("0.0.0.0:8000")
}
