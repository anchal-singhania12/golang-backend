package usecase

import repo "gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/catalog/repository"

// Positions
func (u *catalogUseCase) ListPositions() ([]repo.Position, error) {
	return u.repository.ListPositions()
}

func (u *catalogUseCase) CreatePosition(position *repo.Position) (*repo.Position, error) {
	err := u.repository.CreatePosition(position)
	return position, err
}

func (u *catalogUseCase) GetPositionByID(id uint) (*repo.Position, error) {
	return u.repository.GetPositionByID(id)
}

func (u *catalogUseCase) UpdatePosition(position *repo.Position) (*repo.Position, error) {
	err := u.repository.UpdatePosition(position)
	return position, err
}

func (u *catalogUseCase) DeletePosition(id uint) error {
	return u.repository.DeletePosition(id)
}
