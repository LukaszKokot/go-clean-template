package storage

import (
	"io"

	"github.com/LukaszKokot/go-clean-template/pkg/models"
)

// UserRepository is for interacting with users
type UserRepository interface {
	FindByEmail(email string) (*models.User, error)
	FindByID(userID string) (*models.User, error)
	Create(user models.User) error
	Update(user models.User) error
}

// UserAvatarRepository is for interacting with user avatars (blobs)
type UserAvatarRepository interface {
	CreateOrUpdate(userID string, blob io.Reader) (*string, error)
}
