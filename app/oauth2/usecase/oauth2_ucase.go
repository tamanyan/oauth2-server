package usecase

import (
	"context"
	errs "errors"
	"time"

	"github.com/tamanyan/oauth2-server/app/oauth2"
)

// OAuth2Usecase usecase
type OAuth2Usecase struct {
}

// NewOAuth2Usecase will create new an articleUsecase object representation of article.Usecase interface
func NewOAuth2Usecase(timeout time.Duration) oauth2.Usecase {
	return &OAuth2Usecase{}
}

// IssueAccessToken Issue access token
func (a *OAuth2Usecase) IssueAccessToken(ctx context.Context) (string, error) {

	return "", errs.New("unknown")
}
