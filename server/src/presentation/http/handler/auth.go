package handler

import (
	"fmt"
	"net/http"

	"github.com/Doer-org/google-cloud-challenge-2022/config"
	res "github.com/Doer-org/google-cloud-challenge-2022/presentation/http/response"
	"github.com/Doer-org/google-cloud-challenge-2022/usecase"
)

const oneWeek = 60 * 60 * 24 * 7

type IAuth interface {
	Login(w http.ResponseWriter, r *http.Request)
	Callback(w http.ResponseWriter, r *http.Request)
	Validate(w http.ResponseWriter, r *http.Request)
	User(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

// TODO:logout apiも作る必要あり
type Auth struct {
	authUC usecase.IAuth
	userUC usecase.IUser
}

func NewAuth(auc usecase.IAuth, uuc usecase.IUser) IAuth {
	return &Auth{
		authUC: auc,
		userUC: uuc,
	}
}

func (h *Auth) Login(w http.ResponseWriter, r *http.Request) {
	redirectURL := r.FormValue("redirect_url")
	url, state, err := h.authUC.GetAuthURL(r.Context(), redirectURL)
	if err != nil {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: GetAuthURL: %w", err)), http.StatusBadRequest)
		return
	}
	url += "&approval_prompt=force&access_type=offline"

	sameSite := http.SameSiteNoneMode
	if config.IsDev() {
		sameSite = http.SameSiteLaxMode
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    state,
		Path:     "/",
		Secure:   !config.IsDev(),
		HttpOnly: true,
		SameSite: sameSite,
		Domain:   config.CLIENT_DOMAIN,
	})
	http.Redirect(w, r, url, http.StatusFound)
}

// logout には session の削除か、cokkieの削除とかで出来て、今回はどっちもやってます。
// maxage をマイナスにすることで、cokkieを消せます。　https://tech-up.hatenablog.com/entry/2019/01/03/121435
func (h *Auth) Logout(w http.ResponseWriter, r *http.Request) {
	sessCookie, err := r.Cookie("session")
	if err != nil {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: falid to get session: %w", err)), http.StatusBadRequest)
		return
	}
	err = h.authUC.DeleteSession(r.Context(), sessCookie.Value)
	if err != nil {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: falid to delete session: %w", err)), http.StatusBadRequest)
		return
	}

	sessCookie.MaxAge = -1
	http.SetCookie(w, sessCookie)
	res.WriteJson(w, res.New200SuccessJson("logout success"), http.StatusOK)
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
	redirectURL, sessionID, err := h.authUC.Authorization(r.Context(), state, code)
	if err != nil {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: Authorization: %w", err)), http.StatusBadRequest)
		return
	}
	if redirectURL == "" {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: redirect url is empty")), http.StatusBadRequest)
		return
	}
	sameSite := http.SameSiteNoneMode
	if config.IsDev() {
		sameSite = http.SameSiteLaxMode
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    sessionID,
		Path:     "/",
		MaxAge:   oneWeek,
		Secure:   !config.IsDev(),
		HttpOnly: true,
		SameSite: sameSite,
		Domain:   config.CLIENT_DOMAIN,
	})
	http.Redirect(w, r, redirectURL, http.StatusFound)
}

func (h *Auth) Validate(w http.ResponseWriter, r *http.Request) {
	res.WriteJson(w, res.New200SuccessJson("validate success"), http.StatusOK)
}

func (h *Auth) User(w http.ResponseWriter, r *http.Request) {
	sessCookie, err := r.Cookie("session")
	if err != nil {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: Cookie: %w", err)), http.StatusBadRequest)
		return
	}
	userId, err := h.authUC.GetUserIdFromSession(r.Context(), sessCookie.Value)
	if err != nil {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: GetUserIdFromSession: %w", err)), http.StatusBadRequest)
		return
	}
	user, err := h.userUC.GetUserById(r.Context(), userId.String())
	if err != nil {
		res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: GetUserById: %w", err)), http.StatusBadRequest)
		return
	}
	res.WriteJson(w, user, http.StatusOK)
}
