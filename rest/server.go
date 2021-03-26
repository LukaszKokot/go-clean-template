package rest

import (
	"github.com/LukaszKokot/go-clean-template/config"

	"github.com/jmoiron/sqlx"
	echo "github.com/labstack/echo/v4"
)

// Server is a simple interface for a web server that can be started and stopped
type Server interface {
	Close() error
	Start(address string) error
}

// Options for configuring the app
type Options struct {
	Web config.WebServerConfiguration
}

// defaultServer is the default implementation of Server
type defaultServer struct {
	e        *echo.Echo
	database *sqlx.DB
}

// NewServer builds a new startable web server
// We're going to hide here that we use Echo, nobody needs to know from the outside
func NewServer(options Options) (Server, error) {
	e := echo.New()

	// This is the place where

	s := &defaultServer{
		e:        e,
		database: nil, // We should connect to the database and provide it here
	}

	// When we have our server and database, we can now configure the routes that
	// will connect HTTP endpoints to our handlers.
	//
	// For example:
	// s.e.POST("users", s.h.user.Register)

	return s, nil
}

func (s defaultServer) Close() error {
	return s.e.Close()
}

func (s defaultServer) Start(address string) error {
	return s.e.Start(address)
}
