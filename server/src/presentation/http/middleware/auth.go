package middleware

import (
	"fmt"
	"net/http"

	res "github.com/Doer-org/google-cloud-challenge-2022/presentation/http/response"
	"github.com/Doer-org/google-cloud-challenge-2022/usecase"
	mycontext "github.com/Doer-org/google-cloud-challenge-2022/utils/context"
	"golang.org/x/oauth2"
)

type Auth struct {
	uc *usecase.Auth
}

func NewAuth(uc *usecase.Auth) *Auth {
	return &Auth{uc: uc}
}

func (m *Auth) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessCookie, err := r.Cookie("session")
		if err != nil {
			res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: falid to get session: %w", err)), http.StatusBadRequest)
			return
		}
		userId, err := m.uc.GetUserIdFromSession(sessCookie.Value)
		if err != nil {
			res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: GetUserIdFromSession: %w", err)), http.StatusBadRequest)
			return
		}
		token, err := m.uc.GetTokenByUserId(userId)
		if err != nil {
			res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: GetTokenByUserId: %w", err)), http.StatusBadRequest)
			return
		}
		newToken, err := m.uc.RefreshAccessToken(userId, token)
		if err != nil {
			res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: RefreshAccessToken: %w", err)), http.StatusBadRequest)
			return
		}
		token = newToken
		r = setToContext(r, userId.String(), token)
		next.ServeHTTP(w, r)
	})
}

func setToContext(r *http.Request, userID string, token *oauth2.Token) *http.Request {
	ctx := r.Context()
	ctx = mycontext.SetUserId(ctx, userID)
	ctx = mycontext.SetToken(ctx, token)
	return r.WithContext(ctx)
}
