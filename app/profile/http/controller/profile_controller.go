package controller

import (
	"net/http"
	"github.com/labstack/echo"

	"github.com/tamanyan/oauth2-server/app/middleware"
	"github.com/tamanyan/oauth2-server/app/profile"
)

// ProfileHandler represent the httphandler for profile
type ProfileHandler struct {
	ProfileUsecase profile.Usecase
}

// NewProfileHandler will initialize the profile resources endpoint
func NewProfileHandler(e *echo.Echo, middleware *middleware.GoMiddleware, us profile.Usecase) {
	handler := &ProfileHandler{
		ProfileUsecase: us,
	}
	e.GET("/profile", handler.GetData)
}

// GetData will get data by ID
func (h *ProfileHandler) GetData(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}