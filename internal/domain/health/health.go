package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HealthHandler struct {
	db *gorm.DB
}

func NewHealthHandler(db *gorm.DB) *HealthHandler {
	handler := &HealthHandler{
		db: db,
	}

	return handler
}

func (h *HealthHandler) Check(c *gin.Context) {
	// get underlying *sql.DB from GORM
	sqlDB, err := h.db.DB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "DOWN",
			"error":   "unable to get raw DB connection",
			"details": err.Error(),
		})
		return
	}

	// ping with a short timeout context
	if err := sqlDB.Ping(); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":  "DOWN",
			"error":   "database ping failed",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "UP",
		"version": "1.0.0",
	})
}
