package response

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/tamanyan/oauth2-server/common"
	"github.com/tamanyan/oauth2-server/user"
)

// UserResponse usecase
type UserResponse struct {
	StatucCode int
	Data       echo.Map
}

// NewUserResponse will create a new UserResponse
func NewUserResponse(data interface{}) common.Response {
	return &UserResponse{
		StatucCode: http.StatusOK,
		Data:       echo.Map{},
	}
}

// GetStatusCode Get HTTP status code
func (r *UserResponse) GetStatusCode() int {
	return r.StatucCode
}

// GetData Get ehco Map data
func (r *UserResponse) GetData() echo.Map {
	return r.Data
}
