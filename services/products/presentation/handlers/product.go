package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application/product"
)

type ProductHandler struct {
	productService product.ProductService
}

func NewProductHandler(services *application.Services) *ProductHandler {
	return &ProductHandler{
		productService: services.ProductService,
	}
}

func (handler *ProductHandler) GetProductBySKU(c *gin.Context) {
	sku := c.Param("sku")

	productRequest := &product.ProductRequestDTO{
		SKU: sku,
	}
	productResponse, errorResponse := handler.productService.GetProductBySKU(productRequest)

	if errorResponse != nil {
		c.JSON(errorResponse.Status, errorResponse)
		return
	}
	c.JSON(http.StatusOK, productResponse)
}

func (handler *ProductHandler) GetProducts(c *gin.Context) {
	productsResponse, errorResponse := handler.productService.GetProducts()

	if errorResponse != nil {
		c.JSON(errorResponse.Status, errorResponse)
		return
	}
	c.JSON(http.StatusOK, productsResponse)
}
