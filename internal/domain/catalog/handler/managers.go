package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	repo "gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/catalog/repository"
	errorhandler "gitlab.com/fanligafc-group/fanligafc-backend/pkg/error_handler"
	successhandler "gitlab.com/fanligafc-group/fanligafc-backend/pkg/success_handler"
)

// Managers
func (h *handler) ListManagers(c *gin.Context) {
	items, err := h.usecase.ListManagers()
	if err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	successhandler.HandleResponse(c, items, http.StatusOK)
}

func (h *handler) CreateManager(c *gin.Context) {
	var manager repo.Manager
	if err := c.ShouldBindJSON(&manager); err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	result, err := h.usecase.CreateManager(&manager)
	if err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	successhandler.HandleResponse(c, result, http.StatusCreated)
}

func (h *handler) GetManager(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	manager, err := h.usecase.GetManagerByID(uint(id))
	if err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	successhandler.HandleResponse(c, manager, http.StatusOK)
}

func (h *handler) UpdateManager(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	var manager repo.Manager
	if err := c.ShouldBindJSON(&manager); err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	manager.ID = uint(id)
	result, err := h.usecase.UpdateManager(&manager)
	if err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	successhandler.HandleResponse(c, result, http.StatusOK)
}

func (h *handler) DeleteManager(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	if err := h.usecase.DeleteManager(uint(id)); err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	successhandler.HandleResponse(c, gin.H{"message": "Manager deleted successfully"}, http.StatusOK)
}
