package usecase

import (
	"u-fes-2021-team-c/model"
	"u-fes-2021-team-c/repository"
)

type UserUsecase struct {
	UserRepo repository.UserRepository
}

func (uc *UserUsecase) RegisterNewUser(userName string, password string) (int, error) {
	user := &model.User{
		Name:     userName,
		Password: password,
	}

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

func (uc *UserUsecase) Login(name string, password string) (*model.User, error) {
	user, err := uc.UserRepo.GetUserByNameAndPasssword(name, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
