package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sar0868/otus_go_basic_hw/hw16_docker/internal/app"
	"github.com/sar0868/otus_go_basic_hw/hw16_docker/internal/repository"
	"github.com/sar0868/otus_go_basic_hw/hw16_docker/internal/service"
	"github.com/sar0868/otus_go_basic_hw/hw16_docker/internal/trx"
)

type Product struct {
	Name     string  `db:"name" json:"product"`
	Quantity float64 `db:"quantity" json:"quantity"`
}

type CreateOrderParams struct {
	User     string `db:"name" json:"user"`
	Products []Product
}

type ProductAddOrderParams struct {
	OrderID  int32   `db:"order_id" json:"orderId"`
	Name     string  `db:"name" json:"name"`
	Quantity float64 `db:"quantity" json:"quantity"`
}

type ShopOrder struct {
	ID          int     `db:"id" json:"id"`
	UserID      int32   `db:"user_id" json:"user_id"`           //nolint: tagliatelle
	OrderDate   string  `db:"order_date" json:"order_date"`     //nolint: tagliatelle
	TotalAmount float64 `db:"total_amount" json:"total_amount"` //nolint: tagliatelle
}

type OrderProductUpdateParams struct {
	OrderID   int32   `json:"order_id"`   //nolint: tagliatelle
	ProductID int32   `json:"product_id"` //nolint: tagliatelle
	Quantity  float64 `json:"quantity"`
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

// Add product in order
// @Summary Order add product
// @Tags OrderAddProduct
// @Accept			json
// @Produce		json
// @Param input body ProductAddOrderParams true "Модель которую принимает метод"
// @Success 200 {string} string "create order"
// @Failure 400 {string} string "Error"
// @Router /order_add_product [post].
func (h *Handler) OrderAddProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var product repository.ProductAddOrderParams
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request",
			})
			return
		}
		order, err := trx.OrderAddProduct(app.Ctx, product, app.DB)
		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{
				"error": fmt.Sprintf("Error create order: %v", err),
			})
			return
		}
		c.JSON(http.StatusOK, order)
	}
}

type OrdersProductsFullRow struct {
	OrderID  int32   `db:"order_id" json:"order_id"` //nolint: tagliatelle
	Name     string  `db:"name" json:"name"`
	Quantity float64 `db:"quantity" json:"quantity"`
}

// Get Orders full
// @Summary get orders full
// @Tags getOrdersFull
// @Accept			json
// @Produce		json
// @Success 200 {array} OrdersProductsFullRow
// @Failure 400 {string} string "Error"
// @Router /orders_products [get].
func (h *Handler) OrdersProductsFull() gin.HandlerFunc {
	return func(c *gin.Context) {
		orders, err := service.OrdersProductsFull(app.Ctx, app.Repo)
		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{
				"error": fmt.Sprintf("Error get full orders: %v", err),
			})
			return
		}
		c.JSON(http.StatusOK, orders)
	}
}

type OrderRecountParams struct {
	OrderID int `json:"orderId"`
}

// Get Order recount
// @Summary order recount
// @Tags orderRecount
// @Accept			json
// @Produce		json
// @Param input body OrderRecountParams true "Модель которую принимает метод"
// @Success 200 {string} string "Order after recount"
// @Failure 400 {string} string "Error"
// @Router /order_recount [post].
func (h *Handler) OrderRecount() gin.HandlerFunc {
	return func(c *gin.Context) {
		var param OrderRecountParams
		if err := c.ShouldBindJSON(&param); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request",
			})
			return
		}
		id := param.OrderID
		order, err := service.OrderRecount(app.Ctx, app.Repo, id)
		if err != nil {
			msg := fmt.Sprintf("Error recount order id=%v: %s", id, err)
			c.JSON(http.StatusNotAcceptable, gin.H{
				"error": msg,
			})
			return
		}
		c.JSON(http.StatusOK, order)
	}
}

// Get Order update
// @Summary order update
// @Tags orderUpdateQuantity
// @Accept			json
// @Produce		json
// @Param input body OrderProductUpdateParams true "Модель которую принимает метод"
// @Success 200 {array} ShopOrder
// @Failure 400 {string} string "Error"
// @Router /order_update [post].
func (h *Handler) OrderProductUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params repository.OrderProductUpdateParams
		if err := c.ShouldBindJSON(&params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request",
			})
			return
		}
		order, err := service.OrderProductUpdate(app.Ctx, app.Repo, params)
		if err != nil {
			msg := fmt.Sprintf("Error update order: %s", err)
			c.JSON(http.StatusNotAcceptable, gin.H{
				"error": msg,
			})
			return
		}
		c.JSON(http.StatusOK, order)
	}
}

// Delete order by id
// @Summary delete order by id
// @Tags deleteOrder
// @Accept			json
// @Produce		json
// @Param id query string true "string valid"
// @Success 200 {string} string "Delete order by id"
// @Failure 400 {string} string "Error"
// @Router /del_order [delete].
func (h *Handler) OrderDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Query("id")
		err := service.OrderDelete(app.Ctx, app.Repo, idStr)
		if err != nil {
			msg := fmt.Sprintf("Error delete order id=%s: %s", idStr, err)
			c.JSON(http.StatusNotAcceptable, gin.H{
				"message": msg,
			})
			return
		}
		msg := fmt.Sprintf("Order id=%s delete", idStr)
		c.JSON(http.StatusOK, gin.H{
			"message": msg,
		})
	}
}
