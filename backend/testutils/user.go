package testutils

import "u-fes-2021-team-c/model"

type FakeUserRepository struct {
	FakeCreateUser  func(user *model.User) (int, error)
	FakeGetAllUsers func() ([]*model.User, error)
	FakeGetUserById func(userId int) (*model.User, error)
}

func (repo FakeUserRepository) CreateUser(user *model.User) (int, error) {
	return repo.FakeCreateUser(user)
}

func (repo FakeUserRepository) GetAllUsers() ([]*model.User, error) {
	return repo.FakeGetAllUsers()
}

func (repo FakeUserRepository) GetUserById(userId int) (*model.User, error) {
	return repo.FakeGetUserById(userId)
}
