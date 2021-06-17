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

func NewFormatHandler() *FormatHandler {
	uc := usecase.FormatUsecase{}

	return &FormatHandler{uc: uc}
}

func (handler *FormatHandler) ConvertImageToPdf(c *gin.Context) {
	// file, header, err := c.Request.FormFile("image")
	_, header, err := c.Request.FormFile("image")
	if err != nil {
		log.Print(err)
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	filename := header.Filename

	c.JSON(http.StatusOK, gin.H{"filename": filename})
}
