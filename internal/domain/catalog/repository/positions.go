package repository

// Positions CRUD
func (r *repository) ListPositions() ([]Position, error) {
	var items []Position
	err := r.db.Table("positions").Order("id").Find(&items).Error
	return items, err
}

func (r *repository) CreatePosition(position *Position) error {
	return r.db.Create(position).Error
}

func (r *repository) GetPositionByID(id uint) (*Position, error) {
	var position Position
	err := r.db.First(&position, id).Error
	return &position, err
}

func (r *repository) UpdatePosition(position *Position) error {
	return r.db.Save(position).Error
}

func (r *repository) DeletePosition(id uint) error {
	return r.db.Delete(&Position{}, id).Error
}
