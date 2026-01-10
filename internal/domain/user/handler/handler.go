package handler

import (
	"github.com/gin-gonic/gin"
	usecase "gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/user/usecase"
)

type UserHandler interface {
	FetchUserProfile(*gin.Context)
	CreateUser(*gin.Context)
	FetchUserPlayers(*gin.Context)
	AddUserPlayers(*gin.Context)
	FollowUser(*gin.Context)
	UnfollowUser(*gin.Context)
}

type userHandler struct {
	useCase usecase.UserUseCase
}

func NewUserHandler(useCase usecase.UserUseCase) UserHandler {
	return &userHandler{
		useCase: useCase,
	}
}
