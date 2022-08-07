package product

import (
	"github.com/go-playground/validator/v10"
	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"
)

type ProductRequestDTO struct {
	SKU string `validate:"required"`
}

func (request *ProductRequestDTO) toEntity() *domain.Product {
	return &domain.Product{
		SKU: request.SKU,
	}
}

func (request *ProductRequestDTO) validateRequest() map[string]string {
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
