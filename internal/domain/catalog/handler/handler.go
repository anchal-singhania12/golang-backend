package handler

import (
	"github.com/gin-gonic/gin"
	uc "gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/catalog/usecase"
)

type Handler interface {
	// Positions
	ListPositions(*gin.Context)
	CreatePosition(*gin.Context)
	GetPosition(*gin.Context)
	UpdatePosition(*gin.Context)
	DeletePosition(*gin.Context)

	// Clubs
	ListClubs(*gin.Context)
	CreateClub(*gin.Context)
	GetClub(*gin.Context)
	UpdateClub(*gin.Context)
	DeleteClub(*gin.Context)

	// Managers
	ListManagers(*gin.Context)
	CreateManager(*gin.Context)
	GetManager(*gin.Context)
	UpdateManager(*gin.Context)
	DeleteManager(*gin.Context)

	// Countries
	ListCountries(*gin.Context)
	CreateCountry(*gin.Context)
	GetCountry(*gin.Context)
	UpdateCountry(*gin.Context)
	DeleteCountry(*gin.Context)

	// Players
	ListPlayers(*gin.Context)
	CreatePlayer(*gin.Context)
	GetPlayer(*gin.Context)
	UpdatePlayer(*gin.Context)
	DeletePlayer(*gin.Context)

	// Entities
	ListEntities(*gin.Context)
	CreateEntity(*gin.Context)
	GetEntity(*gin.Context)
	UpdateEntity(*gin.Context)
	DeleteEntity(*gin.Context)

	// Ranked Entities
	ListRankedEntities(*gin.Context)
	CreateRankedEntity(*gin.Context)
	GetRankedEntity(*gin.Context)
	UpdateRankedEntity(*gin.Context)
	DeleteRankedEntity(*gin.Context)
}

type handler struct{ usecase uc.CatalogUseCase }

func NewHandler(u uc.CatalogUseCase) Handler { return &handler{usecase: u} }

// Helper functions
// func writeOK(c *gin.Context, data interface{}) {
// 	c.JSON(http.StatusOK, gin.H{"data": data, "success": true})
// }

// func writeCreated(c *gin.Context, data interface{}) {
// 	c.JSON(http.StatusCreated, gin.H{"data": data, "success": true})
// }

// func writeError(c *gin.Context, err error) {
// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "success": false})
// }

// func writeBadRequest(c *gin.Context, message string) {
// 	c.JSON(http.StatusBadRequest, gin.H{"error": message, "success": false})
// }

// func writeNotFound(c *gin.Context, message string) {
// 	c.JSON(http.StatusNotFound, gin.H{"error": message, "success": false})
// }
