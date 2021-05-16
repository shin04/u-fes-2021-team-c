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
