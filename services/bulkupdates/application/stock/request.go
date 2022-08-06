package stock

import "github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/domain"

type StockRequestDTO struct {
	SKU      string
	Country  string
	Quantity int
}

func (request *StockRequestDTO) toEntity() *domain.Stock {
	return &domain.Stock{
		Count: request.Quantity,
	}
}
