package contract

type PlayerResponse struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	PositionID uint   `json:"position_id"`
}

type UserTeamResponse struct {
	Players []PlayerResponse `json:"players"`
}
