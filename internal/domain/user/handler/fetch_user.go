package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/user/contracts"
	"gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/user/errors"
	errorhandler "gitlab.com/fanligafc-group/fanligafc-backend/pkg/error_handler"
	successhandler "gitlab.com/fanligafc-group/fanligafc-backend/pkg/success_handler"
)

func (h *userHandler) FetchUserProfile(c *gin.Context) {

	var headers contracts.FetchUserProfileRequest

	if err := c.ShouldBindHeader(&headers); err != nil {
		log.Printf("Error in request validation: %v", err)
		errorhandler.HandleError(c, errors.ErrorBadRequest, errors.ErrorDetailsMap)
		return
	}

	log.Printf("Fetching user profile for user ID: %d", headers.UserID)

	resp, err := h.useCase.FetchUserProfileByID(headers.UserID)
	if err != nil {
		log.Printf("Error fetching user profile: %v", err)
		errorhandler.HandleError(c, err, errors.ErrorDetailsMap)
		return
	}

	successhandler.HandleResponse(c, resp, http.StatusOK)
}

func (h *userHandler) FetchUserPlayers(c *gin.Context) {
	var headers contracts.FetchUserPlayersRequest
	if err := c.ShouldBindHeader(&headers); err != nil {
		log.Printf("Error in request validation: %v", err)
		errorhandler.HandleError(c, errors.ErrorBadRequest, errors.ErrorDetailsMap)
		return
	}

	log.Printf("Fetching players for user ID: %d", headers.UserID)

	resp, err := h.useCase.FetchPlayersByUserID(headers.UserID)
	if err != nil {
		log.Printf("Error fetching user players: %v", err)
		errorhandler.HandleError(c, err, errors.ErrorDetailsMap)
		return
	}

	if len(resp) == 0 {
		errorhandler.HandleError(c, errors.ErrorUserPlayerMappingNotFound, errors.ErrorDetailsMap)
		return
	}

	successhandler.HandleResponse(c, resp, http.StatusOK)
}
