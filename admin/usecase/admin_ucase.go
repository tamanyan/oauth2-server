package usecase

import (
	"github.com/tamanyan/oauth2-server/admin"
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
