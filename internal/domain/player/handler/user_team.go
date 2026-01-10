package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/player/contract"
	"gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/player/errors"
	errorhandler "gitlab.com/fanligafc-group/fanligafc-backend/pkg/error_handler"
)

func (h *playerHandler) FetchUserTeam(c *gin.Context) {
	var req contract.FetchUserTeamRequest
	if err := c.ShouldBindHeader(&req); err != nil {
		log.Printf("Error in verifying request: %v", err)
		errorhandler.HandleError(c, err, errors.ErrorDetailsMap)
		return
	}

	team, err := h.playerUseCase.FetchUserTeam(req.UserID)
	if err != nil {
		log.Printf("Error in fetching user team: %v", err)
		errorhandler.HandleError(c, err, errors.ErrorDetailsMap)
		return
	}

	HandleResponse(c, team, http.StatusOK)
}
