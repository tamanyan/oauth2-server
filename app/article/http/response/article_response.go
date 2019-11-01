package response

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/tamanyan/oauth2-server/app/common"
	"github.com/tamanyan/oauth2-server/article"
)

// ArticleResponse usecase
type ArticleResponse struct {
	StatucCode int
	Data       echo.Map
}

// NewArticleResponse will create a new ArticleResponse
func NewArticleResponse(data: interface{}) common.Response {
	return &ArticleResponse{
		StatucCode: http.StatusOK,
		Data:       echo.Map{},
	}
}

// GetStatusCode Get HTTP status code
func (r *ArticleResponse) GetStatusCode() int {
	return r.StatucCode
}

// GetData Get ehco Map data
func (r *ArticleResponse) GetData() echo.Map {
	return r.Data
}
