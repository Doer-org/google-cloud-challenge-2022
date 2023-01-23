package handler

import (
	"fmt"
	"net/http"

	res "github.com/Doer-org/google-cloud-challenge-2022/presentation/http/response"
	"github.com/Doer-org/google-cloud-challenge-2022/usecase"
	"github.com/Doer-org/google-cloud-challenge-2022/utils/env"
)

const oneWeek = 60 * 60 * 24 * 7

type Auth struct {
	authUC usecase.IAuth
}

func NewAuth(authUC usecase.IAuth) *Auth {
	return &Auth{authUC: authUC}
}

func (h *Auth) Login(w http.ResponseWriter, r *http.Request) {
	redirectURL := r.FormValue("redirect_url")
	url, state, err := h.authUC.GetAuthURL(redirectURL)
	if err != nil {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: GetAuthURL: %w", err)), http.StatusBadRequest)
		return
	}
	url += "&approval_prompt=force&access_type=offline"

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    state,
		Path:     "/",
		Secure:   !env.IsLocal(),
		HttpOnly: true,
	})

	http.Redirect(w, r, url, http.StatusFound)
}

func (h *Auth) Callback(w http.ResponseWriter, r *http.Request) {
	if errFormValue := r.FormValue("error"); errFormValue != "" {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: error is empty")), http.StatusBadRequest)
		return
	}
	state := r.FormValue("state")
	if state == "" {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: state is empty")), http.StatusBadRequest)
		return
	}

	sessCookie, err := r.Cookie("session")

	if sessCookie.Value != state {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: state is not correct")), http.StatusBadRequest)
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
	sameSite := http.SameSiteNoneMode
	if env.IsLocal() {
		sameSite = http.SameSiteLaxMode
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    sessionID,
		Path:     "/",
		MaxAge:   oneWeek,
		Secure:   !env.IsLocal(),
		HttpOnly: true,
		SameSite: sameSite,
	})
	http.Redirect(w, r, redirectURL, http.StatusFound)
}

func (h *Auth) Validate(w http.ResponseWriter, r *http.Request) {
	res.New200SuccessJson("validate success")
}
