package country

import (
	"fmt"

	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"
)

type CountryResponseDTO struct {
	ID          string `json:"id"`
	CountryName string `json:"countrName"`
}

func (c *CountryResponseDTO) toResponseDTO(entity *domain.Country) {
	c.ID = fmt.Sprint(entity.ID)
	c.CountryName = entity.Name

}
