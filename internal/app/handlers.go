package app

import (
	"gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/health"
	playerHandler "gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/player/handler"
	playerRepository "gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/player/repository"
	playerUseCase "gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/player/usecase"
	userHandler "gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/user/handler"
	userRepository "gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/user/repository"
	userUsecase "gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/user/usecase"

	catalogHandler "gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/catalog/handler"
	catalogRepository "gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/catalog/repository"
	catalogUsecase "gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/catalog/usecase"

	postsHandler "gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/posts/handler"
	postsRepository "gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/posts/repository"
	postsUsecase "gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/posts/usecase"
)

type Handlers struct {
	HealthHandler  health.HealthHandler
	UserHandler    userHandler.UserHandler
	CatalogHandler catalogHandler.Handler
	PlayerHandler  playerHandler.PlayerHandler
	PostsHandler   *postsHandler.FetchPostsHandler
}

func InitiateHandlers(deps *Dependencies) *Handlers {
	//health
	healthHandler := health.NewHealthHandler(deps.db)

	// User initialization
	userRepository := userRepository.NewUserRepository(deps.db)
	userUseCase := userUsecase.NewUserUseCase(deps.cfg, userRepository)
	userHandler := userHandler.NewUserHandler(userUseCase)

	// Catalog initialization
	catalogRepo := catalogRepository.NewRepository(deps.db)
	catalogUC := catalogUsecase.NewCatalogUseCase(catalogRepo)
	catHandler := catalogHandler.NewHandler(catalogUC)

	// Players initialization
	playerRepository := playerRepository.NewPlayerRepository(deps.db)
	playerUseCase := playerUseCase.NewPlayerUseCase(deps.cfg, playerRepository)
	playerHandler := playerHandler.NewPlayerHandler(playerUseCase)

	// Posts initialization
	postsRepo := postsRepository.NewPostRepository(deps.db)
	postsUC := postsUsecase.NewFetchPostsUsecase(postsRepo)
	postsHandler := postsHandler.NewFetchPostsHandler(postsUC)

	return &Handlers{
		HealthHandler:  *healthHandler,
		UserHandler:    userHandler,
		CatalogHandler: catHandler,
		PlayerHandler:  playerHandler,
		PostsHandler:   postsHandler,
	}
}
