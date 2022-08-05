package presentation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application"
	"github.com/samuelemwangi/jumia-mds-test/services/products/domain"
)

type CountryHandler struct {
	countryService application.CountryService
}

func NewCountryHandler(countryService application.CountryService) *CountryHandler {
	return &CountryHandler{
		countryService: countryService,
	}
}

func (ch *CountryHandler) SaveCountry(c *gin.Context) {
	var country domain.Country

	if err := c.BindJSON(&country); err != nil {
		c.JSON(http.StatusBadRequest, "bad request")
		return
	}

	savedCountry, savingError := ch.countryService.SaveCountry(&country)

	if savingError != nil {
		c.JSON(http.StatusInternalServerError, savingError)
		return
	}

	c.JSON(http.StatusOK, savedCountry)

}
