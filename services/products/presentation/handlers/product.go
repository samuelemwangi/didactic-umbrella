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

func (ph *ProductHandler) GetProductBySKU(c *gin.Context) {
	sku := c.Param("sku")

	productResponse, errorResponse := ph.productService.GetProductBySKU(sku)

	if errorResponse != nil {
		c.JSON(errorResponse.Status, errorResponse)
		return
	}

	c.JSON(http.StatusOK, productResponse)
}