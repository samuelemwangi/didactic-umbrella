package stock

import "github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/domain"

type StockRequestDTO struct {
	Quantity  int
	CountryId uint
	ProductId uint
}

func (request *StockRequestDTO) toEntity() *domain.Stock {
	return &domain.Stock{
		Count:     request.Quantity,
		CountryID: request.CountryId,
		ProductID: request.ProductId,
	}
}
