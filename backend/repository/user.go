package repository

import "u-fes-2021-team-c/model"

type UserRepository interface {
	CreateUser(user *model.User) (int, error)
	GetAllUsers() ([]*model.User, error)
	GetUserById(userId int) (*model.User, error)
}
