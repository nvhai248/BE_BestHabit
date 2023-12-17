package oauthprovider

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GGOAuthProvider interface {
	GetGoogleOauthConfig() *oauth2.Config
	GetOauthStateString() string
}

type GGOAuthConfig struct {
	ClientId     string
	ClientSecret string
	RedirectURL  string
	Scopes       []string
}

func NewGGOAuthProvider(clientId, clientSecret, redirectUrl string, scopes []string) *GGOAuthConfig {
	return &GGOAuthConfig{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		RedirectURL:  redirectUrl,
		Scopes:       scopes,
	}
}

func (cf *GGOAuthConfig) GetGoogleOauthConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     cf.ClientId,
		ClientSecret: cf.ClientSecret,
		RedirectURL:  cf.RedirectURL,
		Scopes:       cf.Scopes,
		Endpoint:     google.Endpoint,
	}
}

func (cf *GGOAuthConfig) GetOauthStateString() string {
	return "random"
}
