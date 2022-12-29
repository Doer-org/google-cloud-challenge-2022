package handler

import (
	"context"
	"net/http"

	"github.com/Doer-org/google-cloud-challenge-2022/usecase"
)


type UserHandler struct {
	UC usecase.IUserUsecase
}

func NewUserHandler(uc usecase.IUserUsecase) *UserHandler {
	return &UserHandler{
		UC: uc,
	}
}

func (h *UserHandler) Create(w http.ResponseWriter,r *http.Request){
	err := h.UC.Create(context.Background(),1,"hoge",false,"","")
	if err!=nil{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	w.Write([]byte("ok"))
}
