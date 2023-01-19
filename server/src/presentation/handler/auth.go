package handler

import (
	"fmt"
	"net/http"

	"github.com/Doer-org/google-cloud-challenge-2022/usecase"
	"github.com/Doer-org/google-cloud-challenge-2022/utils/helper"
	"github.com/Doer-org/google-cloud-challenge-2022/utils/http/response"
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
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(
				http.StatusBadRequest,
				"StatusBadRequest",
				fmt.Errorf("Get auth url faild : %w", err),
			),
			http.StatusBadRequest,
		)
		return
	}

	url += "&approval_prompt=force&access_type=offline"

	http.Redirect(w, r, url, http.StatusFound)
}

func (h *AuthHandler) Callback(w http.ResponseWriter, r *http.Request) {

	if err := r.FormValue("error"); err != "" {
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(
				http.StatusBadRequest,
				"StatusBadRequest",
				fmt.Errorf("google auth faild : %s", err),
			),
			http.StatusBadRequest,
		)
		return
	}

	state := r.FormValue("state")
	if state == "" {
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(
				http.StatusBadRequest,
				"StatusBadRequest",
				fmt.Errorf("google auth failed: state is empty"),
			),
			http.StatusBadRequest,
		)
		return
	}

	code := r.FormValue("code")
	if code == "" {
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(
				http.StatusBadRequest,
				"StatusBadRequest",
				fmt.Errorf("google auth failed: code is empty"),
			),
			http.StatusBadRequest,
		)
		return
	}

	redirectURL, sessionID, err := h.authUC.Authorization(state, code)
	if err != nil {
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(
				http.StatusBadRequest,
				"StatusBadRequest",
				fmt.Errorf("authorization error : %w", err),
			),
			http.StatusBadRequest,
		)
		return
	}

	if redirectURL == "" {
		response.WriteJsonResponse(
			w,
			response.NewErrResponse(
				http.StatusBadRequest,
				"StatusBadRequest",
				fmt.Errorf("redirect url empty"),
			),
			http.StatusBadRequest,
		)
		return
	}

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
