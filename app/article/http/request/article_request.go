package request

// ArticleRequest GET /article validation
type ArticleRequest struct {
	Sample    string `form:"sample" validate:"required"`
}
