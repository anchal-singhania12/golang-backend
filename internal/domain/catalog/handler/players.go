package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	repo "gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/catalog/repository"
	errorhandler "gitlab.com/fanligafc-group/fanligafc-backend/pkg/error_handler"
	successhandler "gitlab.com/fanligafc-group/fanligafc-backend/pkg/success_handler"
)

// Players
func (h *handler) ListPlayers(c *gin.Context) {
	items, err := h.usecase.ListPlayers()
	if err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	successhandler.HandleResponse(c, items, http.StatusOK)
}

func (h *handler) CreatePlayer(c *gin.Context) {
	var player repo.Player
	if err := c.ShouldBindJSON(&player); err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	result, err := h.usecase.CreatePlayer(&player)
	if err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	successhandler.HandleResponse(c, result, http.StatusCreated)
}

func (h *handler) GetPlayer(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	player, err := h.usecase.GetPlayerByID(uint(id))
	if err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	successhandler.HandleResponse(c, player, http.StatusOK)
}

func (h *handler) UpdatePlayer(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	var player repo.Player
	if err := c.ShouldBindJSON(&player); err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	player.ID = uint(id)
	result, err := h.usecase.UpdatePlayer(&player)
	if err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	successhandler.HandleResponse(c, result, http.StatusOK)
}

func (h *handler) DeletePlayer(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	if err := h.usecase.DeletePlayer(uint(id)); err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	successhandler.HandleResponse(c, gin.H{"message": "Player deleted successfully"}, http.StatusOK)
	// writeOK(c, gin.H{"message": "Player deleted successfully"})
}
