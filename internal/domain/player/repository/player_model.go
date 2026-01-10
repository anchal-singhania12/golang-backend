package repository

import "gitlab.com/fanligafc-group/fanligafc-backend/pkg/basecontracts"

type Player struct {
	basecontracts.BaseModel
	Name             string `json:"name"`
	PositionID       uint   `json:"position_id"`
	ProviderPlayerID uint   `json:"provider_player_id"`
}
