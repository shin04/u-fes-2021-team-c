package handler

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"u-fes-2021-team-c/repository"
	"u-fes-2021-team-c/usecase"

	"github.com/gin-gonic/gin"
)

type StudentInfoHandler struct {
	uc usecase.StudentInfoUsecase
}

type RegisterStudentInfoReq struct {
	UserId        int
	Name          string
	StudentNumber int
}

func NewStudentinfoHandler(studentinfoRepo repository.StudentInfoRepository) *StudentInfoHandler {
	uc := usecase.StudentInfoUsecase{StudentInfoRepo: studentinfoRepo}

	return &StudentInfoHandler{uc: uc}
}

func (handler *StudentInfoHandler) RegisterStudentInfo(c *gin.Context) {
	req := RegisterStudentInfoReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}

	if req.Name == "" {
		err = errors.New("student name field not null")
		log.Print(err)
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}

	// TODO: 学生氏名と学籍番号のvalidation追加

	id, err := handler.uc.RegisterNewStudentInfo(req.UserId, req.Name, req.StudentNumber)
	if id == -1 {
		err = errors.New("register new student info failed")
		log.Print(err)
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	if err != nil {
		err = errors.New("register new student info failed")
		log.Print(err)
		c.JSON(500, gin.H{"err": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
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

func (handler *StudentInfoHandler) GetStudentInfoByUserId(c *gin.Context) {
	idStr := c.Query("id")
	userId, err := strconv.Atoi(idStr)
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	studentInfo, err := handler.uc.GetStudentInfoByUserId(userId)
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, studentInfo)
}
