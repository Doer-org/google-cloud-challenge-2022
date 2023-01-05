package google

import (
	"context"

	"golang.org/x/oauth2"
)

type Auth interface {
	GetAuthURL(state string) string
	Exchange(ctx context.Context, code string) (*oauth2.Token, error)
	Refresh(ctx context.Context, token *oauth2.Token) (*oauth2.Token, error)
}
