package handler

import (
	"fmt"
	"net/http"

	res "github.com/Doer-org/google-cloud-challenge-2022/presentation/http/response"
	"github.com/Doer-org/google-cloud-challenge-2022/usecase"
	helper "github.com/Doer-org/google-cloud-challenge-2022/utils/env"
)

const sevenDays = 60 * 60 * 24 * 7

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
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: GetAuthURL: %w", err)), http.StatusBadRequest)
		return
	}
	url += "&approval_prompt=force&access_type=offline"
	http.Redirect(w, r, url, http.StatusFound)
}

func (h *AuthHandler) Callback(w http.ResponseWriter, r *http.Request) {
	// TODO: usecase?
	if errFormValue := r.FormValue("error"); errFormValue != "" {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: error is empty")), http.StatusBadRequest)
		return
	}
	state := r.FormValue("state")
	if state == "" {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: state is empty")), http.StatusBadRequest)
		return
	}
	code := r.FormValue("code")
	if code == "" {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: code is empty")), http.StatusBadRequest)
		return
	}
	redirectURL, sessionID, err := h.authUC.Authorization(state, code)
	if err != nil {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: Authorization: %w", err)), http.StatusBadRequest)
		return
	}
	if redirectURL == "" {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: redirect url is empty")), http.StatusBadRequest)
		return
	}
	// TODO: これなに
	sameSite := http.SameSiteNoneMode
	if helper.IsLocal() {
		sameSite = http.SameSiteLaxMode
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    sessionID,
		Path:     "/",
		MaxAge:   sevenDays,
		Secure:   !helper.IsLocal(),
		HttpOnly: true,
		SameSite: sameSite,
	})
	http.Redirect(w, r, redirectURL, http.StatusFound)
}
