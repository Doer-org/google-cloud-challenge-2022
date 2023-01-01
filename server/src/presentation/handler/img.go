package handler

import (
	"net/http"

	"github.com/Doer-org/google-cloud-challenge-2022/usecase"
)

type ImageHandler struct {
	UC usecase.IImageUsecase
}

func NewImgHandler(uc usecase.IImageUsecase) *ImageHandler {
	return &ImageHandler{
		UC: uc,
	}
}

// Debug用
func (h *ImageHandler) Create(w http.ResponseWriter, r *http.Request) {
	
}
