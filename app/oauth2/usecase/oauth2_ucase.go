package usecase

import (
	"context"
	"log"
	"time"

	_oauth2 "github.com/tamanyan/oauth2-server/app/oauth2"
	"github.com/tamanyan/oauth2-server/app/oauth2/http/request"
	"github.com/tamanyan/oauth2"
	"github.com/tamanyan/oauth2/errors"
)

// OAuth2Usecase usecase
type OAuth2Usecase struct {
	Manager oauth2.Manager
}

// NewOAuth2Usecase will create new an articleUsecase object representation of article.Usecase interface
func NewOAuth2Usecase(Manager oauth2.Manager, mamatimeout time.Duration) _oauth2.Usecase {
	return &OAuth2Usecase{
		Manager: Manager,
	}
}

// IssuePasswordCredentialAccessToken Issue password credential access token
func (a *OAuth2Usecase) IssuePasswordCredentialAccessToken(ctx context.Context, request request.OAuth2PasswordCredentialRequest) (ti oauth2.TokenInfo, err error) {
	// Check Client ID and Secret
	cli, err := a.Manager.GetClient(request.ClientID)

	if err != nil || cli.GetSecret() != request.ClientSecret {
		err = errors.ErrInvalidClient
		return
	}

	// Check username and password
	if !(request.Username == "test" && request.Password == "test") {
		err = errors.ErrInvalidGrant
		return
	}

	tgr := &oauth2.TokenGenerateRequest{
		ClientID:     request.ClientID,
		ClientSecret: request.ClientSecret,
		Scope:        request.Scope,
		UserID:       request.Username,
	}

	ti, err = a.Manager.GenerateAccessToken(oauth2.GrantType(request.GrantType), tgr)

	return
}

// IssueRefreshAccessToken Issue refresh token
func (a *OAuth2Usecase) IssueRefreshAccessToken(ctx context.Context, request request.OAuth2RefreshTokenRequest) (ti oauth2.TokenInfo, err error) {
	// Check Client ID and Secret
	cli, err := a.Manager.GetClient(request.ClientID)

	if err != nil || cli.GetSecret() != request.ClientSecret {
		err = errors.ErrInvalidClient
		return
	}

	tgr := &oauth2.TokenGenerateRequest{
		ClientID:     request.ClientID,
		ClientSecret: request.ClientSecret,
		Scope:        request.Scope,
		Refresh:      request.RefreshToken,
	}

	ti, err = a.Manager.RefreshAccessToken(tgr)

	log.Println(ti)

	return
}

// IssueClientCredentialAccessToken will issue client credential access token
func (a *OAuth2Usecase) IssueClientCredentialAccessToken(ctx context.Context, request request.OAuth2ClientCredentialRequest) (ti oauth2.TokenInfo, err error) {
	// Check Client ID and Secret
	cli, err := a.Manager.GetClient(request.ClientID)

	if err != nil || cli.GetSecret() != request.ClientSecret {
		err = errors.ErrInvalidClient
		return
	}

	tgr := &oauth2.TokenGenerateRequest{
		ClientID:     request.ClientID,
		ClientSecret: request.ClientSecret,
		Scope:        request.Scope,
	}

	ti, err = a.Manager.GenerateAccessToken(oauth2.GrantType(request.GrantType), tgr)

	return
}

// RevokeAccessToken will revoke access token
func (a *OAuth2Usecase) RevokeAccessToken(ctx context.Context, token string) error {
	return a.Manager.RemoveAccessToken(token)
}

// VerifyAccessToken will verify access token
func (a *OAuth2Usecase) VerifyAccessToken(ctx context.Context, token string) (ti oauth2.TokenInfo, err error) {
	return a.Manager.LoadAccessToken(token)
}
