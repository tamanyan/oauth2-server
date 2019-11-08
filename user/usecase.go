package user

import "github.com/tamanyan/oauth2-server/model"

// Usecase represent the user's usecases
type Usecase interface {
	Save(user *model.User) error
}
