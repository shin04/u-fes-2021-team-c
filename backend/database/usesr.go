package database

import "u-fes-2021-team-c/model"

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) CreateUser(user *model.User) (int, error) {
	result := repo.SqlHandler.Conn.Create(&user)
	if err := result.Error; err != nil {
		return -1, err
	}

	return user.Id, nil
}

func (repo *UserRepository) GetAllUsers() ([]*model.User, error) {
	users := []*model.User{}
	result := repo.SqlHandler.Conn.Find(&users)
	if err := result.Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (repo *UserRepository) GetUserById(userId int) (*model.User, error) {
	user := &model.User{}
	result := repo.SqlHandler.Conn.Where("ID = ?", userId).First(&user)
	if err := result.Error; err != nil {
		return nil, err
	}

	return user, nil
}
