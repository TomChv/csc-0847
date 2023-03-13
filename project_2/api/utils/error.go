package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// HTTPError is a wrapper around gin error to format error.
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

// NewHTTPError creates a new gin error.
func NewHTTPError(c *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}

	fmt.Println("error: ", err.Error())
	c.JSON(status, er)
}
