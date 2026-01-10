package repository

// Clubs CRUD
func (r *repository) ListClubs() ([]Club, error) {
	var items []Club
	err := r.db.Table("clubs").Order("id").Find(&items).Error
	return items, err
}

func (r *repository) CreateClub(club *Club) error {
	return r.db.Create(club).Error
}

func (r *repository) GetClubByID(id uint) (*Club, error) {
	var club Club
	err := r.db.First(&club, id).Error
	return &club, err
}

func (r *repository) UpdateClub(club *Club) error {
	return r.db.Save(club).Error
}

func (r *repository) DeleteClub(id uint) error {
	return r.db.Delete(&Club{}, id).Error
}
