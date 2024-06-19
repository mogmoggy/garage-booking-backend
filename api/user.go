package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mogmoggy/garage-booking-backend/token"
)

func (server *Server) getBookings(ctx echo.Context) error {

	authPayload := ctx.Get(authorizationPayloadKey).(*token.Payload)
	fmt.Println(authPayload)
	return ctx.JSON(http.StatusOK, struct{}{})
}
