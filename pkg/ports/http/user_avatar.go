package http

import (
	"fmt"
	"net/http"

	"github.com/LukaszKokot/go-clean-template/pkg/usecases"
	"github.com/labstack/echo/v4"
)

type UserAvatarPort interface {
	CreateOrUpdate(ctx echo.Context) error
}

type defaultUserAvatarPort struct {
	userAvatar usecases.UserAvatarUsecase
}

// NewDefaultUserAvatarPort returns the default implementation of the user avatar port.
func NewDefaultUserAvatarPort(userAvatar usecases.UserAvatarUsecase) UserAvatarPort {
	return defaultUserAvatarPort{
		userAvatar: userAvatar,
	}
}

func (p defaultUserAvatarPort) CreateOrUpdate(ctx echo.Context) error {
	// We get the id of the user from the URL placeholder
	// Say: PUT /users/{userID}/avatar
	userID := ctx.Param("userID")

	// We'll try to get the image part of the multipart request
	formFile, err := ctx.FormFile("image")
	if err != nil {
		return ctx.String(http.StatusBadRequest, fmt.Sprintf("No image found in multipart request: %s", err.Error()))
	}
	file, err := formFile.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	// We have everything, we can call the usecase to do the operation.
	err = p.userAvatar.CreateOrUpdate(userID, file)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusNoContent)
}
