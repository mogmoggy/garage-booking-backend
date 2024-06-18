package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	db "github.com/mogmoggy/garage-booking-backend/db/sqlc"
	"github.com/mogmoggy/garage-booking-backend/util"
)

type Server struct {
	store  db.Store
	router *echo.Echo
	config util.Config
}

func NewServer(store db.Store, config util.Config) *Server {
	server := &Server{store: store, config: config}
	router := echo.New()

	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: server.config.AllowedOrigins,
	}))
	router.Use(middleware.Logger())

	v1 := router.Group("/v1")
	v1.GET("/car-registration/:reg", server.GetCarRegistration)

	server.router = router
	return server
}

func (server *Server) Start() error {
	return server.router.Start(server.config.ServerAddress)
}

func errorResponse(errMessage string) map[string]string {
	return map[string]string{"error": errMessage}
}
