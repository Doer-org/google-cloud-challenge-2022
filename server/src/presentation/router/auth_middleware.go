package router

import (
	"fmt"
	"net/http"

	"github.com/Doer-org/google-cloud-challenge-2022/usecase"
	"github.com/Doer-org/google-cloud-challenge-2022/utils"
	"github.com/Doer-org/google-cloud-challenge-2022/utils/http/response"
	"golang.org/x/oauth2"
)

type AuthMiddleware struct {
	uc *usecase.AuthUsecase
}

// NewAuthMiddleware web.AuthMiddlewareのポインタを生成します。
func NewAuthMiddleware(uc *usecase.AuthUsecase) *AuthMiddleware {
	return &AuthMiddleware{uc: uc}
}

func (m *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessCookie, err := r.Cookie("session")
		if err != nil {
			response.NewErrResponse(w, fmt.Errorf("falid to get session err : %w", err))
			return
		}
		userID, err := m.uc.GetUserIDFromSession(sessCookie.Value)
		if err != nil {
			response.NewErrResponse(w, fmt.Errorf("faild to get userId from sessinId: %s , error :  %w", sessCookie.Value, err))
			return
		}
		token, err := m.uc.GetTokenByUserID(userID)
		if err != nil {
			response.NewErrResponse(w, fmt.Errorf("falid to get token err : %w", err))
			return
		}

		newToken, err := m.uc.RefreshAccessToken(userID, token)
		if err != nil {
			response.NewErrResponse(w, fmt.Errorf("falid to get token err : %w", err))
			return
		}

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
