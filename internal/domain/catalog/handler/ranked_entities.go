package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	repo "gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/catalog/repository"
	errorhandler "gitlab.com/fanligafc-group/fanligafc-backend/pkg/error_handler"
	successhandler "gitlab.com/fanligafc-group/fanligafc-backend/pkg/success_handler"
)

// Ranked Entities
func (h *handler) ListRankedEntities(c *gin.Context) {
	items, err := h.usecase.ListRankedEntities()
	if err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	successhandler.HandleResponse(c, items, http.StatusOK)
}

func (h *handler) CreateRankedEntity(c *gin.Context) {
	var rankedEntity repo.RankedEntity
	if err := c.ShouldBindJSON(&rankedEntity); err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	result, err := h.usecase.CreateRankedEntity(&rankedEntity)
	if err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	successhandler.HandleResponse(c, result, http.StatusCreated)
}

func (h *handler) GetRankedEntity(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	rankedEntity, err := h.usecase.GetRankedEntityByID(uint(id))
	if err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	successhandler.HandleResponse(c, rankedEntity, http.StatusOK)
}

func (h *handler) UpdateRankedEntity(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	var rankedEntity repo.RankedEntity
	if err := c.ShouldBindJSON(&rankedEntity); err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	rankedEntity.ID = uint(id)
	result, err := h.usecase.UpdateRankedEntity(&rankedEntity)
	if err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	successhandler.HandleResponse(c, result, http.StatusOK)
}

func (h *handler) DeleteRankedEntity(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	if err := h.usecase.DeleteRankedEntity(uint(id)); err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	successhandler.HandleResponse(c, gin.H{"message": "Ranked entity deleted successfully"}, http.StatusOK)
}
