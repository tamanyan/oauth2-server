package usecase

import (
	"github.com/tamanyan/oauth2-server/model"
	"github.com/tamanyan/oauth2-server/user"
)

// UserUsecase usecase
type UserUsecase struct {
	Repository user.Repository
}

// NewUserUsecase will create new an UserUsecase object representation of user.Usecase interface
func NewUserUsecase(Repository user.Repository) user.Usecase {
	return &UserUsecase{
		Repository: Repository,
	}
}

// Save will create user
func (us *UserUsecase) Save(user *model.User) error {
	return nil
}
