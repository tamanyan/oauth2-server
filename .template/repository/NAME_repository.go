package repository

import (
	"github.com/tamanyan/oauth2-server/{{lowercaseletters .NAME}}"
)

// {{camelcase .NAME}}Repository repository
type {{camelcase .NAME}}Repository struct {
}

// New{{camelcase .NAME}}Repository will create new an {{camelcase .NAME}}Repository  object representation of {{lowercaseletters .NAME}}.Repository interface
func New{{camelcase .NAME}}Repository() {{lowercaseletters .NAME}}.Repository {
	return &{{camelcase .NAME}}Repository{}
}

// GetByID will get data by ID
func (r *{{camelcase .NAME}}Repository) GetByID(id string) (data interface{}, err error) {
	data = nil
	err = nil
	return
}