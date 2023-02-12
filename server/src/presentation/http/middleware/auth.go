package middleware

import (
	"fmt"
	"net/http"

	res "github.com/Doer-org/google-cloud-challenge-2022/presentation/http/response"
	"github.com/Doer-org/google-cloud-challenge-2022/usecase"
	mycontext "github.com/Doer-org/google-cloud-challenge-2022/utils/context"
)

type IAuth interface {
	Authenticate(next http.Handler) http.Handler
}

type Auth struct {
	uc usecase.IAuth
}

func NewAuth(uc usecase.IAuth) IAuth {
	return &Auth{uc: uc}
}

func (m *Auth) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessCookie, err := r.Cookie("session")
		if err != nil {
			res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: falid to get session: %w", err)), http.StatusBadRequest)
			return
		}
		ok, err := m.uc.CheckSessionExpiry(r.Context(), sessCookie.Value)
		if err != nil {
			res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: check session expiry error: %w", err)), http.StatusBadRequest)
			return
		}
		if !ok {
			res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: session expiry time over: %w", err)), http.StatusBadRequest)
			return
		}
		userId, err := m.uc.GetUserIdFromSession(r.Context(), sessCookie.Value)
		if err != nil {
			res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: GetUserIdFromSession: %w", err)), http.StatusBadRequest)
			return
		}
		token, err := m.uc.GetTokenByUserId(r.Context(), userId)
		if err != nil {
			res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: GetTokenByUserId: %w", err)), http.StatusBadRequest)
			return
		}
		// TODO 切れてたらrefresh
		newToken, err := m.uc.RefreshAccessToken(r.Context(), userId, token)
		if err != nil {
			res.WriteJson(w, res.New404ErrJson(fmt.Errorf("error: RefreshAccessToken: %w", err)), http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		ctx = mycontext.SetUserId(ctx, userId.String())
		ctx = mycontext.SetToken(ctx, newToken)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
