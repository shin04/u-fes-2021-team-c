package usecase

import (
	"u-fes-2021-team-c/database"
	"u-fes-2021-team-c/model"
)

type UserUsecase struct {
	UserRepo database.UserRepository
}

func (uc *UserUsecase) registerNewUser(user *model.User) (int, error) {
	userId, err := uc.UserRepo.CreateUser(user)
	if err != nil {
		return -1, err
	}

	return userId, nil
}
