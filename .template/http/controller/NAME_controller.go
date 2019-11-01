package controller

import (
	"net/http"
	"github.com/labstack/echo"

	"github.com/tamanyan/oauth2-server/app/middleware"
	"github.com/tamanyan/oauth2-server/app/{{lowercaseletters .NAME}}"
)

// {{camelcase .NAME}}Handler represent the httphandler for {{lowercaseletters .NAME}}
type {{camelcase .NAME}}Handler struct {
	{{camelcase .NAME}}Usecase {{lowercaseletters .NAME}}.Usecase
}

// New{{camelcase .NAME}}Handler will initialize the {{lowercaseletters .NAME}} resources endpoint
func New{{camelcase .NAME}}Handler(e *echo.Echo, middleware *middleware.GoMiddleware, us {{lowercaseletters .NAME}}.Usecase) {
	handler := &{{camelcase .NAME}}Handler{
		{{camelcase .NAME}}Usecase: us,
	}
	e.GET("/{{lowercaseletters .NAME}}", handler.GetData)
}

// GetData will get data by ID
func (h *{{camelcase .NAME}}Handler) GetData(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}