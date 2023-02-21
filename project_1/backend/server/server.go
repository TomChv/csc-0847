package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/TomChv/csc-0847/project_1/backend/db"
)

type Server struct {
	*gin.Engine
	c *db.Client
}

func New(c *db.Client) *Server {
	r := gin.Default()

	loadConfig()

	corsConfig := cors.Config{
		AllowAllOrigins:  allowAllOrigin,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: false,
	}

	if !corsConfig.AllowAllOrigins {
		corsConfig.AllowOrigins = corsOrigins
	}

	r.Use(cors.New(corsConfig))

	r.GET("/healthcheck", healthcheck)

	BindUserController(c, r.RouterGroup)

	return &Server{r, c}
}
