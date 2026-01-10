package contract

type PlayerRequest struct {
	Name             string `json:"name" binding:"required,min=2,max=100"`
	PositionID       uint   `json:"position_id" binding:"required,gt=0"`
	ProviderPlayerID uint   `json:"provider_player_id" binding:"required,gt=0"`
}

type FetchUserTeamRequest struct {
	UserID uint `header:"user-id" binding:"required"`
}
