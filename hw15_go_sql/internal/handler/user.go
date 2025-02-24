package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sar0868/otus_go_basic_hw/hw15_go_sql/internal/app"
	"github.com/sar0868/otus_go_basic_hw/hw15_go_sql/internal/repository"
)

// Get Users
// @Summary get users
// @Tags getUsers
// @Accept			json
// @Produce		json
// @Success 200 {string} string "Get users"
// @Failure 400 {string} string "Error"
// @Router /users [get].
func (h *Handler) GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Request received get data users")
		users, err := repository.Querier.Users(app.Repo, app.Ctx)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		c.JSON(http.StatusOK, users)
	}
}

// Get User by id
// @Summary get User by ID
// @Tags getUserById
// @Accept			json
// @Produce		json
// @Param id query string false "string valid"
// @Success 200 {string} string "Get user by id"
// @Failure 400 {string} string "Error"
// @Router /user [get].
func (h *Handler) GetUserByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.DefaultQuery("id", "")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Error convert string to int")
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid note id",
			})
			return
		}

		user, errGetUser := repository.Querier.UserByID(app.Repo, app.Ctx, id)
		if errGetUser != nil {
			msg := fmt.Sprintf("Don't found data for id=%d", id)
			c.JSON(http.StatusNotAcceptable, gin.H{
				"message": msg,
			})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

// Add User
// @Summary addUser
// @Tags addUser
// @Accept			json
// @Produce		json
// @Param input body models.AddUser true "Модель которую принимает метод"
// @Success 200 {string} string "Get user by id"
// @Failure 400 {string} string "Error"
// @Router /add_user [post].
func (h *Handler) AddUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser repository.UserAddParams
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request",
			})
			return
		}
		user, err := repository.Querier.UserAdd(app.Repo, app.Ctx, newUser)
		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{
				"error": fmt.Sprintf("Error add user: %v", err),
			})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}
