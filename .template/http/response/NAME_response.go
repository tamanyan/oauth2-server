package response

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/tamanyan/oauth2-server/app/common"
	"github.com/tamanyan/oauth2-server/{{lowercaseletters .NAME}}"
)

// {{camelcase .NAME}}Response usecase
type {{camelcase .NAME}}Response struct {
	StatucCode int
	Data       echo.Map
}

// New{{camelcase .NAME}}Response will create a new {{camelcase .NAME}}Response
func New{{camelcase .NAME}}Response(data: interface{}) common.Response {
	return &{{camelcase .NAME}}Response{
		StatucCode: http.StatusOK,
		Data:       echo.Map{},
	}
}

// GetStatusCode Get HTTP status code
func (r *{{camelcase .NAME}}Response) GetStatusCode() int {
	return r.StatucCode
}

// GetData Get ehco Map data
func (r *{{camelcase .NAME}}Response) GetData() echo.Map {
	return r.Data
}
