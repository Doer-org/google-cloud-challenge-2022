package google

import (
	"os"

	"golang.org/x/oauth2"
	googleOAuth "golang.org/x/oauth2/google"
)

type Google struct {
	Config *oauth2.Config
}

func NewGoogle(redirecturl string) *Google {
	return newGoogle(redirecturl)
}

func newGoogle(redirecturl string) *Google {
	google := &Google{
		Config: &oauth2.Config{
			ClientID:     os.Getenv("Google_ID"),
			ClientSecret: os.Getenv("Google_SECRET"),
			Endpoint:     googleOAuth.Endpoint,
			Scopes:       []string{"openid", "email", "profile"},
			RedirectURL:  redirecturl,
		},
	}
	return google
}
