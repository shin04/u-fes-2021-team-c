package handler

import "u-fes-2021-team-c/usecase"

type FormatHandler struct {
	uc usecase.FormatUsecase
}

func NewFormatHandler() *FormatHandler {
	uc := usecase.FormatUsecase{}

	return &FormatHandler{uc: uc}
}

func (handler *FormatHandler) ConvertImageToPdf() {

}
