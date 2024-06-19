package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	db "github.com/mogmoggy/garage-booking-backend/db/sqlc"
	"github.com/mogmoggy/garage-booking-backend/token"
	"github.com/mogmoggy/garage-booking-backend/util"
)

type Server struct {
	store      db.Store
	router     *echo.Echo
	config     util.Config
	tokenMaker token.Maker
}

func NewServer(store db.Store, config util.Config) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSecret) // can replace this maker with other token other than jwt
	if err != nil {
		return nil, err
	}

	server := &Server{store: store, config: config, tokenMaker: tokenMaker}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := echo.New()

	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: server.config.AllowedOrigins,
	}))

	v1 := router.Group("/v1")
	user := v1.Group("/user")
	user.Use(authMiddleware(server.tokenMaker))

	user.GET("/bookings", server.getBookings)

	v1.GET("/car-registration/:reg", server.GetCarRegistration)

	server.router = router
}

func (server *Server) Start() error {
	return server.router.Start(server.config.ServerAddress)
}

func errorResponse(errMessage error) map[string]string {
	return map[string]string{"error": errMessage.Error()}
}
