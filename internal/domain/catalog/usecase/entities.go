package usecase

import repo "gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/catalog/repository"

// Entities
func (u *catalogUseCase) ListEntities() ([]repo.Entity, error) {
	return u.repository.ListEntities()
}

func (u *catalogUseCase) CreateEntity(entity *repo.Entity) (*repo.Entity, error) {
	err := u.repository.CreateEntity(entity)
	return entity, err
}

func (u *catalogUseCase) GetEntityByID(id uint) (*repo.Entity, error) {
	return u.repository.GetEntityByID(id)
}

func (u *catalogUseCase) UpdateEntity(entity *repo.Entity) (*repo.Entity, error) {
	err := u.repository.UpdateEntity(entity)
	return entity, err
}

func (u *catalogUseCase) DeleteEntity(id uint) error {
	return u.repository.DeleteEntity(id)
}
