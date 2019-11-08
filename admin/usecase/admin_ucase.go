package usecase

import (
	"github.com/tamanyan/oauth2-server/admin"
	"github.com/tamanyan/oauth2-server/model"
)

// AdminUsecase usecase
type AdminUsecase struct {
	Repository admin.Repository
}

// NewAdminUsecase will create new an AdminUsecase object representation of admin.Usecase interface
func NewAdminUsecase(Repository admin.Repository) admin.Usecase {
	return &AdminUsecase{
		Repository: Repository,
	}
}

// SaveUser will save user
func (us *AdminUsecase) SaveUser(user *model.User) error {
	return nil
}
