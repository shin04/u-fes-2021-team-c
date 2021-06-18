package repository

import "u-fes-2021-team-c/model"

type StudentInfoRepository interface {
	CreateStudentInfo(stundentInfo *model.StudentInfo) (int, error)
	GetAllStudentInfo() ([]*model.StudentInfo, error)
	GetStudentInfoByUserId(userId int) (*model.StudentInfo, error)
}
