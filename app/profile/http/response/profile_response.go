package response

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/tamanyan/oauth2-server/app/common"
	"github.com/tamanyan/oauth2-server/profile"
)

// ProfileResponse usecase
type ProfileResponse struct {
	StatucCode int
	Data       echo.Map
}

// NewProfileResponse will create a new ProfileResponse
func NewProfileResponse(data: interface{}) common.Response {
	return &ProfileResponse{
		StatucCode: http.StatusOK,
		Data:       echo.Map{},
	}
}

// GetStatusCode Get HTTP status code
func (r *ProfileResponse) GetStatusCode() int {
	return r.StatucCode
}

// GetData Get ehco Map data
func (r *ProfileResponse) GetData() echo.Map {
	return r.Data
}
