package repository

// Managers CRUD
func (r *repository) ListManagers() ([]Manager, error) {
	var items []Manager
	err := r.db.Table("managers").Order("id").Find(&items).Error
	return items, err
}

func (r *repository) CreateManager(manager *Manager) error {
	return r.db.Create(manager).Error
}

func (r *repository) GetManagerByID(id uint) (*Manager, error) {
	var manager Manager
	err := r.db.First(&manager, id).Error
	return &manager, err
}

func (r *repository) UpdateManager(manager *Manager) error {
	return r.db.Save(manager).Error
}

func (r *repository) DeleteManager(id uint) error {
	return r.db.Delete(&Manager{}, id).Error
}
