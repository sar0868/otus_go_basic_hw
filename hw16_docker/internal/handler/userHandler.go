package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sar0868/otus_go_basic_hw/hw16_docker/internal/app"
	"github.com/sar0868/otus_go_basic_hw/hw16_docker/internal/repository"
	"github.com/sar0868/otus_go_basic_hw/hw16_docker/internal/service"
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
		user, err := service.AddUser(app.Ctx, app.Repo, newUser)
		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{
				"error": fmt.Sprintf("Error add user: %v", err),
			})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

// Delete User by name
// @Summary delete User by name
// @Tags deleteUserByName
// @Accept			json
// @Produce		json
// @Param name query string true "string valid"
// @Success 200 {string} string "Delete user by name"
// @Failure 400 {string} string "Error"
// @Router /del_user [delete].
func (h *Handler) DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.DefaultQuery("name", "")
		err := service.DeleteUser(app.Ctx, app.Repo, name)
		if err != nil {
			msg := fmt.Sprintf("Error delete user by name %s: %s", name, err)
			c.JSON(http.StatusNotAcceptable, gin.H{
				"message": msg,
			})
			return
		}
		msg := fmt.Sprintf("User name %s delete", name)
		c.JSON(http.StatusOK, gin.H{
			"message": msg,
		})
	}
}

// Update User name
// @Summary updateUser
// @Tags updateUser
// @Accept			json
// @Produce		json
// @Param input body repository.UserUpdateParams true "Модель которую принимает метод"
// @Success 200 {string} string "Update user name"
// @Failure 400 {string} string "Error"
// @Router /edit_user [post].
func (h *Handler) UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var updateUser repository.UserUpdateParams
		if err := c.ShouldBindJSON(&updateUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request",
			})
			return
		}
		user, err := service.UserUpdate(app.Ctx, app.Repo, updateUser)
		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{
				"error": fmt.Sprintf("Error update user: %v", err),
			})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

// Get users statistic
// @Summary get users statistic
// @Tags getUsersStatistic
// @Accept			json
// @Produce		json
// @Success 200 {array} repository.UsersStatisticRow
// @Failure 400 {string} string "Error"
// @Router /statistic [get].
func (h *Handler) Statistic() gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := service.UsersStatistic(app.Ctx, app.Repo)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		c.JSON(http.StatusOK, users)
	}
}
