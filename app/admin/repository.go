package admin

// Repository represent the admin's repository
type Repository interface {
	GetByID(id string) (data interface{}, err error)
}
