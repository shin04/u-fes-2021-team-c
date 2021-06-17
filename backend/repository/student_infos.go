package repository

import "u-fes-2021-team-c/model"

type StudentInfoRepository interface {
	GetAllStudentInfo() ([]*model.StudentInfo, error)
	GetStudentInfoByUserId(userId int) (*model.StudentInfo, error)
}
