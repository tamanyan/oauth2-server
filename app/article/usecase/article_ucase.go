package usecase

import (
	"github.com/tamanyan/oauth2-server/app/article"
)

// ArticleUsecase usecase
type ArticleUsecase struct {
	Repository article.Repository
}

// NewArticleUsecase will create new an ArticleUsecase object representation of article.Usecase interface
func NewArticleUsecase(Repository article.Repository) article.Usecase {
	return &ArticleUsecase{
		Repository: Repository,
	}
}
