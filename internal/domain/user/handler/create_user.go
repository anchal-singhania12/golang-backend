package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/user/contracts"
	"gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/user/errors"
	"gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/user/repository"
	errorhandler "gitlab.com/fanligafc-group/fanligafc-backend/pkg/error_handler"
	successhandler "gitlab.com/fanligafc-group/fanligafc-backend/pkg/success_handler"
)

func (h *userHandler) CreateUser(c *gin.Context) {
	var user contracts.CreateUserRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("Error in request validation: %v", err)
		errorhandler.HandleError(c, errors.ErrorBadRequest, errors.ErrorDetailsMap)
		return
	}

	log.Printf("Creating user with name: %s, user_name: %s", user.Name, user.UserName)

	userModel := repository.User{
		Name:          user.Name,
		UserName:      user.UserName,
		Photo:         user.Photo,
		Chant:         user.Chant,
		Bio:           user.Bio,
		Badge:         user.Badge,
		BestPlayerID:  user.BestPlayerID,
		BestManagerID: user.BestManagerID,
		BestClubID:    user.BestClubID,
		BestCountryID: user.BestCountryID,
	}

	resp, err := h.useCase.CreateUser(userModel)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		errorhandler.HandleError(c, err, errors.ErrorDetailsMap)
		return
	}

	successhandler.HandleResponse(c, resp, http.StatusCreated)
}

func (h *userHandler) AddUserPlayers(c *gin.Context) {
	var headers contracts.FetchUserPlayersRequest
	if err := c.ShouldBindHeader(&headers); err != nil {
		log.Printf("Error in request validation: %v", err)
		errorhandler.HandleError(c, errors.ErrorBadRequest, errors.ErrorDetailsMap)
		return
	}

	var body contracts.AddUserPlayersRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Printf("Error in request validation: %v", err)
		errorhandler.HandleError(c, errors.ErrorBadRequest, errors.ErrorDetailsMap)
		return
	}

	if len(body.PlayerIDs) == 0 {
		errorhandler.HandleError(c, errors.ErrorBadRequest, errors.ErrorDetailsMap)
		return
	}

	if err := h.useCase.AddPlayersForUser(headers.UserID, body.PlayerIDs); err != nil {
		log.Printf("Error adding user players: %v", err)
		errorhandler.HandleError(c, err, errors.ErrorDetailsMap)
		return
	}

	successhandler.HandleResponse(c, gin.H{"added": len(body.PlayerIDs)}, http.StatusCreated)
}

func (h *userHandler) FollowUser(c *gin.Context) {
	var headers contracts.FetchUserPlayersRequest
	if err := c.ShouldBindHeader(&headers); err != nil {
		log.Printf("Error in request validation: %v", err)
		errorhandler.HandleError(c, errors.ErrorBadRequest, errors.ErrorDetailsMap)
		return
	}

	var body contracts.FollowUserRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Printf("Error in request validation: %v", err)
		errorhandler.HandleError(c, errors.ErrorBadRequest, errors.ErrorDetailsMap)
		return
	}

	if err := h.useCase.FollowUser(headers.UserID, body.FollowingID); err != nil {
		errorhandler.HandleError(c, err, errors.ErrorDetailsMap)
		return
	}
	successhandler.HandleResponse(c, gin.H{"followed": true}, http.StatusCreated)
}

func (h *userHandler) UnfollowUser(c *gin.Context) {
	var headers contracts.FetchUserPlayersRequest
	if err := c.ShouldBindHeader(&headers); err != nil {
		log.Printf("Error in request validation: %v", err)
		errorhandler.HandleError(c, errors.ErrorBadRequest, errors.ErrorDetailsMap)
		return
	}

	var body contracts.FollowUserRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Printf("Error in request validation: %v", err)
		errorhandler.HandleError(c, errors.ErrorBadRequest, errors.ErrorDetailsMap)
		return
	}

	if err := h.useCase.UnfollowUser(headers.UserID, body.FollowingID); err != nil {
		errorhandler.HandleError(c, err, errors.ErrorDetailsMap)
		return
	}
	successhandler.HandleResponse(c, gin.H{"unfollowed": true}, http.StatusOK)
}
