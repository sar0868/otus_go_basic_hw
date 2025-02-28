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
	api.POST("/edit_user", h.UpdateUser())
	api.DELETE("/del_user", h.DeleteUser())
	api.GET("/products", h.GetProducts())
	api.GET("/product", h.GetProductByName())
	api.POST("/add_product", h.ProductCreate())
	api.POST("/edit_product", h.ProductUpdate())
	api.POST("/products_prices", h.ProductsGetRangePrice())
	api.DELETE("/del_product", h.ProductDelete())
	api.GET("/orders", h.GetOrders())
	api.POST("/create_order", h.CreateOrderWithProducts())
	api.POST("/order_add_product", h.OrderAddProduct())
	api.GET("/orders_products", h.OrdersProductsFull())
	api.POST("/order_recount", h.OrderRecount())
	api.POST("/order_update", h.OrderProductUpdate())
	api.GET("/statistic", h.Statistic())
	api.DELETE("/del_order", h.OrderDelete())
}
