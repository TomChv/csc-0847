package server

import "github.com/gin-gonic/gin"

type Server struct {
	*gin.Engine
}

func New() *Server {
	r := gin.Default()

	return &Server{r}
}
