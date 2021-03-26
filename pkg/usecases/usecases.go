package usecases

import (
	"io"

	"github.com/LukaszKokot/go-clean-template/pkg/models"
)

type UserUsecase interface {
	Login(email, password string) (*models.User, error)
	Register(user models.NewUser) (*models.User, error)
}

type UserAvatarUsecase interface {
	CreateOrUpdate(userID string, blob io.Reader) error
}
