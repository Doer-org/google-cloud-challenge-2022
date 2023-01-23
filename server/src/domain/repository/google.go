package repository

import (
	"context"

	"github.com/Doer-org/google-cloud-challenge-2022/infrastructure/ent"
	"golang.org/x/oauth2"
)

type IGoogle interface {
	GetAuthURL(state string) string
	Exchange(ctx context.Context, code string) (*oauth2.Token, error)
	Refresh(ctx context.Context, token *oauth2.Token) (*oauth2.Token, error)
	GetMe(ctx context.Context) (*ent.User, error)
}
