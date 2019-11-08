package controller

import (
	"net/http"
	"github.com/labstack/echo"

	"github.com/tamanyan/oauth2-server/middleware"
	"github.com/tamanyan/oauth2-server/user"
)

// UserHandler represent the httphandler for user
type UserHandler struct {
	UserUsecase user.Usecase
}

// NewUserHandler will initialize the user resources endpoint
func NewUserHandler(e *echo.Echo, middleware *middleware.GoMiddleware, us user.Usecase) {
	handler := &UserHandler{
		UserUsecase: us,
	}
	e.GET("/user", handler.GetData)
}

// GetData will get data by ID
func (h *UserHandler) GetData(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}