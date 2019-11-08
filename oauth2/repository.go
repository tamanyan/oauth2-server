package oauth2

import (
	"context"
)

// Repository represent the oauth2's repository
type Repository interface {
	Store(context.Context) error
}
