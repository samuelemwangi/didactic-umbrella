package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application/product"
)

type ProductHandler struct {
	productService product.ProductService
}

func NewProductHandler(productService product.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
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
