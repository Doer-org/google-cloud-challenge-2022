package handler

import (
	"net/http"

	"github.com/Doer-org/google-cloud-challenge-2022/usecase"
)

type AuthHandler struct {
	authUC      *usecase.AuthUsecase
	frontendURL string
}

func NewAuthHandler(authUC *usecase.AuthUsecase, frontendURL string) *AuthHandler {
	return &AuthHandler{authUC: authUC, frontendURL: frontendURL}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {

	redirectURL := r.FormValue("redirect_url")

	url, err := h.authUC.GetAuthURL(redirectURL)
	if err != nil {
		
	}

	http.Redirect(w, r, url, http.StatusFound)
}
