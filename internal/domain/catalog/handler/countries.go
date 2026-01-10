package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	repo "gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/catalog/repository"
	errorhandler "gitlab.com/fanligafc-group/fanligafc-backend/pkg/error_handler"
	successhandler "gitlab.com/fanligafc-group/fanligafc-backend/pkg/success_handler"
)

// Countries
func (h *handler) ListCountries(c *gin.Context) {
	items, err := h.usecase.ListCountries()
	if err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	successhandler.HandleResponse(c, items, http.StatusOK)
}

func (h *handler) CreateCountry(c *gin.Context) {
	var country repo.Country
	if err := c.ShouldBindJSON(&country); err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	result, err := h.usecase.CreateCountry(&country)
	if err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	successhandler.HandleResponse(c, result, http.StatusCreated)
}

func (h *handler) GetCountry(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	country, err := h.usecase.GetCountryByID(uint(id))
	if err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	successhandler.HandleResponse(c, country, http.StatusOK)
}

func (h *handler) UpdateCountry(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	var country repo.Country
	if err := c.ShouldBindJSON(&country); err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	country.ID = uint(id)
	result, err := h.usecase.UpdateCountry(&country)
	if err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	successhandler.HandleResponse(c, result, http.StatusOK)
}

func (h *handler) DeleteCountry(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	if err := h.usecase.DeleteCountry(uint(id)); err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	successhandler.HandleResponse(c, gin.H{"message": "Country deleted successfully"}, http.StatusOK)
}
