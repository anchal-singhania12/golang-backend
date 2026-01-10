package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/player/contract"
	"gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/player/errors"
	"gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/player/repository"
	errorhandler "gitlab.com/fanligafc-group/fanligafc-backend/pkg/error_handler"
)

func (h *playerHandler) CreatePlayer(c *gin.Context) {
	var req contract.PlayerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	player := &repository.Player{
		Name:             req.Name,
		PositionID:       req.PositionID,
		ProviderPlayerID: req.ProviderPlayerID,
	}

	err := h.playerUseCase.CreatePlayer(player)
	if err != nil {
		log.Printf("Error fetching player: %v", err)
		errorhandler.HandleError(c, err, errors.ErrorDetailsMap)
		return
	}
	HandleResponse(c, player, http.StatusOK)
}

func (h *playerHandler) DeletePlayer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.playerUseCase.DeletePlayer(uint(id))

	if err != nil {
		log.Printf("Error fetching player: %v", err)
		errorhandler.HandleError(c, err, errors.ErrorDetailsMap)
		return
	}
	HandleResponse(c, nil, http.StatusOK)

}

func (h *playerHandler) GetPlayerByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	player, err := h.playerUseCase.FetchPlayerByID(uint(id))

	if err != nil {
		log.Printf("Error fetching player: %v", err)
		errorhandler.HandleError(c, err, errors.ErrorDetailsMap)
		return
	}

	HandleResponse(c, player, http.StatusOK)
}

func (h *playerHandler) GetPlayersByPositionID(c *gin.Context) {
	posID, _ := strconv.Atoi(c.Param("position_id"))

	players, err := h.playerUseCase.FetchPlayersByPositionID(uint(posID))

	if err != nil {
		log.Printf("Error fetching player: %v", err)
		errorhandler.HandleError(c, err, errors.ErrorDetailsMap)
		return
	}

	HandleResponse(c, players, http.StatusOK)

}

func (h *playerHandler) UpdatePlayer(c *gin.Context) {
	var req contract.PlayerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	player := &repository.Player{
		Name:             req.Name,
		PositionID:       req.PositionID,
		ProviderPlayerID: req.ProviderPlayerID,
	}

	err := h.playerUseCase.UpdatePlayer(player)

	if err != nil {
		log.Printf("Error fetching player: %v", err)
		errorhandler.HandleError(c, err, errors.ErrorDetailsMap)
		return
	}
	HandleResponse(c, player, http.StatusOK)

}
