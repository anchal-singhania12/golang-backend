package contract

type CreatePositionRequest struct {
	ProviderID uint   `json:"provider_id" binding:"required,gt=0"`
	Name       string `json:"name" binding:"required,min=2"`
}

type UpdatePositionRequest struct {
	Name string `json:"name" binding:"required,min=2"`
}
