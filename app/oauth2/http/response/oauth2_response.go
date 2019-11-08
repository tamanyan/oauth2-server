package response

import (
	"net/http"
	"time"

	"github.com/labstack/echo"

	"github.com/tamanyan/oauth2"
	"github.com/tamanyan/oauth2-server/app/common"
)

// OAuth2Response usecase
type OAuth2Response struct {
	StatucCode int
	Data       echo.Map
}

// NewOAuth2Response will create a new OAuth2Response
func NewOAuth2Response(ti oauth2.TokenInfo) common.Response {
	data := echo.Map{
		"access_token": ti.GetAccess(),
		"token_type":   "Bearer",
		"expires_in":   int64(ti.GetAccessExpiresIn() / time.Second),
	}

	if scope := ti.GetScope(); scope != "" {
		data["scope"] = scope
	}

	if refresh := ti.GetRefresh(); refresh != "" {
		data["refresh_token"] = refresh
	}

	return &OAuth2Response{
		StatucCode: http.StatusOK,
		Data:       data,
	}
}

// GetStatusCode Get HTTP status code
func (r *OAuth2Response) GetStatusCode() int {
	return r.StatucCode
}

// GetData Get ehco Map data
func (r *OAuth2Response) GetData() echo.Map {
	return r.Data
}
