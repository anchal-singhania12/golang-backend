package usecase

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/position/contract"
	position "gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/position/model"
	usecase "gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/position/usecase"
)

type PositionHandler struct {
	usecase usecase.PositionUseCase
}

func NewPositionHandler(r *gin.Engine, uc usecase.PositionUseCase) {
	h := &PositionHandler{usecase: uc}

	group := r.Group("/positions")
	group.POST("/", h.Create)
	group.GET("/", h.List)
	group.GET("/:id", h.GetByID)
	group.PUT("/:id", h.Update)
	group.DELETE("/:id", h.Delete)
}

func (h *PositionHandler) Create(c *gin.Context) {
	var req contract.CreatePositionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pos := &position.PositionModel{
		ProviderID: req.ProviderID,
		Name:       req.Name,
	}
	if err := h.usecase.Create(c, pos); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create position"})
		return
	}
	c.JSON(http.StatusCreated, pos)
}

func (h *PositionHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	pos, err := h.usecase.GetByID(c, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch position"})
		return
	}
	if pos == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Position not found"})
		return
	}
	c.JSON(http.StatusOK, pos)
}

func (h *PositionHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req contract.UpdatePositionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.usecase.Update(c, uint(id), req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update position"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Position updated"})
}

func (h *PositionHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.usecase.Delete(c, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete position"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Position deleted"})
}

func (h *PositionHandler) List(c *gin.Context) {
	positions, err := h.usecase.List(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list positions"})
		return
	}
	c.JSON(http.StatusOK, positions)
}
