package usecase

import (
	"u-fes-2021-team-c/model"
	"u-fes-2021-team-c/repository"
)

type StudentInfoUsecase struct {
	StudentInfoRepo repository.StudentInfoRepository
}

func (uc *StudentInfoUsecase) RegisterNewStudentInfo(userId int, studentName string, studentNumber int) (int, error) {
	studentInfo := &model.StudentInfo{
		UseId:         userId,
		Name:          studentName,
		StudentNumber: studentNumber,
	}

	id, err := uc.StudentInfoRepo.CreateStudentInfo(studentInfo)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (uc *StudentInfoUsecase) GetAllStudentInfo() ([]*model.StudentInfo, error) {
	studentInfos, err := uc.StudentInfoRepo.GetAllStudentInfo()
	if err != nil {
		return nil, err
	}
	return studentInfos, nil
}

func (uc *StudentInfoUsecase) GetStudentInfoByUserId(userId int) (*model.StudentInfo, error) {
	studentInfo, err := uc.StudentInfoRepo.GetStudentInfoByUserId(userId)
	if err != nil {
		return nil, err
	}
	return studentInfo, nil
}
