package profile

// Repository represent the profile's repository
type Repository interface {
	GetByID(id string) (data interface{}, err error)
}
