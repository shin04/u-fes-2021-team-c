package handler

import (
	"net/http"
	"strconv"
	"u-fes-2021-team-c/database"
	"u-fes-2021-team-c/model"
	"u-fes-2021-team-c/usecase"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	uc usecase.UserUsecase
}

type getUsesrReq struct {
	userId int `json:"userId"`
}

func NewUserHandler(sqlHandler database.SqlHandler) *UserHandler {
	uc := usecase.UserUsecase{
		UserRepo: database.UserRepository{
			SqlHandler: sqlHandler,
		},
	}

	return &UserHandler{uc: uc}
}

func (handler *UserHandler) RegisterUser(c *gin.Context) {
	user := model.User{}
	err := c.Bind(&user)
	if err != nil {
		c.JSON(500, err)
	}

	userId, err := handler.uc.RegisterNewUser(&user)
	if err != nil {
		c.JSON(500, err)
		return
	}
	if userId == -1 {
		c.JSON(500, "regisster usesr failed")
	}

	c.JSON(http.StatusOK, gin.H{"userId": userId})
}

func (handler *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := handler.uc.GetAllUsers()
	if err != nil {
		c.JSON(500, err)
	}
	if len(users) < 1 {
		c.JSON(500, "users not found")
	}

	c.JSON(http.StatusOK, users)
}

func (handler *UserHandler) GetUser(c *gin.Context) {
	idStr := c.Param("id")
	userId, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(500, "invalid param")
	}
	user, err := handler.uc.GetUserById(userId)
	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(http.StatusOK, user)
}
