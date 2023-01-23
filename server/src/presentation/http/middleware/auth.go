package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Doer-org/google-cloud-challenge-2022/usecase"
	mycontext "github.com/Doer-org/google-cloud-challenge-2022/utils/context"
	"golang.org/x/oauth2"
)

type Auth struct {
	uc          usecase.IAuth
	frontendURL string
}

func NewAuth(uc usecase.IAuth, url string) *Auth {
	return &Auth{uc: uc, frontendURL: url}
}

func (m *Auth) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessCookie, err := r.Cookie("session")
		if err != nil {
			log.Println(fmt.Errorf("error: falid to get session: %w", err))
			http.Redirect(w, r, m.frontendURL, http.StatusFound)
			return
		}
		userId, err := m.uc.GetUserIdFromSession(sessCookie.Value)
		if err != nil {
			log.Println(fmt.Errorf("error: GetUserIdFromSession: %w", err))
			http.Redirect(w, r, m.frontendURL, http.StatusFound)
			return
		}
		token, err := m.uc.GetTokenByUserId(userId)
		if err != nil {
			log.Println(fmt.Errorf("error: GetTokenByUserId: %w", err))
			http.Redirect(w, r, m.frontendURL, http.StatusFound)
			return
		}
		newToken, err := m.uc.RefreshAccessToken(userId, token)
		if err != nil {
			log.Println(fmt.Errorf("error: RefreshAccessToken: %w", err))
			http.Redirect(w, r, m.frontendURL, http.StatusFound)
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
