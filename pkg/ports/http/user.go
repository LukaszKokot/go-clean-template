package http

import (
	"net/http"

	"github.com/LukaszKokot/go-clean-template/pkg/models"
	"github.com/LukaszKokot/go-clean-template/pkg/usecases"
	"github.com/labstack/echo/v4"
)

type UserPort interface {
	Register(ctx echo.Context) error
}

type defaultUserPort struct {
	user usecases.UserUsecase
}

// NewDefaultUserPort returns the default implementation of the user port.
func NewDefaultUserPort(user usecases.UserUsecase) UserPort {
	return defaultUserPort{
		user: user,
	}
}

func (p defaultUserPort) Register(ctx echo.Context) error {
	var newUser models.NewUser
	if err := ctx.Bind(&newUser); err != nil {
		return ctx.String(http.StatusBadRequest, "Unreadable form")
	}

	user, err := p.user.Register(newUser)
	if err != nil {
		// We would log that error here
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, user)
}
