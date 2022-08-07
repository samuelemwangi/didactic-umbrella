package stock

import (
	"net/http"

	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"
)

type ConsumeStockDetailDTO struct {
	ID        uint `json:"id"`
	Quantity  int  `json:"count"`
	ProductID uint `json:"productId"`
	CountryID uint `json:"countryId"`
}

type ConsumeStockResponseDTO struct {
	Status  int                    `json:"responseStatus"`
	Message string                 `json:"responseMessage"`
	Item    *ConsumeStockDetailDTO `json:"itemDetails"`
}

func (response *ConsumeStockResponseDTO) toResponseDTO(stock *domain.Stock) {
	stockDetail := &ConsumeStockDetailDTO{
		ID:        stock.ID,
		Quantity:  stock.Quantity,
		CountryID: stock.CountryID,
		ProductID: stock.ProductID,
	}

	response.Status = http.StatusOK
	response.Message = "request successful"
	response.Item = stockDetail

}
