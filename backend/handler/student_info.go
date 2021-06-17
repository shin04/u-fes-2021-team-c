package handler

import (
	"errors"
	"log"

	"u-fes-2021-team-c/repository"
	"u-fes-2021-team-c/usecase"

	"github.com/gin-gonic/gin"
)

type StudentInfoHandler struct {
	uc usecase.StudentInfoUsecase
}

func NewStudentinfoHandler(studentinfoRepo repository.StudentInfoRepository) *StudentInfoHandler {
	uc := usecase.StudentInfoUsecase{StudentInfoRepo: studentinfoRepo}

	return &StudentInfoHandler{uc: uc}
}

func (handler *StudentInfoHandler) GetAllStudentInfo(c *gin.Context) {
	studentinfos, err := handler.uc.GetAllStudentInfo()
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	if len(studentinfos) < 1 {
		err = errors.New("studentinfo not found")
		log.Print(err)
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}

	c.JSON(200, studentinfos)
}