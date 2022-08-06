package stock

import (
	"net/http"

	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"
)

type StockDetailDTO struct {
	ID        uint `json:"id"`
	Count     int  `json:"count"`
	ProductID uint `json:"productId"`
	CountryID uint `json:"countryId"`
}

type ConsumeStockResponseDTO struct {
	Status  int             `json:"responseStatus"`
	Message string          `json:"responseMessage"`
	Item    *StockDetailDTO `json:"itemDetails"`
}

func (cr *ConsumeStockResponseDTO) toResponseDTO(stock *domain.Stock) {
	stockDetail := &StockDetailDTO{
		ID:        stock.ID,
		Count:     stock.Count,
		CountryID: stock.CountryID,
		ProductID: stock.ProductID,
	}

	cr.Status = http.StatusOK
	cr.Message = "request successful"
	cr.Item = stockDetail

}
