package repository

// Entities CRUD
func (r *repository) ListEntities() ([]Entity, error) {
	var items []Entity
	err := r.db.Table("entity").Order("id").Find(&items).Error
	return items, err
}

func (r *repository) CreateEntity(entity *Entity) error {
	return r.db.Table("entity").Create(entity).Error
}

func (r *repository) GetEntityByID(id uint) (*Entity, error) {
	var entity Entity
	err := r.db.Table("entity").Where("id = ?", id).First(&entity).Error
	return &entity, err
}

func (r *repository) UpdateEntity(entity *Entity) error {
	return r.db.Table("entity").Where("id = ?", entity.ID).Updates(entity).Error
}

func (r *repository) DeleteEntity(id uint) error {
	return r.db.Table("entity").Where("id = ?", id).Delete(&Entity{}).Error
}
