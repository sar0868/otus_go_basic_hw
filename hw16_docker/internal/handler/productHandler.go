package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sar0868/otus_go_basic_hw/hw16_docker/internal/app"
	"github.com/sar0868/otus_go_basic_hw/hw16_docker/internal/repository"
	"github.com/sar0868/otus_go_basic_hw/hw16_docker/internal/service"
)

type ProductUpdateParams struct {
	Price float64 `db:"price" json:"price"`
	Name  string  `db:"name" json:"name"`
}

type ProductCreateParams struct {
	Name  string  `db:"name" json:"name"`
	Price float64 `db:"price" json:"price"`
}

type ProductGetRangePriceParams struct {
	Price  float64 `db:"price" json:"price"`
	Price2 float64 `json:"price_2"` //nolint: tagliatelle
}

// Get Products
// @Summary get products
// @Tags getProducts
// @Accept			json
// @Produce		json
// @Success 200 {string} string "Get products"
// @Failure 400 {string} string "Error"
// @Router /products [get].
func (h *Handler) GetProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Request received get data users")
		products, err := service.Products(app.Ctx, app.Repo)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		c.JSON(http.StatusOK, products)
	}
}

// Get Product by name
// @Summary get Product by name
// @Tags getProductByName
// @Accept			json
// @Produce		json
// @Param name query string false "string valid"
// @Success 200 {string} string "Get product by name"
// @Failure 400 {string} string "Error"
// @Router /product [get].
func (h *Handler) GetProductByName() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.DefaultQuery("name", "")
		product, err := service.ProductByName(app.Ctx, app.Repo, name)
		if err != nil {
			msg := fmt.Sprintf("Error get product by name: %s", err)
			c.JSON(http.StatusNotAcceptable, gin.H{
				"message": msg,
			})
			return
		}
		c.JSON(http.StatusOK, product)
	}
}

// Add Product
// @Summary addProduct
// @Tags addProduct
// @Accept			json
// @Produce		json
// @Param input body ProductCreateParams true "Модель которую принимает метод"
// @Success 200 {string} string "Added product"
// @Failure 400 {string} string "Error"
// @Router /add_product [post].
func (h *Handler) ProductCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newProduct repository.ProductCreateParams
		if err := c.ShouldBindJSON(&newProduct); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request",
			})
			return
		}
		product, err := service.ProductCreate(app.Ctx, app.Repo, newProduct)
		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{
				"error": fmt.Sprintf("Error add product: %v", err),
			})
			return
		}
		c.JSON(http.StatusOK, product)
	}
}

// Update Product price
// @Summary update product price
// @Tags updateProductPrice
// @Accept			json
// @Produce		json
// @Param input body ProductUpdateParams true "Модель которую принимает метод"
// @Success 200 {string} string "Update product"
// @Failure 400 {string} string "Error"
// @Router /edit_product [post].
func (h *Handler) ProductUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var updateProduct repository.ProductUpdateParams
		if err := c.ShouldBindJSON(&updateProduct); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request",
			})
			return
		}
		product, err := service.ProductUpdate(app.Ctx, app.Repo, updateProduct)
		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{
				"error": fmt.Sprintf("Error update product: %v", err),
			})
			return
		}
		c.JSON(http.StatusOK, product)
	}
}

// Delete product
// @Summary delete product
// @Tags deleteProduct
// @Accept			json
// @Produce		json
// @Param name query string true "string valid"
// @Success 200 {string} string "Delete product"
// @Failure 400 {string} string "Error"
// @Router /del_product [delete].
func (h *Handler) ProductDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.DefaultQuery("name", "")
		err := service.DeleteProduct(app.Ctx, app.Repo, name)
		if err != nil {
			msg := fmt.Sprintf("Error delete product %s: %s", name, err)
			c.JSON(http.StatusNotAcceptable, gin.H{
				"message": msg,
			})
			return
		}
		msg := fmt.Sprintf("Product %s delete", name)
		c.JSON(http.StatusOK, gin.H{
			"message": msg,
		})
	}
}

// Get Products for range
// @Summary get Products for range
// @Tags getProductsForEange
// @Accept			json
// @Produce		json
// @Param input body ProductGetRangePriceParams true "Модель которую принимает метод"
// @Success 200 {string} string "get products"
// @Failure 400 {string} string "Error"
// @Router /products_prices [post].
func (h *Handler) ProductsGetRangePrice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var paramsPrice repository.ProductGetRangePriceParams
		if err := c.ShouldBindJSON(&paramsPrice); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request",
			})
			return
		}
		products, err := service.ProductsGetRangePrice(app.Ctx, app.Repo, paramsPrice)
		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{
				"error": fmt.Sprintf("Error get products: %v", err),
			})
			return
		}
		c.JSON(http.StatusOK, products)
	}
}
