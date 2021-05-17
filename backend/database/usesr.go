package database

import (
	"os"
	"u-fes-2021-team-c/model"
)

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) CreateUser(user *model.User) (int, error) {
	if os.Getenv("MODE") == "test" {
		return 1, nil
	}

	result := repo.SqlHandler.Conn.Create(&user)
	if err := result.Error; err != nil {
		return -1, err
	}

	return user.Id, nil
}

func (repo *UserRepository) GetAllUsers() ([]*model.User, error) {
	if os.Getenv("MODE") == "test" {
		return []*model.User{
			{Id: 1, Name: "name", Password: "pass"},
		}, nil
	}

	users := []*model.User{}
	result := repo.SqlHandler.Conn.Find(&users)
	if err := result.Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (repo *UserRepository) GetUserById(userId int) (*model.User, error) {
	if os.Getenv("MODE") == "test" {
		return &model.User{Id: 1, Name: "name", Password: "pass"}, nil
	}

	user := &model.User{}
	result := repo.SqlHandler.Conn.Where("ID = ?", userId).First(&user)
	if err := result.Error; err != nil {
		return nil, err
	}

	return user, nil
}
