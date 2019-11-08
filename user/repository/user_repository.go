package repository

import (
	"github.com/tamanyan/oauth2-server/model"
	"github.com/tamanyan/oauth2-server/user"
)

// UserRepository repository
type UserRepository struct {
}

// NewUserRepository will create new an UserRepository  object representation of user.Repository interface
func NewUserRepository() user.Repository {
	return &UserRepository{}
}

// GetByID will get data by ID
func (r *UserRepository) GetByID(id string) (data interface{}, err error) {
	data = nil
	err = nil
	return
}

// Save will create user
func (r *UserRepository) Save(user *model.User) error {
	return nil
}
