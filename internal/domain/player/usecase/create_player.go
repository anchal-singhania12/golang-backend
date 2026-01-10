package usecase

import (
	"gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/player/repository"
)

func (uc *playerUseCase) CreatePlayer(player *repository.Player) error {
	return uc.repository.Create(player)
}
