package admin

import "github.com/tamanyan/oauth2-server/model"

// Repository represent the admin's repository
type Repository interface {
	GetByID(id string) (data interface{}, err error)
	SaveUser(user *model.User) error
}
