package usecase

import (
	"gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/player/repository"
)

func (uc *playerUseCase) FetchPlayerByID(id uint) (*repository.Player, error) {
	return uc.repository.FindByID(id)
}
