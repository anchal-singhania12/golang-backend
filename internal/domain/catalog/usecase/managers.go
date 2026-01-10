package usecase

import repo "gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/catalog/repository"

// Managers
func (u *catalogUseCase) ListManagers() ([]repo.Manager, error) {
	return u.repository.ListManagers()
}

func (u *catalogUseCase) CreateManager(manager *repo.Manager) (*repo.Manager, error) {
	err := u.repository.CreateManager(manager)
	return manager, err
}

func (u *catalogUseCase) GetManagerByID(id uint) (*repo.Manager, error) {
	return u.repository.GetManagerByID(id)
}

func (u *catalogUseCase) UpdateManager(manager *repo.Manager) (*repo.Manager, error) {
	err := u.repository.UpdateManager(manager)
	return manager, err
}

func (u *catalogUseCase) DeleteManager(id uint) error {
	return u.repository.DeleteManager(id)
}
