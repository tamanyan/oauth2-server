package controller

import (
	"net/http"
	"github.com/labstack/echo"

	"github.com/tamanyan/oauth2-server/app/middleware"
	"github.com/tamanyan/oauth2-server/app/admin"
)

// AdminHandler represent the httphandler for admin
type AdminHandler struct {
	AdminUsecase admin.Usecase
}

// NewAdminHandler will initialize the admin resources endpoint
func NewAdminHandler(e *echo.Echo, middleware *middleware.GoMiddleware, us admin.Usecase) {
	handler := &AdminHandler{
		AdminUsecase: us,
	}
	e.GET("/admin", handler.GetData)
}

// GetData will get data by ID
func (h *AdminHandler) GetData(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}