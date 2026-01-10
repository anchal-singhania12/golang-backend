package repository

// Countries CRUD
func (r *repository) ListCountries() ([]Country, error) {
	var items []Country
	err := r.db.Table("country").Order("id").Find(&items).Error
	return items, err
}

func (r *repository) CreateCountry(country *Country) error {
	return r.db.Table("country").Create(country).Error
}

func (r *repository) GetCountryByID(id uint) (*Country, error) {
	var country Country
	err := r.db.Table("country").Where("id = ?", id).First(&country).Error
	return &country, err
}

func (r *repository) UpdateCountry(country *Country) error {
	return r.db.Table("country").Where("id = ?", country.ID).Updates(country).Error
}

func (r *repository) DeleteCountry(id uint) error {
	return r.db.Table("country").Where("id = ?", id).Delete(&Country{}).Error
}
