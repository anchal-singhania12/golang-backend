package repository

import (
	"time"

	"gitlab.com/fanligafc-group/fanligafc-backend/pkg/basecontracts"
	"gorm.io/datatypes"
)

// Post represents a post in the fanliga feed
type Post struct {
	basecontracts.BaseModel
	UserID        int64              `gorm:"column:user_id;not null;index:idx_posts_user_created_at_id_active" json:"user_id"`
	Content       string             `gorm:"column:content;type:text;not null" json:"content"`
	Images        datatypes.JSON     `gorm:"column:images;type:jsonb;default:'[]'::jsonb;not null" json:"images"`
	IsDeleted     bool               `gorm:"column:is_deleted;default:false;index:idx_posts_created_at_id_active" json:"is_deleted"`
	IsBlocked     bool               `gorm:"column:is_blocked;default:false;index:idx_posts_created_at_id_active" json:"is_blocked"`
	CreatedAt     time.Time          `gorm:"column:created_at;autoCreateTime:milli;index:idx_posts_created_at_id_active;index:idx_posts_user_created_at_id_active" json:"created_at"`
	UpdatedAt     time.Time          `gorm:"column:updated_at;autoUpdateTime:milli" json:"updated_at"`
	UserName      string             `gorm:"-" json:"-"` // User's name (not stored, fetched from join)
	UserPhoto     string             `gorm:"-" json:"-"` // User's profile image (not stored, fetched from join)
	FavouriteTeam string             `gorm:"-" json:"-"` // User's favourite club name (not stored, fetched from join)
}

// TableName specifies the table name for the Post model
func (Post) TableName() string {
	return "posts"
}

// PostImage represents an image object with URL and display order
type PostImage struct {
	URL   string `json:"url"`
	Order int    `json:"order"`
}
