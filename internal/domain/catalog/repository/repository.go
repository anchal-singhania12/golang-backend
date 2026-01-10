package repository

import "gorm.io/gorm"

type Operations interface {
	// Positions
	ListPositions() ([]Position, error)
	CreatePosition(position *Position) error
	GetPositionByID(id uint) (*Position, error)
	UpdatePosition(position *Position) error
	DeletePosition(id uint) error

	// Clubs
	ListClubs() ([]Club, error)
	CreateClub(club *Club) error
	GetClubByID(id uint) (*Club, error)
	UpdateClub(club *Club) error
	DeleteClub(id uint) error

	// Managers
	ListManagers() ([]Manager, error)
	CreateManager(manager *Manager) error
	GetManagerByID(id uint) (*Manager, error)
	UpdateManager(manager *Manager) error
	DeleteManager(id uint) error

	// Countries
	ListCountries() ([]Country, error)
	CreateCountry(country *Country) error
	GetCountryByID(id uint) (*Country, error)
	UpdateCountry(country *Country) error
	DeleteCountry(id uint) error

	// Players
	ListPlayers() ([]Player, error)
	CreatePlayer(player *Player) error
	GetPlayerByID(id uint) (*Player, error)
	UpdatePlayer(player *Player) error
	DeletePlayer(id uint) error

	// Entities
	ListEntities() ([]Entity, error)
	CreateEntity(entity *Entity) error
	GetEntityByID(id uint) (*Entity, error)
	UpdateEntity(entity *Entity) error
	DeleteEntity(id uint) error

	// Ranked Entities
	ListRankedEntities() ([]RankedEntity, error)
	CreateRankedEntity(rankedEntity *RankedEntity) error
	GetRankedEntityByID(id uint) (*RankedEntity, error)
	UpdateRankedEntity(rankedEntity *RankedEntity) error
	DeleteRankedEntity(id uint) error
}

type repository struct{ db *gorm.DB }

func NewRepository(db *gorm.DB) Operations { return &repository{db: db} }
