package rest

import (
	"github.com/LukaszKokot/go-clean-template/config"
	"github.com/LukaszKokot/go-clean-template/pkg/context"
	"github.com/LukaszKokot/go-clean-template/pkg/ports/http"

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

	ports struct {
		user       http.UserPort
		userAvatar http.UserAvatarPort
	}
}

// NewServer builds a new startable web server
// We're going to hide here that we use Echo, nobody needs to know from the outside
func NewServer(options Options) (Server, error) {
	e := echo.New()
	var database *sqlx.DB // We should connect to the database and provide it here

	s := &defaultServer{
		e:        e,
		database: database,
	}

	// We configure our HTTP ports
	s.ports.user = context.NewUserPort(database)
	s.ports.userAvatar = context.NewUserAvatarPort(database)

	// And finally we connect the HTTP endpoints with our ports
	s.e.POST("users", s.ports.user.Register)
	s.e.PUT("users/:userID", s.ports.userAvatar.CreateOrUpdate)

	// What's missing here?
	//
	// The configured routes don't have any security middleware.
	// For example, the avatar update route should have a middleware that
	// checks whether the user making the request has the right credentials
	// for updating that user's adapter.

	return s, nil
}

func (s defaultServer) Close() error {
	return s.e.Close()
}

func (s defaultServer) Start(address string) error {
	return s.e.Start(address)
}
