package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application/errorhelper"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application/stock"
)

type StockHandler struct {
	stockService stock.StockService
	errorService errorhelper.ErrorService
}

func NewStockHandler(services *application.Services) *StockHandler {
	return &StockHandler{
		stockService: services.StockService,
		errorService: services.ErrorService,
	}
}

func (handler *StockHandler) ConsumeStock(c *gin.Context) {
	var consumeStockRequest stock.ConsumeStockRequestDTO

	if err := c.BindJSON(&consumeStockRequest); err != nil {
		response := handler.errorService.GetGeneralError(http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	updateResponse, errorResponse := handler.stockService.ConsumeStock(&consumeStockRequest)

	if errorResponse != nil {
		c.JSON(errorResponse.Status, errorResponse)
		return
	}

	c.JSON(updateResponse.Status, updateResponse)
}
