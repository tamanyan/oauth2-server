package user

import "github.com/tamanyan/oauth2-server/model"

// Repository represent the user's repository
type Repository interface {
	GetByID(id string) (data interface{}, err error)
	Save(user *model.User) error
}
