package controller

import (
	"net/http"
	"github.com/labstack/echo"

	"github.com/tamanyan/oauth2-server/app/middleware"
	"github.com/tamanyan/oauth2-server/app/article"
)

// ArticleHandler  represent the httphandler for article
type ArticleHandler struct {
	ArticleUsecase article.Usecase
}

// NewArticleHandler will initialize the /article"/ resources endpoint
func NewArticleHandler(e *echo.Echo, middleware *middleware.GoMiddleware, us article.Usecase) {
	handler := &ArticleHandler{
		ArticleUsecase: us,
	}
	e.GET("/article", handler.GetData)
}

// GetData will get data by ID
func (h *ArticleHandler) GetData(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}