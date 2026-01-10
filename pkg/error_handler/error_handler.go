package errorhandler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GenericError struct {
	Code      string                 `json:"code"`
	Message   string                 `json:"message"`
	Timestamp *timestamppb.Timestamp `json:"timestamp"`
}

type ErrorDetails struct {
	Message string
	Code    int
}

func HandleError(c *gin.Context, err error, errorDetailsMap map[string]ErrorDetails) {

	errorDetails, ok := errorDetailsMap[err.Error()]

	if err != nil || !ok {
		c.JSON(errorDetails.Code, gin.H{
			"data":    nil,
			"success": false,
			"errors": GenericError{
				Code:      err.Error(),
				Message:   errorDetails.Message,
				Timestamp: timestamppb.Now(),
			},
		})
		return
	} else {

		log.Printf("Unable to find error")
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":    nil,
			"success": false,
			"errors": GenericError{
				Code:    "internal_server_error",
				Message: "Something went wrong, please try again later.",
			}})
		return
	}

}
