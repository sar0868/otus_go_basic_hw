package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sar0868/otus_go_basic_hw/hw15_go_sql/internal/app"
	"github.com/sar0868/otus_go_basic_hw/hw15_go_sql/internal/service"
)

// Get Orders
// @Summary get orders
// @Tags getOrders
// @Accept			json
// @Produce		json
// @Success 200 {string} string "Get orders"
// @Failure 400 {string} string "Error"
// @Router /orders [get].
func (h *Handler) GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		orders, err := service.Orders(app.Ctx, app.Repo)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		c.JSON(http.StatusOK, orders)
	}
}
