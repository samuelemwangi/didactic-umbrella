package stock

import (
	"net/http"

	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"
)

type consumeStockDetailDTO struct {
	ID              uint `json:"id"`
	QuantityBalance int  `json:"quantityBalance"`
	ProductID       uint `json:"productId"`
	CountryID       uint `json:"countryId"`
}

type ConsumeStockResponseDTO struct {
	Status  int                    `json:"responseStatus"`
	Message string                 `json:"responseMessage"`
	Item    *consumeStockDetailDTO `json:"itemDetails"`
}

func (response *ConsumeStockResponseDTO) toResponseDTO(stock *domain.Stock) {
	stockDetail := &consumeStockDetailDTO{
		ID:              stock.ID,
		QuantityBalance: stock.Quantity,
		ProductID:       stock.ProductID,
		CountryID:       stock.CountryID,
	}

	response.Status = http.StatusOK
	response.Message = "request successful"
	response.Item = stockDetail

}
