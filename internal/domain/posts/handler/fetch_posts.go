package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FetchPostsHandler struct {
	usecase FetchPostsUsecase
}

type FetchPostsUsecase interface {
	CommonHomeFeed(ctx *gin.Context) (interface{}, error)
}

func NewFetchPostsHandler(usecase FetchPostsUsecase) *FetchPostsHandler {
	return &FetchPostsHandler{
		usecase: usecase,
	}
}

func (h *FetchPostsHandler) FetchCommonHomeFeed(c *gin.Context) {
	result, err := h.usecase.CommonHomeFeed(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
