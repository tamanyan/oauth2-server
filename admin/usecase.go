package admin

import (
	"github.com/tamanyan/oauth2-server/model"
)

// Usecase represent the admin's usecases
type Usecase interface {
	SaveUser(user *model.User) error
}
