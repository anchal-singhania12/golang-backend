package usecase

import (
	"gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/user/errors"
	"gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/user/repository"
	"strconv"

	"gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/user/contracts"
)

func (u *userUseCase) FetchUserProfileByID(id uint) (*contracts.UserProfileResponse, error) {
	data, err := u.repository.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	// Convert created_at to epoch string
	joinedAtEpoch := strconv.FormatInt(data.CreatedAt.Unix(), 10)

	// Build favorites array from best entities
	var favorites []contracts.FavoriteEntity

	// Add best club if exists
	if data.BestClub != nil {
		favorites = append(favorites, contracts.FavoriteEntity{
			EntityName: "club",
			ID:         data.BestClub.ID,
			Name:       data.BestClub.ClubName,
			Image:      data.BestClub.ClubImage,
		})
	}

	// Add best player if exists
	if data.BestPlayer != nil {
		favorites = append(favorites, contracts.FavoriteEntity{
			EntityName: "player",
			ID:         data.BestPlayer.ID,
			Name:       data.BestPlayer.Name,
			Image:      "", // Players don't have images in current schema
		})
	}

	// Add best country if exists
	if data.BestCountry != nil {
		favorites = append(favorites, contracts.FavoriteEntity{
			EntityName: "country",
			ID:         data.BestCountry.ID,
			Name:       data.BestCountry.CountryName,
			Image:      data.BestCountry.CountryImage,
		})
	}

	// Add best manager if exists
	if data.BestManager != nil {
		favorites = append(favorites, contracts.FavoriteEntity{
			EntityName: "manager",
			ID:         data.BestManager.ID,
			Name:       data.BestManager.ManagerName,
			Image:      data.BestManager.ManagerImage,
		})
	}

	// Create response
	response := &contracts.UserProfileResponse{
		ID:             data.ID,
		JoinedAt:       joinedAtEpoch,
		UserName:       data.UserName,
		Name:           data.Name,
		Photo:          data.Photo,
		Badge:          data.Badge,
		Chant:          data.Chant,
		Bio:            data.Bio,
		FollowerCount:  data.FollowerCount,
		FollowingCount: data.FollowingCount,
		Favorites:      favorites,
	}

	return response, nil
}

func (u *userUseCase) FetchPlayersByUserID(userID uint) ([]repository.Player, error) {
	data, err := u.repository.GetPlayersByUserID(userID)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.ErrorUserPlayerMappingNotFound
	}
	return data, nil
}
