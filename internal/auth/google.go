package auth

import (
	"context"

	"github.com/coreos/go-oidc"
	"github.com/pkg/errors"
	"github.com/tukejonny/go-oauth2-example/internal/config"
	"golang.org/x/oauth2"
)

const (
	googleProvider = "https://accounts.google.com"
)

// GoogleAuth ... Google OIDC Authentication Client
type GoogleAuth struct {
	provider *oidc.Provider
	cfg      *oauth2.Config
}

func NewGoogleAuth(ctx context.Context, cfg *config.AuthConfig) (*GoogleAuth, error) {
	provider, err := oidc.NewProvider(ctx, googleProvider)
	if err != nil {
		return nil, err
	}

	scopes := make([]string, len(cfg.Scopes)+1)
	scopes[0] = oidc.ScopeOpenID
	for i := 0; i < len(cfg.Scopes); i++ {
		scopes[i+1] = cfg.Scopes[i]
	}

	return &GoogleAuth{
		provider: provider,
		cfg: &oauth2.Config{
			ClientID:     cfg.ClientID,
			ClientSecret: cfg.ClientSecret,

			Scopes: cfg.Scopes,

			Endpoint:    provider.Endpoint(),
			RedirectURL: cfg.RedirectURL,
		},
	}, nil
}

func (a *GoogleAuth) Config() *oauth2.Config {
	return a.cfg
}

// FetchUserInfo ... Fetch user's information with oidc authentication
func (a *GoogleAuth) FetchUserInfo(code string) (*oidc.UserInfo, error) {
	ctx := context.Background()
	oauth2Token, err := a.cfg.Exchange(ctx, code)
	if err != nil {
		return nil, errors.Wrap(err, "トークンの取得に失敗しました")
	}

	userInfo, err := a.provider.UserInfo(ctx, oauth2.StaticTokenSource(oauth2Token))
	if err != nil {
		return nil, errors.Wrap(err, "ユーザ情報の取得に失敗しました")
	}

	return userInfo, nil
}
