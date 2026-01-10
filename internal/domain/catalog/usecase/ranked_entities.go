package usecase

import repo "gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/catalog/repository"

// Ranked Entities
func (u *catalogUseCase) ListRankedEntities() ([]repo.RankedEntity, error) {
	return u.repository.ListRankedEntities()
}

func (u *catalogUseCase) CreateRankedEntity(rankedEntity *repo.RankedEntity) (*repo.RankedEntity, error) {
	err := u.repository.CreateRankedEntity(rankedEntity)
	return rankedEntity, err
}

func (u *catalogUseCase) GetRankedEntityByID(id uint) (*repo.RankedEntity, error) {
	return u.repository.GetRankedEntityByID(id)
}

func (u *catalogUseCase) UpdateRankedEntity(rankedEntity *repo.RankedEntity) (*repo.RankedEntity, error) {
	err := u.repository.UpdateRankedEntity(rankedEntity)
	return rankedEntity, err
}

func (u *catalogUseCase) DeleteRankedEntity(id uint) error {
	return u.repository.DeleteRankedEntity(id)
}
