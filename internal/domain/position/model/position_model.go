package position

import "time"

type PositionModel struct {
	ID         uint      `gorm:"primaryKey"`
	ProviderID uint      `gorm:"not null"`
	Name       string    `gorm:"not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
	DeletedAt  time.Time `gorm:"autoDeleteTime"`
}
