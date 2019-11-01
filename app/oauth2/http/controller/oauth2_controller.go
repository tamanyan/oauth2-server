package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"

	"github.com/tamanyan/oauth2-server/app/common"
	"github.com/tamanyan/oauth2-server/app/middleware"
	_oauth2 "github.com/tamanyan/oauth2-server/app/oauth2"
	"github.com/tamanyan/oauth2-server/app/oauth2/http/request"
	"github.com/tamanyan/oauth2-server/app/oauth2/http/response"
	"github.com/tamanyan/oauth2-server/errors"
	"github.com/tamanyan/oauth2-server/oauth2"
)

// OAuth2Handler  represent the httphandler for oauth2
type OAuth2Handler struct {
	OAuth2Usecase _oauth2.Usecase
}

// NewOAuth2Handler will initialize the articles/ resources endpoint
func NewOAuth2Handler(e *echo.Echo, middleware *middleware.GoMiddleware, manager oauth2.Manager, us _oauth2.Usecase) {
	handler := &OAuth2Handler{
		OAuth2Usecase: us,
	}
	e.POST("/oauth2/token", handler.IssueAccessToken)

	r := e.Group("/oauth2")
	r.Use(middleware.JWT())
	r.DELETE("/token", handler.RevokeAccessToken)
	r.GET("/verify", handler.VerifyAccessToken)
}

func validateRequest(c echo.Context, request interface{}) error {
	if err := c.Bind(request); err != nil {
		return errors.ErrInvalidRequest
	}

	validate := validator.New()
	err := validate.Struct(request)

	if err != nil {
		return errors.ErrInvalidRequest
	}

	return nil
}

// HandleJwtError will handle jwt error
func (h *OAuth2Handler) HandleJwtError(err error) error {
	errRes := common.NewErrorResponse(errors.ErrUnauthorized)
	return echo.NewHTTPError(
		errRes.GetStatusCode(),
		errRes.GetData(),
	)
}

// IssueAccessToken will create access token
func (h *OAuth2Handler) IssueAccessToken(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	gt := oauth2.GrantType(c.Request().FormValue("grant_type"))

	if gt.String() == "" {
		errRes := common.NewErrorResponse(errors.ErrUnsupportedGrantType)
		return echo.NewHTTPError(
			errRes.GetStatusCode(),
			errRes.GetData(),
		)
	}

	switch gt {
	case oauth2.PasswordCredentials:
		request := request.OAuth2PasswordCredentialRequest{}
		err := validateRequest(c, &request)

		if err != nil {
			errRes := common.NewErrorResponse(err)
			return echo.NewHTTPError(
				errRes.GetStatusCode(),
				errRes.GetData(),
			)
		}

		ti, err := h.OAuth2Usecase.IssuePasswordCredentialAccessToken(ctx, request)

		if err != nil {
			errRes := common.NewErrorResponse(err)
			return echo.NewHTTPError(
				errRes.GetStatusCode(),
				errRes.GetData(),
			)
		}

		res := response.NewOAuth2Response(ti)
		return c.JSON(res.GetStatusCode(), res.GetData())
	case oauth2.Refreshing:
		request := request.OAuth2RefreshTokenRequest{}
		err := validateRequest(c, &request)

		if err != nil {
			errRes := common.NewErrorResponse(err)
			return echo.NewHTTPError(
				errRes.GetStatusCode(),
				errRes.GetData(),
			)
		}

		ti, err := h.OAuth2Usecase.IssueRefreshAccessToken(ctx, request)

		if err != nil {
			errRes := common.NewErrorResponse(err)
			return echo.NewHTTPError(
				errRes.GetStatusCode(),
				errRes.GetData(),
			)
		}

		res := response.NewOAuth2Response(ti)
		return c.JSON(res.GetStatusCode(), res.GetData())
	case oauth2.ClientCredentials:
		request := request.OAuth2ClientCredentialRequest{}
		err := validateRequest(c, &request)

		if err != nil {
			errRes := common.NewErrorResponse(err)
			return echo.NewHTTPError(
				errRes.GetStatusCode(),
				errRes.GetData(),
			)
		}

		ti, err := h.OAuth2Usecase.IssueClientCredentialAccessToken(ctx, request)

		if err != nil {
			errRes := common.NewErrorResponse(err)
			return echo.NewHTTPError(
				errRes.GetStatusCode(),
				errRes.GetData(),
			)
		}

		res := response.NewOAuth2Response(ti)
		return c.JSON(res.GetStatusCode(), res.GetData())
	}

	errRes := common.NewErrorResponse(errors.ErrUnsupportedGrantType)
	return c.JSON(errRes.GetStatusCode(), errRes.GetData())
}

// RevokeAccessToken will revoke access token
func (h *OAuth2Handler) RevokeAccessToken(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	token := c.Get("user").(*jwt.Token)

	if token == nil {
		errRes := common.NewErrorResponse(errors.ErrUnauthorized)
		return echo.NewHTTPError(
			errRes.GetStatusCode(),
			errRes.GetData(),
		)
	}

	err := h.OAuth2Usecase.RevokeAccessToken(ctx, token.Raw)

	if err != nil {
		errRes := common.NewErrorResponse(err)
		return echo.NewHTTPError(
			errRes.GetStatusCode(),
			errRes.GetData(),
		)
	}

	return c.NoContent(http.StatusOK)
}

// VerifyAccessToken will verify access token
func (h *OAuth2Handler) VerifyAccessToken(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	token := c.Get("user").(*jwt.Token)

	if token == nil {
		errRes := common.NewErrorResponse(errors.ErrUnauthorized)
		return echo.NewHTTPError(
			errRes.GetStatusCode(),
			errRes.GetData(),
		)
	}

	ti, err := h.OAuth2Usecase.VerifyAccessToken(ctx, token.Raw)

	if err != nil {
		errRes := common.NewErrorResponse(err)
		return echo.NewHTTPError(
			errRes.GetStatusCode(),
			errRes.GetData(),
		)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"scope":      ti.GetScope(),
		"client_id":  ti.GetClientID(),
		"expires_in": int64(ti.GetAccessExpiresIn() / time.Second),
	})
}
