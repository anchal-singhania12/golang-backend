package usecase

import repo "gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/catalog/repository"

// Countries
func (u *catalogUseCase) ListCountries() ([]repo.Country, error) {
	return u.repository.ListCountries()
}

func (u *catalogUseCase) CreateCountry(country *repo.Country) (*repo.Country, error) {
	err := u.repository.CreateCountry(country)
	return country, err
}

func (u *catalogUseCase) GetCountryByID(id uint) (*repo.Country, error) {
	return u.repository.GetCountryByID(id)
}

func (u *catalogUseCase) UpdateCountry(country *repo.Country) (*repo.Country, error) {
	err := u.repository.UpdateCountry(country)
	return country, err
}

func (u *catalogUseCase) DeleteCountry(id uint) error {
	return u.repository.DeleteCountry(id)
}
