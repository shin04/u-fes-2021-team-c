package database

import (
	"u-fes-2021-team-c/model"
	"u-fes-2021-team-c/repository"
)

type StudentInfoRepository struct {
	SqlHandler
}

func NewStudentRepository(sqlHandler SqlHandler) repository.StudentInfoRepository {
	return &StudentInfoRepository{sqlHandler}
}

func (repo *StudentInfoRepository) GetAllStudentInfo() ([]*model.StudentInfo, error) {
	studentinfos := []*model.StudentInfo{}

	result := repo.SqlHandler.Conn.Find(&studentinfos)
	if err := result.Error; err != nil {
		return nil, err
	}
	return studentinfos, nil
}
