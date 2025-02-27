package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sar0868/otus_go_basic_hw/hw15_go_sql/internal/app"
	"github.com/sar0868/otus_go_basic_hw/hw15_go_sql/internal/service"
	"github.com/sar0868/otus_go_basic_hw/hw15_go_sql/internal/trx"
)

type CreateOrderParams struct {
	User     string  `db:"name" json:"user"`
	Name     string  `db:"name" json:"product"`
	Quantity float64 `db:"quantity" json:"quantity"`
}

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

// Create order
// @Summary Order Create
// @Tags OrderCreate
// @Accept			json
// @Produce		json
// @Param input body CreateOrderParams true "Модель которую принимает метод"
// @Success 200 {string} string "create order"
// @Failure 400 {string} string "Error"
// @Router /create_order [post].
func (h *Handler) CreateOrderWithProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		var products trx.CreateOrderParams
		if err := c.ShouldBindJSON(&products); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request",
			})
			return
		}
		order, err := trx.CreateOrderWithProducts(app.Ctx, products, app.DB)
		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{
				"error": fmt.Sprintf("Error create order: %v", err),
			})
			return
		}
		c.JSON(http.StatusOK, order)
	}
}
