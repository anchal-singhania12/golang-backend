package usecase

func (uc *playerUseCase) DeletePlayer(id uint) error {
	return uc.repository.Delete(id)
}
