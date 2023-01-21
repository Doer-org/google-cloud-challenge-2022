package middleware

import (
	"fmt"
	"net/http"

	res "github.com/Doer-org/google-cloud-challenge-2022/presentation/http/response"
	"github.com/Doer-org/google-cloud-challenge-2022/usecase"
	"github.com/Doer-org/google-cloud-challenge-2022/utils"
	"github.com/go-chi/chi/v5"
	"golang.org/x/oauth2"
)

type AuthMiddleware struct {
	uc *usecase.Auth
}

func NewAuthMiddleware(uc *usecase.Auth) *AuthMiddleware {
	return &AuthMiddleware{uc: uc}
}

func (m *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessCookie, err := r.Cookie("session")
		if err != nil {
			res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: falid to get session: %w", err)), http.StatusBadRequest)
			return
		}
		userID, err := m.uc.GetUserIDFromSession(sessCookie.Value)
		if err != nil {
			res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: GetUserIDFromSession: %w", err)), http.StatusBadRequest)
			return
		}
		token, err := m.uc.GetTokenByUserID(userID)
		if err != nil {
			res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: GetTokenByUserID: %w", err)), http.StatusBadRequest)
			return
		}
		newToken, err := m.uc.RefreshAccessToken(userID, token)
		if err != nil {
			res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: RefreshAccessToken: %w", err)), http.StatusBadRequest)
			return
		}
		// TODO: なぜ
		token = newToken
		r = setToContext(r, userID, token)
		next.ServeHTTP(w, r)
	})
}

func setToContext(r *http.Request, userID string, token *oauth2.Token) *http.Request {
	ctx := r.Context()
	ctx = utils.SetUserIDToContext(ctx, userID)
	ctx = utils.SetTokenToContext(ctx, token)
	r = r.WithContext(ctx)
	return r
}

// TODO: 使われていない
func setAuthMiddleware(r *chi.Mux, uc *usecase.Auth) {
	authmiddle := NewAuthMiddleware(uc)
	r.Use(authmiddle.Authenticate)
}
