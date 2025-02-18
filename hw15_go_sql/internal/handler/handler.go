package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sar0868/otus_go_basic_hw/hw15_go_sql/internal/models"
	"github.com/sar0868/otus_go_basic_hw/hw15_go_sql/server/dao"
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

// Get Users
// @Summary get users
// @Tags getusers
// @Accept			json
// @Produce		json
// @Success 200 {string} string "Get users"
// @Failure 400 {string} string "Error"
// @Router /users [get].
func (h *Handler) GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Request received get data users")

		c.JSON(http.StatusOK, dao.Users)
	}
}

// Get User by id
// @Summary getUserById
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
		user, result := dao.GetUser(id)
		if !result {
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
// @Param input body models.User true "Модель которую принимает метод"
// @Success 200 {string} string "Get user by id"
// @Failure 400 {string} string "Error"
// @Router /add_user [post].
func (h *Handler) AddUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser models.User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request",
			})
			return
		}
		user, err := dao.CreateUser(newUser.Name, newUser.Age, newUser.Address)
		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{
				"error": fmt.Sprintf("Error add user: %v", err),
			})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

func InitHandler(api *gin.Engine, h *Handler) {
	api.GET("/users", h.GetUsers())
	api.GET("/user", h.GetUserByID())
	api.POST("/add_user", h.AddUser())
}
