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

func (h *Handler) GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Request received get data users")

		c.JSON(http.StatusOK, dao.Users)
	}
}

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
