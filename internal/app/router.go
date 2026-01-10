package app

import (
	"github.com/gin-gonic/gin"
)

func InitializeRoutesEngine(handlers *Handlers) *gin.Engine {
	r := gin.Default()

	//ping
	r.GET("/ping", handlers.HealthHandler.Check)

	//user Routes
	userRoutes := r.Group("/user")
	{
		userRoutes.GET("/profile", handlers.UserHandler.FetchUserProfile)
		userRoutes.POST("/create", handlers.UserHandler.CreateUser)
		userRoutes.GET("/players", handlers.UserHandler.FetchUserPlayers)
		userRoutes.POST("/players", handlers.UserHandler.AddUserPlayers)
		userRoutes.POST("/follow", handlers.UserHandler.FollowUser)
		userRoutes.POST("/unfollow", handlers.UserHandler.UnfollowUser)
	}

	// catalog routes
	catalog := r.Group("/catalog")
	{
		// Positions CRUD
		catalog.GET("/positions", handlers.CatalogHandler.ListPositions)
		catalog.POST("/positions", handlers.CatalogHandler.CreatePosition)
		catalog.GET("/positions/:id", handlers.CatalogHandler.GetPosition)
		catalog.PUT("/positions/:id", handlers.CatalogHandler.UpdatePosition)
		catalog.DELETE("/positions/:id", handlers.CatalogHandler.DeletePosition)

		// Clubs CRUD
		catalog.GET("/clubs", handlers.CatalogHandler.ListClubs)
		catalog.POST("/clubs", handlers.CatalogHandler.CreateClub)
		catalog.GET("/clubs/:id", handlers.CatalogHandler.GetClub)
		catalog.PUT("/clubs/:id", handlers.CatalogHandler.UpdateClub)
		catalog.DELETE("/clubs/:id", handlers.CatalogHandler.DeleteClub)

		// Managers CRUD
		catalog.GET("/managers", handlers.CatalogHandler.ListManagers)
		catalog.POST("/managers", handlers.CatalogHandler.CreateManager)
		catalog.GET("/managers/:id", handlers.CatalogHandler.GetManager)
		catalog.PUT("/managers/:id", handlers.CatalogHandler.UpdateManager)
		catalog.DELETE("/managers/:id", handlers.CatalogHandler.DeleteManager)

		// Countries CRUD
		catalog.GET("/countries", handlers.CatalogHandler.ListCountries)
		catalog.POST("/countries", handlers.CatalogHandler.CreateCountry)
		catalog.GET("/countries/:id", handlers.CatalogHandler.GetCountry)
		catalog.PUT("/countries/:id", handlers.CatalogHandler.UpdateCountry)
		catalog.DELETE("/countries/:id", handlers.CatalogHandler.DeleteCountry)

		// Players CRUD
		catalog.GET("/players", handlers.CatalogHandler.ListPlayers)
		catalog.POST("/players", handlers.CatalogHandler.CreatePlayer)
		catalog.GET("/players/:id", handlers.CatalogHandler.GetPlayer)
		catalog.PUT("/players/:id", handlers.CatalogHandler.UpdatePlayer)
		catalog.DELETE("/players/:id", handlers.CatalogHandler.DeletePlayer)

		// Entities CRUD
		catalog.GET("/entities", handlers.CatalogHandler.ListEntities)
		catalog.POST("/entities", handlers.CatalogHandler.CreateEntity)
		catalog.GET("/entities/:id", handlers.CatalogHandler.GetEntity)
		catalog.PUT("/entities/:id", handlers.CatalogHandler.UpdateEntity)
		catalog.DELETE("/entities/:id", handlers.CatalogHandler.DeleteEntity)

		// Ranked Entities CRUD
		catalog.GET("/ranked-entities", handlers.CatalogHandler.ListRankedEntities)
		catalog.POST("/ranked-entities", handlers.CatalogHandler.CreateRankedEntity)
		catalog.GET("/ranked-entities/:id", handlers.CatalogHandler.GetRankedEntity)
		catalog.PUT("/ranked-entities/:id", handlers.CatalogHandler.UpdateRankedEntity)
		catalog.DELETE("/ranked-entities/:id", handlers.CatalogHandler.DeleteRankedEntity)
	}

	// player Routes
	playerRoutes := r.Group("/player")
	{

		playerRoutes.POST("/", handlers.PlayerHandler.CreatePlayer)
		playerRoutes.PUT("/:id", handlers.PlayerHandler.UpdatePlayer)
		playerRoutes.DELETE("/:id", handlers.PlayerHandler.DeletePlayer)
		playerRoutes.GET("/:id", handlers.PlayerHandler.GetPlayerByID)
		playerRoutes.GET("/position/:position_id", handlers.PlayerHandler.GetPlayersByPositionID)
		playerRoutes.GET("/user-team", handlers.PlayerHandler.FetchUserTeam)
	}

	postsRoutes := r.Group("/posts")
	{
		postsRoutes.GET("/home/feed", handlers.PostsHandler.FetchCommonHomeFeed)
	}

	return r
}
