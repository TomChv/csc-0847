package server

import (
	"github.com/gin-gonic/gin"

	"github.com/TomChv/csc-0847/project_1/backend/db"
)

type Server struct {
	*gin.Engine
	c *db.Client
}

func New(c *db.Client) *Server {
	r := gin.Default()

	r.GET("/healthcheck", healthcheck)

	return &Server{r, c}
}
