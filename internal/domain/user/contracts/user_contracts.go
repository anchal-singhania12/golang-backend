package contracts

type FetchUserProfileRequest struct {
	UserID uint `header:"user-id" binding:"required"`
}

type CreateUserRequest struct {
	Name          string `json:"name" binding:"required,min=2,max=100"`       // Full name
	UserName      string `json:"user_name" binding:"required,min=3,max=100"` // Username
	Photo         string `json:"photo,omitempty"`                            // Optional
	Chant         string `json:"chant,omitempty"`                            // Optional
	Bio           string `json:"bio,omitempty" binding:"max=1024"`           // Optional, text up to 1024 chars
	Badge         string `json:"badge,omitempty"`                            // Optional
	BestPlayerID  *uint  `json:"best_player_id,omitempty"`                   // Optional
	BestClubID    *uint  `json:"best_club_id,omitempty"`                     // Optional
	BestManagerID *uint  `json:"best_manager_id,omitempty"`                  // Optional
	BestCountryID *uint  `json:"best_country_id,omitempty"`                  // Optional
}

// UserProfileResponse represents the structured response for user profile
type UserProfileResponse struct {
	ID            uint                `json:"id"`
	JoinedAt      string              `json:"joined_at"`
	UserName      string              `json:"user_name"`
	Name          string              `json:"name"`
	Photo         string              `json:"photo,omitempty"`
	Badge         string              `json:"badge,omitempty"`
	Chant         string              `json:"chant,omitempty"`
	Bio           string              `json:"bio,omitempty"`
	FollowerCount int                 `json:"follower_count"`
	FollowingCount int                `json:"following_count"`
	Favorites     []FavoriteEntity    `json:"favorites"`
}

// FavoriteEntity represents a favorite entity (club, player, manager, country)
type FavoriteEntity struct {
	EntityName string `json:"entity_name"`
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Image      string `json:"image"`
}

// FetchUserPlayersRequest is used to bind the user-id header
type FetchUserPlayersRequest struct {
	UserID uint `header:"user-id" binding:"required"`
}

type AddUserPlayersRequest struct {
	PlayerIDs []uint `json:"player_ids" binding:"required"`
}

type FollowUserRequest struct {
	FollowingID uint `json:"following_id" binding:"required"`
}
