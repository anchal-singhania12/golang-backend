package usecase

import (
	"gitlab.com/fanligafc-group/fanligafc-backend/internal/config"
	"gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/user/contracts"
	"gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/user/repository"
)

type UserUseCase interface {
	FetchUserProfileByID(uint) (*contracts.UserProfileResponse, error)
	CreateUser(repository.User) (*repository.User, error)
	FetchPlayersByUserID(uint) ([]repository.Player, error)
	AddPlayersForUser(uint, []uint) error
	FollowUser(uint, uint) error
	UnfollowUser(uint, uint) error
}

type userUseCase struct {
	cfg        *config.Config
	repository repository.Operations
}

func NewUserUseCase(cfg *config.Config, rep repository.Operations) UserUseCase {
	return &userUseCase{
		cfg:        cfg,
		repository: rep,
	}
}

type UserDetails struct {
	Name string
}
