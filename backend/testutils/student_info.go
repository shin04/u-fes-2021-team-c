package testutils

import "u-fes-2021-team-c/model"

type FakeStudentInfoRepository struct {
	FakeGetAllStudentInfo       func() ([]*model.StudentInfo, error)
	FakeGetStudentInfoByUSesrID func(usesrId int) (*model.StudentInfo, error)
}

func (repo *FakeStudentInfoRepository) GetAllStudentInfo() ([]*model.StudentInfo, error) {
	return repo.FakeGetAllStudentInfo()
}

func (repo *FakeStudentInfoRepository) GetStudentInfoByUserId(userId int) (*model.StudentInfo, error) {
	return repo.FakeGetStudentInfoByUSesrID(userId)
}
