package middleware

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/tamanyan/oauth2-server/common"
	"github.com/tamanyan/oauth2/errors"
)

// GoMiddleware represent the data-struct for middleware
type GoMiddleware struct {
	// another stuff , may be needed by middleware
}

// JWT will handle the JWT middleware
func (m *GoMiddleware) JWT() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    []byte(os.Getenv("JWT_SECRET_KEY")),
		SigningMethod: "HS512",
		ErrorHandler: func(err error) error {
			errRes := common.NewErrorResponse(errors.ErrUnauthorized)
			return echo.NewHTTPError(
				errRes.GetStatusCode(),
				errRes.GetData(),
			)
		},
	})
}

// InitMiddleware intialize the middleware
func InitMiddleware() *GoMiddleware {
	return &GoMiddleware{}
}
