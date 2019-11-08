package oauth2

import (
	"context"

	"github.com/tamanyan/oauth2-server/oauth2/http/request"
	"github.com/tamanyan/oauth2"
)

// Usecase represent the article's usecases
type Usecase interface {
	IssuePasswordCredentialAccessToken(ctx context.Context, request request.OAuth2PasswordCredentialRequest) (ti oauth2.TokenInfo, err error)
	IssueRefreshAccessToken(ctx context.Context, request request.OAuth2RefreshTokenRequest) (ti oauth2.TokenInfo, err error)
	IssueClientCredentialAccessToken(ctx context.Context, request request.OAuth2ClientCredentialRequest) (ti oauth2.TokenInfo, err error)
	RevokeAccessToken(ctx context.Context, token string) error
	VerifyAccessToken(ctx context.Context, token string) (ti oauth2.TokenInfo, err error)
}
