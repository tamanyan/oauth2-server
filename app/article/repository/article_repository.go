package repository

import (
	"github.com/tamanyan/oauth2-server/app/article"
)

// ArticleRepository repository
type ArticleRepository struct {
}

// NewArticleRepository will create new an ArticleRepository  object representation of article.Repository interface
func NewArticleRepository() article.Repository {
	return &ArticleRepository{}
}

// GetByID will get data by ID
func (r *ArticleRepository) GetByID(id string) (data interface{}, err error) {
	data = nil
	err = nil
	return
}