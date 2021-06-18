package testutils

import "u-fes-2021-team-c/model"

type FakeStudentInfoRepository struct {
	FakeCreateStudentInfo       func(studentInfo *model.StudentInfo) (int, error)
	FakeGetAllStudentInfo       func() ([]*model.StudentInfo, error)
	FakeGetStudentInfoByUSesrID func(usesrId int) (*model.StudentInfo, error)
}

func (repo *FakeStudentInfoRepository) CreateStudentInfo(studentInfo *model.StudentInfo) (int, error) {
	return repo.FakeCreateStudentInfo(studentInfo)
}

func (repo *FakeStudentInfoRepository) GetAllStudentInfo() ([]*model.StudentInfo, error) {
	return repo.FakeGetAllStudentInfo()
}

func (repo *FakeStudentInfoRepository) GetStudentInfoByUserId(userId int) (*model.StudentInfo, error) {
	return repo.FakeGetStudentInfoByUSesrID(userId)
}
