package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application/country"
)

type CountryHandler struct {
	countryService country.CountryService
}

func NewCountryHandler(services *application.Services) *CountryHandler {
	return &CountryHandler{
		countryService: services.CountryService,
	}
}

func (handler *CountryHandler) SaveCountry(c *gin.Context) {
	var countryRequest country.CountryRequestDTO

	if err := c.BindJSON(&countryRequest); err != nil {
		c.JSON(http.StatusBadRequest, "bad request")
		return
	}

	savedCountry, savingError := handler.countryService.SaveCountry(&countryRequest)

	if savingError != nil {
		c.JSON(savingError.Status, savingError)
		return
	}

	savedCountry.Status = http.StatusCreated

	c.JSON(savedCountry.Status, savedCountry)

}

func (handler *CountryHandler) GetCountries(c *gin.Context) {
	countries, errorResponse := handler.countryService.GetCountries()

	if errorResponse != nil {
		c.JSON(errorResponse.Status, errorResponse)
		return
	}
	c.JSON(http.StatusOK, countries)
}
