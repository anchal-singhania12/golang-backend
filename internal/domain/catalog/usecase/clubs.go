package usecase

import repo "gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/catalog/repository"

// Clubs
func (u *catalogUseCase) ListClubs() ([]repo.Club, error) {
	return u.repository.ListClubs()
}

func (u *catalogUseCase) CreateClub(club *repo.Club) (*repo.Club, error) {
	err := u.repository.CreateClub(club)
	return club, err
}

func (u *catalogUseCase) GetClubByID(id uint) (*repo.Club, error) {
	return u.repository.GetClubByID(id)
}

func (u *catalogUseCase) UpdateClub(club *repo.Club) (*repo.Club, error) {
	err := u.repository.UpdateClub(club)
	return club, err
}

func (u *catalogUseCase) DeleteClub(id uint) error {
	return u.repository.DeleteClub(id)
}
