package repository

// Ranked Entities CRUD
func (r *repository) ListRankedEntities() ([]RankedEntity, error) {
	var items []RankedEntity
	err := r.db.Table("ranked_entities").Order("id").Find(&items).Error
	return items, err
}

func (r *repository) CreateRankedEntity(rankedEntity *RankedEntity) error {
	return r.db.Create(rankedEntity).Error
}

func (r *repository) GetRankedEntityByID(id uint) (*RankedEntity, error) {
	var rankedEntity RankedEntity
	err := r.db.First(&rankedEntity, id).Error
	return &rankedEntity, err
}

func (r *repository) UpdateRankedEntity(rankedEntity *RankedEntity) error {
	return r.db.Save(rankedEntity).Error
}

func (r *repository) DeleteRankedEntity(id uint) error {
	return r.db.Delete(&RankedEntity{}, id).Error
}
