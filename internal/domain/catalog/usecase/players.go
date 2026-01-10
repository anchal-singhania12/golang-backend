package usecase

import repo "gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/catalog/repository"

// Players
func (u *catalogUseCase) ListPlayers() ([]repo.Player, error) {
	return u.repository.ListPlayers()
}

func (u *catalogUseCase) CreatePlayer(player *repo.Player) (*repo.Player, error) {
	err := u.repository.CreatePlayer(player)
	return player, err
}

func (u *catalogUseCase) GetPlayerByID(id uint) (*repo.Player, error) {
	return u.repository.GetPlayerByID(id)
}

func (u *catalogUseCase) UpdatePlayer(player *repo.Player) (*repo.Player, error) {
	err := u.repository.UpdatePlayer(player)
	return player, err
}

func (u *catalogUseCase) DeletePlayer(id uint) error {
	return u.repository.DeletePlayer(id)
}
