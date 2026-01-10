package repository

import (
	"time"

	"gitlab.com/fanligafc-group/fanligafc-backend/pkg/basecontracts"
	catalogRepo "gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/catalog/repository"
)

type User struct {
	basecontracts.BaseModel
	Name           string `gorm:"column:name" json:"name"`
	UserName       string `gorm:"column:user_name;not null" json:"user_name"`
	Photo          string `gorm:"column:photo" json:"photo,omitempty"`
	Chant          string `gorm:"column:chant" json:"chant,omitempty"`
	Bio            string `gorm:"column:bio;type:text" json:"bio,omitempty"`
	Badge          string `gorm:"column:badge" json:"badge,omitempty"`
	BestPlayerID   *uint  `gorm:"column:best_player" json:"best_player_id,omitempty"`
	BestClubID     *uint  `gorm:"column:best_club" json:"best_club_id,omitempty"`
	BestManagerID  *uint  `gorm:"column:best_manager" json:"best_manager_id,omitempty"`
	BestCountryID  *uint  `gorm:"column:best_country" json:"best_country_id,omitempty"`
	FollowerCount  int    `gorm:"column:follower_count;default:0" json:"follower_count"`
	FollowingCount int    `gorm:"column:following_count;default:0" json:"following_count"`

	// Relationships - these will be populated when preloaded
	BestPlayer  *catalogRepo.Player  `gorm:"foreignKey:BestPlayerID;references:ID" json:"best_player,omitempty"`
	BestClub    *catalogRepo.Club    `gorm:"foreignKey:BestClubID;references:ID" json:"best_club,omitempty"`
	BestManager *catalogRepo.Manager `gorm:"foreignKey:BestManagerID;references:ID" json:"best_manager,omitempty"`
	BestCountry *catalogRepo.Country `gorm:"foreignKey:BestCountryID;references:ID" json:"best_country,omitempty"`
}

// Player represents a record from the players table.
type Player struct {
	basecontracts.BaseModel
	Name       string `gorm:"column:name;not null" json:"name"`
	PositionID *uint  `gorm:"column:position_id" json:"position_id,omitempty"`
}

// UserPlayer maps users to players via user_players table
type UserPlayer struct {
	basecontracts.BaseModel
	UserID   uint      `gorm:"column:user_id;not null" json:"user_id"`
	PlayerID uint      `gorm:"column:player_id;not null" json:"player_id"`
	AddedAt  time.Time `gorm:"column:added_at;not null;default:now()" json:"added_at"`
}
