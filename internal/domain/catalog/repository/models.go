package repository

import "gitlab.com/fanligafc-group/fanligafc-backend/pkg/basecontracts"

type Position struct {
	basecontracts.BaseModel
	ProviderID   *int   `gorm:"column:provider_id" json:"provider_id,omitempty"`
	PositionName string `gorm:"column:position_name;not null" json:"position_name"`
}

type Club struct {
	basecontracts.BaseModel
	ClubName  string `gorm:"column:club_name" json:"club_name,omitempty"`
	ClubImage string `gorm:"column:club_image" json:"club_image,omitempty"`
}

type Manager struct {
	basecontracts.BaseModel
	ManagerName  string `gorm:"column:manager_name" json:"manager_name,omitempty"`
	ManagerImage string `gorm:"column:manager_image" json:"manager_image,omitempty"`
}

type Country struct {
	basecontracts.BaseModel
	CountryName  string `gorm:"column:country_name" json:"country_name,omitempty"`
	CountryImage string `gorm:"column:country_image" json:"country_image,omitempty"`
}

type Player struct {
	basecontracts.BaseModel
	Name       string `gorm:"column:name;not null" json:"name"`
	PositionID *uint  `gorm:"column:position_id" json:"position_id,omitempty"`
}

type Entity struct {
	basecontracts.BaseModel
	EntityName string `gorm:"column:entity_name;unique" json:"entity_name"`
}

func (Entity) TableName() string {
	return "entity"
}

type RankedEntity struct {
	basecontracts.BaseModel

	EntityID  *uint `gorm:"column:entity_id" json:"entity_id,omitempty"`
	PlayerID  *uint `gorm:"column:player_id" json:"player_id,omitempty"`
	ClubID    *uint `gorm:"column:club_id" json:"club_id,omitempty"`
	ManagerID *uint `gorm:"column:manager_id" json:"manager_id,omitempty"`
	CountryID *uint `gorm:"column:country_id" json:"country_id,omitempty"`
}
