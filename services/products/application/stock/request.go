package stock

import (
	"github.com/go-playground/validator/v10"
	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"
)

type ConsumeStockRequestDTO struct {
	ProductID uint `validate:"required"`
	Quantity  int  `validate:"required"`
	CountryID uint `validate:"required"`
}

func (request *ConsumeStockRequestDTO) toEntity() *domain.Stock {
	return &domain.Stock{
		ProductID: request.ProductID,
		CountryID: request.CountryID,
	}
}

func (request *ConsumeStockRequestDTO) validateRequest() map[string]string {
	errors := make(map[string]string)

	err := validator.New().Struct(request)
	if err == nil {
		return errors
	}

	for _, err := range err.(validator.ValidationErrors) {
		errors[err.Field()] = err.ActualTag()
	}
	return errors
}
