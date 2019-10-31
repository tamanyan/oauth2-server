package oauth2

import (
	"context"
)

// Usecase represent the article's usecases
type Usecase interface {
	IssueAccessToken(ctx context.Context) (string, error)
}
