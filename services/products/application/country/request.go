package country

import (
	"github.com/go-playground/validator/v10"
	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"
)

type CountryRequestDTO struct {
	CountryCode string `validate:"required"`
}

func (request *CountryRequestDTO) toEntity() *domain.Country {
	return &domain.Country{
		Code: request.CountryCode,
	}
}

func (request *CountryRequestDTO) validateRequest() map[string]string {
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
