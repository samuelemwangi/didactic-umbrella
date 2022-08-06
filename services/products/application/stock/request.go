package stock

import "github.com/samuelemwangi/jumia-mds-test/services/products/domain"

type ConsumeStockRequestDTO struct {
	ProductID uint
	Count     int
	CountryID uint
}

func (request *ConsumeStockRequestDTO) toEntity() *domain.Stock {
	return &domain.Stock{
		ProductID: request.ProductID,
		CountryID: request.CountryID,
	}
}
