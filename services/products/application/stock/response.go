package stock

import (
	"net/http"

	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"
)

type consumeStockDetailDTO struct {
	ID              uint   `json:"id"`
	QuantityBalance int    `json:"quantityBalance"`
	ProductID       uint   `json:"productId"`
	ProductSKU      string `json:"productSKU"`
	ProductName     string `json:"productName"`
	CountryID       uint   `json:"countryId"`
	CountryCode     string `json:"countryCode"`
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
		ProductSKU:      stock.Product.SKU,
		ProductName:     stock.Product.Name,
		CountryID:       stock.CountryID,
		CountryCode:     stock.Country.Code,
	}

	response.Status = http.StatusOK
	response.Message = "request successful"
	response.Item = stockDetail

}
