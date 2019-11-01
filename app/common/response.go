package common

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/tamanyan/oauth2-server/errors"
)

// Response represent the oauth2's response
type Response interface {
	GetStatusCode() int
	GetData() echo.Map
}

// ErrorResponse usecase
type ErrorResponse struct {
	StatucCode int
	Data       echo.Map
}

// NewErrorResponse will create a new NewErrorResponse
func NewErrorResponse(err error) Response {
	if v, ok := errors.Descriptions[err]; ok {
		return &ErrorResponse{
			StatucCode: errors.StatusCodes[err],
			Data: echo.Map{
				"error":             err.Error(),
				"error_description": v,
			},
		}
	}

	return &ErrorResponse{
		StatucCode: http.StatusInternalServerError,
		Data: echo.Map{
			"error":             err.Error(),
			"error_description": "The unknowen error happened",
		},
	}
}

// GetStatusCode Get HTTP status code
func (r *ErrorResponse) GetStatusCode() int {
	return r.StatucCode
}

// GetData Get ehco Map data
func (r *ErrorResponse) GetData() echo.Map {
	return r.Data
}
