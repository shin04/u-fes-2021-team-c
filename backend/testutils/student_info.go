package testutils

import "u-fes-2021-team-c/model"

type FakeStudentInfoRepository struct {
	FakeGetAllStudentInfo func() ([]*model.StudentInfo, error)
}

func (repo *FakeStudentInfoRepository) GetAllStudentInfo() ([]*model.StudentInfo, error) {
	return repo.FakeGetAllStudentInfo()
}
