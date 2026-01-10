package successhandler

import "github.com/gin-gonic/gin"

func HandleResponse(c *gin.Context, resp interface{}, statusCode int) {
	c.JSON(statusCode, gin.H{
		"data":    resp,
		"success": true,
	})
}
