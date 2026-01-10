package repository

import (
	"log"

	"gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/user/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Operations interface {
	GetUserByID(id uint) (*User, error)
	SaveUser(user *User) error
	GetPlayersByUserID(userID uint) ([]Player, error)
	AddPlayersForUser(userID uint, playerIDs []uint) (int64, error)
	FollowUser(followerID, followingID uint) error
	UnfollowUser(followerID, followingID uint) error
}

type repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) Operations {
	return &repository{
		db: db,
	}
}

func (r *repository) GetUserByID(id uint) (*User, error) {
	var user User
	dbResp := r.db.Preload("BestPlayer").
		Preload("BestClub").
		Preload("BestManager").
		Preload("BestCountry").
		First(&user, id)

	if dbResp.Error != nil {
		if dbResp.Error == gorm.ErrRecordNotFound {
			log.Printf("User with ID %d not found", id)
			return nil, errors.ErrorUserNotFound
		}

		log.Printf("Error fetching user by ID %d: %v", id, dbResp.Error)
		return nil, errors.ErrorGenericInternalServerError
	}
	return &user, nil
}

func (r *repository) SaveUser(user *User) error {
	return r.db.Create(user).Error
}

// GetPlayersByUserID returns all players mapped to a given user via user_players
func (r *repository) GetPlayersByUserID(userID uint) ([]Player, error) {
	// Join user_players with players to fetch player details
	var players []Player
	// here up is the user_players table alias
	tx := r.db.Table("user_players up").
		Select("p.id, p.created_at, p.updated_at, p.name, p.position_id").
		Joins("JOIN players p ON p.id = up.player_id").
		Where("up.user_id = ?", userID).
		Scan(&players)

	if tx.Error != nil {
		log.Printf("error fetching players for user %d: %v", userID, tx.Error)
		return nil, errors.ErrorGenericInternalServerError
	}
	if tx.RowsAffected == 0 {
		return nil, errors.ErrorUserPlayerMappingNotFound
	}
	return players, nil
}

// AddPlayersForUser inserts user-player mappings; ignores duplicates
func (r *repository) AddPlayersForUser(userID uint, playerIDs []uint) (int64, error) {
	if len(playerIDs) == 0 {
		return 0, nil
	}

	// Validate player IDs exist
	var existing []uint
	if err := r.db.Table("players").
		Where("id IN ?", playerIDs).
		Pluck("id", &existing).Error; err != nil {
		log.Printf("error validating player ids: %v", err)
		return 0, errors.ErrorGenericInternalServerError
	}

	if len(existing) == 0 {
		return 0, errors.ErrorBadRequest
	}

	// Build rows for bulk insert
	type row struct{ UserID, PlayerID uint }
	rows := make([]row, 0, len(existing))
	for _, pid := range existing {
		rows = append(rows, row{UserID: userID, PlayerID: pid})
	}

	// Use GORM create with On Conflict Do Nothing
	tx := r.db.Table("user_players").
		Clauses(clause.OnConflict{DoNothing: true}).
		Create(&rows)

	if tx.Error != nil {
		log.Printf("error adding players for user %d: %v", userID, tx.Error)
		return 0, errors.ErrorGenericInternalServerError
	}
	return tx.RowsAffected, nil
}

func (r *repository) FollowUser(followerID, followingID uint) error {
	if followerID == followingID {
		return errors.ErrorCannotFollowSelf
	}
	type row struct{ FollowerID, FollowingID uint }
	rec := row{FollowerID: followerID, FollowingID: followingID}
	tx := r.db.Table("user_follows").
		Clauses(clause.OnConflict{DoNothing: true}).
		Create(&rec)
	if tx.Error != nil {
		log.Printf("error follow user: %v", tx.Error)
		return errors.ErrorGenericInternalServerError
	}
	return nil
}

func (r *repository) UnfollowUser(followerID, followingID uint) error {
	tx := r.db.Table("user_follows").
		Where("follower_id = ? AND following_id = ?", followerID, followingID).
		Delete(nil)
	if tx.Error != nil {
		log.Printf("error unfollow user: %v", tx.Error)
		return errors.ErrorGenericInternalServerError
	}
	return nil
}
