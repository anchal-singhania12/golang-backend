package usecase

import repo "gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/catalog/repository"

type CatalogUseCase interface {
	// Positions
	ListPositions() ([]repo.Position, error)
	CreatePosition(*repo.Position) (*repo.Position, error)
	GetPositionByID(uint) (*repo.Position, error)
	UpdatePosition(*repo.Position) (*repo.Position, error)
	DeletePosition(uint) error

	// Clubs
	ListClubs() ([]repo.Club, error)
	CreateClub(*repo.Club) (*repo.Club, error)
	GetClubByID(uint) (*repo.Club, error)
	UpdateClub(*repo.Club) (*repo.Club, error)
	DeleteClub(uint) error

	// Managers
	ListManagers() ([]repo.Manager, error)
	CreateManager(*repo.Manager) (*repo.Manager, error)
	GetManagerByID(uint) (*repo.Manager, error)
	UpdateManager(*repo.Manager) (*repo.Manager, error)
	DeleteManager(uint) error

	// Countries
	ListCountries() ([]repo.Country, error)
	CreateCountry(*repo.Country) (*repo.Country, error)
	GetCountryByID(uint) (*repo.Country, error)
	UpdateCountry(*repo.Country) (*repo.Country, error)
	DeleteCountry(uint) error

	// Players
	ListPlayers() ([]repo.Player, error)
	CreatePlayer(*repo.Player) (*repo.Player, error)
	GetPlayerByID(uint) (*repo.Player, error)
	UpdatePlayer(*repo.Player) (*repo.Player, error)
	DeletePlayer(uint) error

	// Entities
	ListEntities() ([]repo.Entity, error)
	CreateEntity(*repo.Entity) (*repo.Entity, error)
	GetEntityByID(uint) (*repo.Entity, error)
	UpdateEntity(*repo.Entity) (*repo.Entity, error)
	DeleteEntity(uint) error

	// Ranked Entities
	ListRankedEntities() ([]repo.RankedEntity, error)
	CreateRankedEntity(*repo.RankedEntity) (*repo.RankedEntity, error)
	GetRankedEntityByID(uint) (*repo.RankedEntity, error)
	UpdateRankedEntity(*repo.RankedEntity) (*repo.RankedEntity, error)
	DeleteRankedEntity(uint) error
}

type catalogUseCase struct{ repository repo.Operations }

func NewCatalogUseCase(r repo.Operations) CatalogUseCase { return &catalogUseCase{repository: r} }
