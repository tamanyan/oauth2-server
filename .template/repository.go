package {{lowercaseletters .NAME}}

// Repository represent the {{lowercaseletters .NAME}}'s repository
type Repository interface {
	GetByID(id string) (data interface{}, err error)
}
