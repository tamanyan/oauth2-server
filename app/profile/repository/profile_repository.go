package repository

import (
	"github.com/tamanyan/oauth2-server/app/profile"
)

// ProfileRepository repository
type ProfileRepository struct {
}

// NewProfileRepository will create new an ProfileRepository  object representation of profile.Repository interface
func NewProfileRepository() profile.Repository {
	return &ProfileRepository{}
}

// GetByID will get data by ID
func (r *ProfileRepository) GetByID(id string) (data interface{}, err error) {
	data = nil
	err = nil
	return
}