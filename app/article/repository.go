package article

// Repository represent the article's repository
type Repository interface {
	GetByID(id string) (data interface{}, err error)
}
