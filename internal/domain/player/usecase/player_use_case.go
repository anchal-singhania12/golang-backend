package usecase

import (
	"gitlab.com/fanligafc-group/fanligafc-backend/internal/config"
	"gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/player/contract"
	"gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/player/repository"
)

type PlayerUseCase interface {
	CreatePlayer(player *repository.Player) error
	UpdatePlayer(player *repository.Player) error
	DeletePlayer(id uint) error
	FetchPlayerByID(id uint) (*repository.Player, error)
	FetchPlayersByPositionID(positionID uint) ([]*repository.Player, error)
	FetchUserTeam(userID uint) (*contract.UserTeamResponse, error)
}

type playerUseCase struct {
	cfg        *config.Config
	repository repository.PlayerRepository
}

func NewPlayerUseCase(cfg *config.Config, repository repository.PlayerRepository) PlayerUseCase {
	return &playerUseCase{
		cfg:        cfg,
		repository: repository,
	}
}
