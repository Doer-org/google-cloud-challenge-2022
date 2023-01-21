package google

import (
	"context"
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
	"golang.org/x/oauth2"
)

type Client struct {
	auth  *Google
	cache *cache.Cache
}

func NewClient(redirecturl string) *Client {
	auth := NewGoogle(redirecturl)
	return &Client{auth: auth, cache: cache.New(10*time.Minute, 20*time.Minute)}
}

func (c *Client) Exchange(ctx context.Context, code string) (*oauth2.Token, error) {
	token, err := c.auth.Config.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("Exchange: %w", err)
	}
	return token, nil
}

func (c *Client) GetAuthURL(state string) string {
	return c.auth.Config.AuthCodeURL(state)
}

func (c *Client) Refresh(ctx context.Context, token *oauth2.Token) (*oauth2.Token, error) {
	newtoken, err := c.auth.Config.TokenSource(ctx, token).Token()
	if err != nil {
		return nil, fmt.Errorf("TokenSource: %w", err)
	}
	return newtoken, err
}
