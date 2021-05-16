package handler

import (
	"net/http"
	"u-fes-2021-team-c/database"
	"u-fes-2021-team-c/model"
	"u-fes-2021-team-c/usecase"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	uc usecase.UserUsecase
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

	userId, err := handler.uc.UserRepo.CreateUser(&user)
	if err != nil {
		c.JSON(500, err)
		return
	}
	if userId == -1 {
		c.JSON(500, "regisster usesr failed")
	}

	c.JSON(http.StatusOK, gin.H{"userId": userId})
}
