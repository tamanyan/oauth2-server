package repository

import (
	"github.com/tamanyan/oauth2-server/app/admin"
)

// AdminRepository repository
type AdminRepository struct {
}

// NewAdminRepository will create new an AdminRepository  object representation of admin.Repository interface
func NewAdminRepository() admin.Repository {
	return &AdminRepository{}
}

// GetByID will get data by ID
func (r *AdminRepository) GetByID(id string) (data interface{}, err error) {
	data = nil
	err = nil
	return
}