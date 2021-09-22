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

type UserHandler struct {
	uc usecase.UserUsecase
}

type RegisteruserReq struct {
	Name     string
	Password string
}

type LoginReq struct {
	Name     string
	Password string
}

func NewUserHandler(userRepo repository.UserRepository) *UserHandler {
	uc := usecase.UserUsecase{UserRepo: userRepo}

	return &UserHandler{uc: uc}
}

func (handler *UserHandler) RegisterUser(c *gin.Context) {
	userReq := RegisteruserReq{}
	err := c.ShouldBindJSON(&userReq)
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}

	if userReq.Name == "" || userReq.Password == "" {
		err = errors.New("username or password field not null")
		log.Print(err)
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}

	userId, err := handler.uc.RegisterNewUser(userReq.Name, userReq.Password)
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	if userId == -1 {
		err = errors.New("register usesr failed")
		log.Print(err)
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"userId": userId})
}

func (handler *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := handler.uc.GetAllUsers()
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	if len(users) < 1 {
		err = errors.New("users not found")
		log.Print(err)
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (handler *UserHandler) GetUser(c *gin.Context) {
	idStr := c.Query("id")
	userId, err := strconv.Atoi(idStr)
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	user, err := handler.uc.GetUserById(userId)
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (handler *UserHandler) Login(c *gin.Context) {
	req := LoginReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}

	if req.Name == "" || req.Password == "" {
		err = errors.New("username or password field not null")
		log.Print(err)
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}

	user, err := handler.uc.Login(req.Name, req.Password)
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"userId": user.Id})
}
