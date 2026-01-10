package usecase

import (
	"log"

	"gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/player/contract"
	"gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/player/repository"
)

func (u *playerUseCase) FetchUserTeam(userID uint) (*contract.UserTeamResponse, error) {
	players, err := u.repository.FetchUserTeam(userID)
	if err != nil {
		log.Printf("Error fetching user team: %v, userID : %v", err, userID)
		return nil, err
	}

	userTeam := createUserTeamFetchResponse(players)
	return userTeam, nil
}

func createUserTeamFetchResponse(players []repository.Player) *contract.UserTeamResponse {
	var response contract.UserTeamResponse

	var team []contract.PlayerResponse
	for _, player := range players {
		playerResp := contract.PlayerResponse{
			ID:         player.ID,
			Name:       player.Name,
			PositionID: player.PositionID,
		}
		team = append(team, playerResp)
	}

	response.Players = team
	return &response
}
