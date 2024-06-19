package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mogmoggy/garage-booking-backend/token"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func authMiddleware(tokenMaker token.Maker) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			authortizationHeader := ctx.Request().Header.Get(authorizationHeaderKey)
			if len(authortizationHeader) == 0 {
				err := errors.New("authorization header is not provided")
				return echo.NewHTTPError(http.StatusUnauthorized, errorResponse(err))
			}

			fields := strings.Fields(authortizationHeader)
			if len(fields) < 2 {
				err := errors.New("invalid authorization header format")
				return echo.NewHTTPError(http.StatusUnauthorized, errorResponse(err))
			}

			authorizationType := strings.ToLower(fields[0])
			if authorizationType != authorizationTypeBearer {
				err := fmt.Errorf("unsupported authorization type %v", authorizationType)
				return echo.NewHTTPError(http.StatusUnauthorized, errorResponse(err))
			}

			accessToken := fields[1]
			payload, err := tokenMaker.VerifyToken(accessToken)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, errorResponse(err))
			}

			ctx.Set(authorizationPayloadKey, payload)
			return next(ctx)
		}
	}
}
