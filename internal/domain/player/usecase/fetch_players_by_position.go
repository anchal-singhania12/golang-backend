package usecase

import "gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/player/repository"

func (uc *playerUseCase) FetchPlayersByPositionID(positionID uint) ([]*repository.Player, error) {
	return uc.repository.FindByPositionID(positionID)
}
