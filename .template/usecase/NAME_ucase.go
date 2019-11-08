package usecase

import (
	"github.com/tamanyan/oauth2-server/{{lowercaseletters .NAME}}"
)

// {{camelcase .NAME}}Usecase usecase
type {{camelcase .NAME}}Usecase struct {
	Repository {{lowercaseletters .NAME}}.Repository
}

// New{{camelcase .NAME}}Usecase will create new an {{camelcase .NAME}}Usecase object representation of {{lowercaseletters .NAME}}.Usecase interface
func New{{camelcase .NAME}}Usecase(Repository {{lowercaseletters .NAME}}.Repository) {{lowercaseletters .NAME}}.Usecase {
	return &{{camelcase .NAME}}Usecase{
		Repository: Repository,
	}
}
