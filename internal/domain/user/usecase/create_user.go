package usecase

import "gitlab.com/fanligafc-group/fanligafc-backend/internal/domain/user/repository"

func (u *userUseCase) CreateUser(user repository.User) (*repository.User, error) {
	err := u.repository.SaveUser(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userUseCase) AddPlayersForUser(userID uint, playerIDs []uint) error {
	_, err := u.repository.AddPlayersForUser(userID, playerIDs)
	return err
}

func (u *userUseCase) FollowUser(followerID, followingID uint) error {
	return u.repository.FollowUser(followerID, followingID)
}

func (u *userUseCase) UnfollowUser(followerID, followingID uint) error {
	return u.repository.UnfollowUser(followerID, followingID)
}
