package repository

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

// PostRepository handles data access for posts
type PostRepository struct {
	db *gorm.DB
}

// NewPostRepository creates a new post repository
func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{
		db: db,
	}
}

// CursorData represents the cursor pagination data
type CursorData struct {
	CreatedAt string `json:"created_at"`
	ID        int64  `json:"id"`
}

// DecodeCursor decodes a base64 encoded cursor into CursorData
func DecodeCursor(cursor string) (*CursorData, error) {
	if cursor == "" {
		return nil, nil
	}

	decodedBytes, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return nil, errors.New("invalid cursor format")
	}

	var cursorData CursorData
	if err := json.Unmarshal(decodedBytes, &cursorData); err != nil {
		return nil, errors.New("invalid cursor data")
	}

	return &cursorData, nil
}

// FetchPostsResponse represents the paginated response structure
type FetchPostsResponse struct {
	Data FetchPostsData `json:"data"`
}

// FetchPostsData contains posts and pagination info
type FetchPostsData struct {
	Posts      []Post     `json:"posts"`
	Pagination Pagination `json:"pagination"`
}

// Pagination contains cursor and has_more flag
type Pagination struct {
	Cursor  string `json:"cursor"`
	HasMore bool   `json:"has_more"`
}

// FetchPosts retrieves posts with cursor-based pagination along with user details
// The query fetches posts ordered by created_at DESC, id DESC
// Joins with users table to get user details
// If cursor is provided, it only fetches posts where (created_at, id) < (cursor.created_at, cursor.id)
func (pr *PostRepository) FetchPosts(ctx context.Context, cursor string, limit int) (*FetchPostsResponse, error) {
	if limit <= 0 || limit > 100 {
		limit = 10
	}

	type PostWithUserDetails struct {
		Post
		UserName     string `gorm:"column:user_name"`
		UserPhoto    string `gorm:"column:user_photo"`
		BestClubID   *uint  `gorm:"column:best_club"`
		BestClubName string `gorm:"column:club"`
	}

	var postsWithDetails []PostWithUserDetails
	query := pr.db.WithContext(ctx).
		Table("posts p").
		Select(
			"p.id, p.user_id, p.content, p.images, p.is_deleted, p.is_blocked, p.created_at, p.updated_at, "+
				"u.user_name, u.photo as user_photo, u.best_club, c.club_image as club",
		).
		Joins("LEFT JOIN users u ON p.user_id = u.id").
		Joins("LEFT JOIN clubs c ON u.best_club = c.id").
		Where("p.is_deleted = ? AND p.is_blocked = ?", false, false).
		Order("p.created_at DESC, p.id DESC")

	// If cursor is provided, apply cursor-based filtering
	if cursor != "" {
		cursorData, err := DecodeCursor(cursor)
		if err != nil {
			return nil, err
		}

		if cursorData != nil {
			// Parse the cursor timestamp
			cursorTime, err := time.Parse(time.RFC3339, cursorData.CreatedAt)
			if err != nil {
				return nil, errors.New("invalid cursor timestamp format")
			}

			// Filter posts: (created_at, id) < (cursor.created_at, cursor.id)
			query = query.Where(
				"(p.created_at < ? OR (p.created_at = ? AND p.id < ?))",
				cursorTime,
				cursorTime,
				cursorData.ID,
			)
		}
	}

	// Fetch limit + 1 posts to determine if there are more posts
	if err := query.Limit(limit + 1).Scan(&postsWithDetails).Error; err != nil {
		return nil, err
	}

	response := &FetchPostsResponse{
		Data: FetchPostsData{
			Posts: []Post{},
			Pagination: Pagination{
				Cursor:  "",
				HasMore: false,
			},
		},
	}

	// If we got more posts than the limit, there are more posts available
	hasMore := false
	if len(postsWithDetails) > limit {
		hasMore = true
		postsWithDetails = postsWithDetails[:limit] // Return only limit number of posts
	}

	// Convert to Post models
	posts := make([]Post, len(postsWithDetails))
	for i, pwd := range postsWithDetails {
		posts[i] = pwd.Post
		// Store user details as metadata (to be accessed in usecase)
		posts[i].UserName = pwd.UserName
		posts[i].UserPhoto = pwd.UserPhoto
		posts[i].FavouriteTeam = pwd.BestClubName
	}

	response.Data.Posts = posts
	response.Data.Pagination.HasMore = hasMore

	// Generate cursor for next pagination
	if len(posts) > 0 && hasMore {
		lastPost := posts[len(posts)-1]
		newCursorData := CursorData{
			CreatedAt: lastPost.CreatedAt.Format(time.RFC3339),
			ID:        int64(lastPost.ID),
		}
		cursorBytes, _ := json.Marshal(newCursorData)
		response.Data.Pagination.Cursor = base64.StdEncoding.EncodeToString(cursorBytes)
	}

	return response, nil
}

// CreatePost creates a new post
func (pr *PostRepository) CreatePost(ctx context.Context, post *Post) error {
	return pr.db.WithContext(ctx).Create(post).Error
}

// GetPostByID retrieves a post by its ID
func (pr *PostRepository) GetPostByID(ctx context.Context, postID int64) (*Post, error) {
	var post Post
	if err := pr.db.WithContext(ctx).Where("id = ? AND is_deleted = ? AND is_blocked = ?", postID, false, false).First(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

// GetPostsByUserID retrieves all active posts from a specific user
func (pr *PostRepository) GetPostsByUserID(ctx context.Context, userID int64) ([]Post, error) {
	var posts []Post
	if err := pr.db.WithContext(ctx).
		Where("user_id = ? AND is_deleted = ? AND is_blocked = ?", userID, false, false).
		Order("created_at DESC, id DESC").
		Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

// UpdatePost updates an existing post
func (pr *PostRepository) UpdatePost(ctx context.Context, post *Post) error {
	return pr.db.WithContext(ctx).
		Model(post).
		Updates(post).
		Error
}

// DeletePost soft deletes a post
func (pr *PostRepository) DeletePost(ctx context.Context, postID int64) error {
	return pr.db.WithContext(ctx).
		Model(&Post{}).
		Where("id = ?", postID).
		Update("is_deleted", true).
		Error
}

// BlockPost blocks a post by fanliga
func (pr *PostRepository) BlockPost(ctx context.Context, postID int64) error {
	return pr.db.WithContext(ctx).
		Model(&Post{}).
		Where("id = ?", postID).
		Update("is_blocked", true).
		Error
}
