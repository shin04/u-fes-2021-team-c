package handler

import (
	"log"
	"net/http"
	"u-fes-2021-team-c/usecase"

	"github.com/gin-gonic/gin"
)

type FormatHandler struct {
	uc usecase.FormatUsecase
}

type ConvertImageReq struct {
	Image       string
	ConvertType string
}

func NewFormatHandler() *FormatHandler {
	uc := usecase.FormatUsecase{}

	return &FormatHandler{uc: uc}
}

func (handler *FormatHandler) ConvertImageToPdf(c *gin.Context) {
	req := ConvertImageReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}

	// requestのvalidation

	// usecaseで処理

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
