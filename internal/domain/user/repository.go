package user

import "gorm.io/gorm"

type Repository interface {
	GetUserByID(id uint) (*UserModel, error)
	SaveUser(user *UserModel) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetUserByID(id uint) (*UserModel, error) {
	var user UserModel
	err := r.db.Where("id = ?", id).First(&user).Error

	return &user, err
}

func (r *repository) SaveUser(user *UserModel) error {
	return r.db.Create(user).Error
}
