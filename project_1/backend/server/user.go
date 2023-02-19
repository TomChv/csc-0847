package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"

	"github.com/TomChv/csc-0847/project_1/backend/db"
	"github.com/TomChv/csc-0847/project_1/backend/db/model"
)

// UserController defines a list of method that can be binds to gin-gonic
// endpoints.
type UserController struct {
	db *db.Client
}

// BindUserController links user endpoints to the given router.
func BindUserController(dbClient *db.Client, group gin.RouterGroup) {
	users := group.Group("/users")

	controller := UserController{db: dbClient}

	users.GET("", controller.ListUsers)
	users.GET("/:userID", controller.GetUser)
	users.POST("", controller.CreateUser)
	users.PUT("/:userID", controller.UpdateUser)
	users.DELETE("/:userID", controller.DeleteUser)
}

// ListUsers returns all users stored in the database.
func (u *UserController) ListUsers(c *gin.Context) {
	users, err := model.ListUser(c, u.db)
	if err != nil {
		NewHTTPError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUser searches for a user stored in the database by its ID and returns it.
func (u *UserController) GetUser(c *gin.Context) {
	userID := c.Param("userID")

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		NewHTTPError(c, http.StatusBadRequest, err)
		return
	}

	user, err := model.GetUser(c, u.db, userUUID)
	if err != nil {
		NewHTTPError(c, http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

// CreateUser inserts a new user in the database.
func (u *UserController) CreateUser(c *gin.Context) {
	var data *model.CreateUserDTO
	if err := c.BindJSON(&data); err != nil {
		NewHTTPError(c, http.StatusBadRequest, err)
		return
	}

	newUser, err := model.CreateUser(c, u.db, data)
	if err != nil {
		NewHTTPError(c, http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusCreated, newUser)
}

// UpdateUser searches a user in the database by its id and updates it
// with given fields.
func (u *UserController) UpdateUser(c *gin.Context) {
	userID := c.Param("userID")

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		NewHTTPError(c, http.StatusBadRequest, err)
		return
	}

	var data *model.UpdateUserDTO
	if err := c.BindJSON(&data); err != nil {
		NewHTTPError(c, http.StatusBadRequest, err)
		return
	}

	updatedUser, err := model.UpdateUser(c, u.db, userUUID, data)
	if err != nil {
		NewHTTPError(c, http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, updatedUser)
}

// DeleteUser searches for a user stored in the database by its ID and deletes it.
func (u *UserController) DeleteUser(c *gin.Context) {
	userID := c.Param("userID")

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		NewHTTPError(c, http.StatusBadRequest, err)
		return
	}

	err = model.DeleteUser(c, u.db, userUUID)
	if err != nil {
		NewHTTPError(c, http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("user %s successfully deleted", userID),
	})
}
