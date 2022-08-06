package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application/error"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application/stock"
)

type StockHandler struct {
	stockService stock.StockService
	errorService error.ErrorService
}

func NewStockHandler(services *application.Services) *StockHandler {
	return &StockHandler{
		stockService: services.StockService,
		errorService: services.ErrorService,
	}
}

func (sh *StockHandler) ConsumeStock(c *gin.Context) {
	var consumeStockRequest stock.ConsumeStockRequestDTO

	if err := c.BindJSON(&consumeStockRequest); err != nil {
		response := sh.errorService.GetGeneralError(http.StatusBadRequest, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	updateResponse, errorResponse := sh.stockService.ConsumeStock(&consumeStockRequest)

	if errorResponse != nil {
		c.JSON(errorResponse.Status, errorResponse)
		return
	}

	c.JSON(updateResponse.Status, updateResponse)
}
