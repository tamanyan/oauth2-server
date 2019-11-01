package usecase

import (
	"github.com/tamanyan/oauth2-server/app/profile"
)

// ProfileUsecase usecase
type ProfileUsecase struct {
	Repository profile.Repository
}

// NewProfileUsecase will create new an ProfileUsecase object representation of profile.Usecase interface
func NewProfileUsecase(Repository profile.Repository) profile.Usecase {
	return &ProfileUsecase{
		Repository: Repository,
	}
}
