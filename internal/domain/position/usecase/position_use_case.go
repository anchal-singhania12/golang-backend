package usecase

import (
	"context"
	"errors"

	position "gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/position/model"
	"gorm.io/gorm"
)

type PositionUseCase interface {
	Create(ctx context.Context, pos *position.PositionModel) error
	GetByID(ctx context.Context, id uint) (*position.PositionModel, error)
	Update(ctx context.Context, id uint, name string) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context) ([]position.PositionModel, error)
}

type positionUseCase struct {
	db *gorm.DB
}

func NewPositionUseCase(db *gorm.DB) PositionUseCase {
	return &positionUseCase{db}
}

func (uc *positionUseCase) Create(ctx context.Context, pos *position.PositionModel) error {
	return uc.db.WithContext(ctx).Create(pos).Error
}

func (uc *positionUseCase) GetByID(ctx context.Context, id uint) (*position.PositionModel, error) {
	var pos position.PositionModel
	err := uc.db.WithContext(ctx).First(&pos, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &pos, err
}

func (uc *positionUseCase) Update(ctx context.Context, id uint, name string) error {
	return uc.db.WithContext(ctx).Model(&position.PositionModel{}).
		Where("id = ?", id).
		Update("name", name).Error
}

func (uc *positionUseCase) Delete(ctx context.Context, id uint) error {
	return uc.db.WithContext(ctx).Delete(&position.PositionModel{}, id).Error
}

func (uc *positionUseCase) List(ctx context.Context) ([]position.PositionModel, error) {
	var positions []position.PositionModel
	err := uc.db.WithContext(ctx).Find(&positions).Error
	return positions, err
}
