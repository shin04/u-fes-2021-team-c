package usecase

import (
	"u-fes-2021-team-c/model"
	"u-fes-2021-team-c/repository"
)

type StudentInfoUsecase struct {
	StudentInfoRepo repository.StudentInfoRepository
}

func (uc *StudentInfoUsecase) GetAllStudentInfo() ([]*model.StudentInfo, error) {
	studentinfos, err := uc.StudentInfoRepo.GetAllStudentInfo()
	if err != nil {
		return nil, err
	}
	return studentinfos, nil
}
