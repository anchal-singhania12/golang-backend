package repository

import (
	"gorm.io/gorm"
)

type PlayerRepository interface {
	Create(player *Player) error
	Update(player *Player) error
	Delete(id uint) error
	FindByID(id uint) (*Player, error)
	FindByPositionID(positionID uint) ([]*Player, error)
	FetchUserTeam(userID uint) ([]Player, error)
}

type playerRepository struct {
	db *gorm.DB
}

func NewPlayerRepository(db *gorm.DB) PlayerRepository {
	return &playerRepository{db: db}
}

func (r *playerRepository) Create(player *Player) error {
	return r.db.Create(player).Error
}

func (r *playerRepository) Update(player *Player) error {
	return r.db.Save(player).Error
}

func (r *playerRepository) Delete(id uint) error {
	return r.db.Delete(&Player{}, id).Error
}

func (r *playerRepository) FindByID(id uint) (*Player, error) {
	var player Player
	err := r.db.First(&player, id).Error
	return &player, err
}

func (r *playerRepository) FindByPositionID(positionID uint) ([]*Player, error) {
	var players []*Player
	err := r.db.Where("position_id = ?", positionID).Find(&players).Error
	return players, err
}

func(r *playerRepository) FetchUserTeam(userID uint) ([]Player, error) {
	var players []Player
	err := r.db.Joins("JOIN user_teams ON user_teams.player_id = players.id").
		Where("user_teams.user_id = ?", userID).
		Find(&players).Error
	return players, err
}
