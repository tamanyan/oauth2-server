package response

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/tamanyan/oauth2-server/common"
	"github.com/tamanyan/oauth2-server/admin"
)

// AdminResponse usecase
type AdminResponse struct {
	StatucCode int
	Data       echo.Map
}

// NewAdminResponse will create a new AdminResponse
func NewAdminResponse(data: interface{}) common.Response {
	return &AdminResponse{
		StatucCode: http.StatusOK,
		Data:       echo.Map{},
	}
}

// GetStatusCode Get HTTP status code
func (r *AdminResponse) GetStatusCode() int {
	return r.StatucCode
}

// GetData Get ehco Map data
func (r *AdminResponse) GetData() echo.Map {
	return r.Data
}
