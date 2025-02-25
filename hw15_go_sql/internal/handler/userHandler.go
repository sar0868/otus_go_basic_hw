package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sar0868/otus_go_basic_hw/hw15_go_sql/internal/app"
	"github.com/sar0868/otus_go_basic_hw/hw15_go_sql/internal/repository"
	"github.com/sar0868/otus_go_basic_hw/hw15_go_sql/internal/service"
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
		users, err := service.Users(app.Ctx, app.Repo)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		c.JSON(http.StatusOK, users)
	}
}

// Get User by parameters
// @Summary get User by parameters
// @Tags getUserByParameters
// @Accept			json
// @Produce		json
// @Param id query string false "string valid"
// @Param name query string false "string valid"
// @Success 200 {string} string "Get user by parameters"
// @Failure 400 {string} string "Error"
// @Router /user [get].
func (h *Handler) GetUserByParameter() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params []service.ParamUser
		idStr := c.DefaultQuery("id", "")
		if idStr != "" {
			param := service.ParamUser{}
			param.Param = "id"
			param.Value = idStr
			params = append(params, param)
		}
		name := c.DefaultQuery("name", "")
		if name != "" {
			param := service.ParamUser{}
			param.Param = "name"
			param.Value = name
			params = append(params, param)
		}
		user, err := service.GetUserByParam(app.Ctx, app.Repo, params)
		if err != nil {
			msg := fmt.Sprintf("Error get user by parameters: %s", err)
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
// @Param input body repository.UserAddParams true "Модель которую принимает метод"
// @Success 200 {string} string "Added user"
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
