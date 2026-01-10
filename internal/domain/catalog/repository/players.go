package repository

// Players CRUD
func (r *repository) ListPlayers() ([]Player, error) {
	var items []Player
	err := r.db.Table("players").Order("id").Find(&items).Error
	return items, err
}

func (r *repository) CreatePlayer(player *Player) error {
	return r.db.Create(player).Error
}

func (r *repository) GetPlayerByID(id uint) (*Player, error) {
	var player Player
	err := r.db.First(&player, id).Error
	return &player, err
}

func (r *repository) UpdatePlayer(player *Player) error {
	return r.db.Save(player).Error
}

func (r *repository) DeletePlayer(id uint) error {
	return r.db.Delete(&Player{}, id).Error
}
