package usecase

import (
	"u-fes-2021-team-c/database"
	"u-fes-2021-team-c/model"
)

type UserUsecase struct {
	UserRepo database.UserRepository
}

func (uc *UserUsecase) RegisterNewUser(user *model.User) (int, error) {
	userId, err := uc.UserRepo.CreateUser(user)
	if err != nil {
		return -1, err
	}
	return userId, nil
}

func (uc *UserUsecase) GetAllUsers() ([]*model.User, error) {
	users, err := uc.UserRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (uc *UserUsecase) GetUserById(userId int) (*model.User, error) {
	user, err := uc.UserRepo.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
