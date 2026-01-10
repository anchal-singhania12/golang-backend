package handler

import (
	"github.com/gin-gonic/gin"
	usecase "gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/player/usecase"
)

type PlayerHandler interface {
	CreatePlayer(*gin.Context)
	UpdatePlayer(*gin.Context)
	DeletePlayer(*gin.Context)
	GetPlayerByID(*gin.Context)
	GetPlayersByPositionID(*gin.Context)
	FetchUserTeam(c *gin.Context)
}

type playerHandler struct {
	playerUseCase usecase.PlayerUseCase
}

// Constructor
func NewPlayerHandler(useCase usecase.PlayerUseCase) PlayerHandler {
	return &playerHandler{
		playerUseCase: useCase,
	}
}

func HandleResponse(c *gin.Context, resp interface{}, statusCode int) {
	c.JSON(statusCode, gin.H{
		"data":    resp,
		"success": true,
	})
}
