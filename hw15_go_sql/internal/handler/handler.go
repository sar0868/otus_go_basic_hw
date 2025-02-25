package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func InitHandler(api *gin.Engine, h *Handler) {
	api.GET("/users", h.GetUsers())
	api.GET("/user", h.GetUserByParameter())
	api.POST("/add_user", h.AddUser())
	api.DELETE("/del_user", h.DeleteUser())
}
