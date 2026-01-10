package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/posts/repository"
)

type PostRepository interface {
	FetchPosts(ctx context.Context, cursor string, limit int) (*repository.FetchPostsResponse, error)
}

type FetchPostsUsecase struct {
	repository PostRepository
}

// API Response Types
type UserDetails struct {
	UserID        int64  `json:"user_id"`
	Name          string `json:"name"`
	ProfileImage  string `json:"profile_image"`
	FavouriteTeam string `json:"favourite_team"`
}

type Image struct {
	URL   string `json:"url"`
	Order int    `json:"order"`
}

type Content struct {
	Text   string  `json:"text"`
	Images []Image `json:"images"`
}

type Interactions struct {
	LikeCount    string `json:"like_count"`
	CommentCount string `json:"comment_count"`
	ShareCount   string `json:"share_count"`
}

type Post struct {
	PostID       int64        `json:"post_id"`
	UserDetails  UserDetails  `json:"user_details"`
	Content      Content      `json:"content"`
	Timestamp    string       `json:"timestamp"`
	Interactions Interactions `json:"interactions"`
}

type Pagination struct {
	Cursor  string `json:"cursor"`
	HasMore bool   `json:"has_more"`
}

type PostsData struct {
	Posts      []Post     `json:"posts"`
	Pagination Pagination `json:"pagination"`
}

type PostsResponse struct {
	Data PostsData `json:"data"`
}

func NewFetchPostsUsecase(repository PostRepository) *FetchPostsUsecase {
	return &FetchPostsUsecase{
		repository: repository,
	}
}

func (u *FetchPostsUsecase) CommonHomeFeed(c *gin.Context) (interface{}, error) {
	cursor := c.DefaultQuery("cursor", "")

	limit := 10
	if limitQuery := c.Query("limit"); limitQuery != "" {
		if parsedLimit, err := strconv.Atoi(limitQuery); err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	ctx := c.Request.Context()

	// Fetch posts from repository
	repoResponse, err := u.repository.FetchPosts(ctx, cursor, limit)
	if err != nil {
		return nil, errors.New("failed to fetch posts")
	}

	// Transform repository response to API response
	apiPosts := make([]Post, 0, len(repoResponse.Data.Posts))
	for _, repoPost := range repoResponse.Data.Posts {
		// Deserialize images from JSONB
		var images []Image
		if len(repoPost.Images) > 0 {
			if err := json.Unmarshal(repoPost.Images, &images); err != nil {
				images = []Image{} // Default to empty array on error
			} else {
				// Sort images by order field in ascending order
				sort.Slice(images, func(i, j int) bool {
					return images[i].Order < images[j].Order
				})
			}
		}

		apiPost := Post{
			PostID: int64(repoPost.ID),
			UserDetails: UserDetails{
				UserID:        repoPost.UserID,
				Name:          repoPost.UserName,
				ProfileImage:  repoPost.UserPhoto,
				FavouriteTeam: repoPost.FavouriteTeam,
			},
			Content: Content{
				Text:   repoPost.Content,
				Images: images,
			},
			Timestamp: repoPost.CreatedAt.Format("2006-01-02T15:04:05Z"),
			Interactions: Interactions{
				LikeCount:    "0",
				CommentCount: "1.2k",
				ShareCount:   "500",
			},
		}
		apiPosts = append(apiPosts, apiPost)
	}

	response := &PostsResponse{
		Data: PostsData{
			Posts: apiPosts,
			Pagination: Pagination{
				Cursor:  repoResponse.Data.Pagination.Cursor,
				HasMore: repoResponse.Data.Pagination.HasMore,
			},
		},
	}

	return response, nil
}
