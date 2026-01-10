package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	repo "gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/catalog/repository"
	errorhandler "gitlab.com/fanligafc-group/fanligafc-backend/pkg/error_handler"
	successhandler "gitlab.com/fanligafc-group/fanligafc-backend/pkg/success_handler"
)

// Clubs
func (h *handler) ListClubs(c *gin.Context) {
	items, err := h.usecase.ListClubs()
	if err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	successhandler.HandleResponse(c, items, http.StatusOK)
}

func (h *handler) CreateClub(c *gin.Context) {
	var club repo.Club
	if err := c.ShouldBindJSON(&club); err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	result, err := h.usecase.CreateClub(&club)
	if err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}

	successhandler.HandleResponse(c, gin.H{"added": result}, http.StatusCreated)

}

func (h *handler) GetClub(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	club, err := h.usecase.GetClubByID(uint(id))
	if err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	successhandler.HandleResponse(c, club, http.StatusOK)
}

func (h *handler) UpdateClub(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	var club repo.Club
	if err := c.ShouldBindJSON(&club); err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	club.ID = uint(id)
	result, err := h.usecase.UpdateClub(&club)
	if err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	successhandler.HandleResponse(c, result, http.StatusOK)
}

func (h *handler) DeleteClub(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	if err := h.usecase.DeleteClub(uint(id)); err != nil {
		errorhandler.HandleError(c, err, nil)
		return
	}
	successhandler.HandleResponse(c, gin.H{"message": "Club deleted successfully"}, http.StatusOK)
}
