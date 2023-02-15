package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func healthcheck(c *gin.Context) {
	c.Status(http.StatusOK)
}
