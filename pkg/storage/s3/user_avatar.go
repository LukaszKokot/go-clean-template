package s3

import (
	"io"

	"github.com/LukaszKokot/go-clean-template/pkg/storage"
)

type defaultUserAvatarRepository struct {
	// Here we might need some s3 info such as the secret API keys
	// to connect to AWS
}

func NewUserAvatarRepository() storage.UserAvatarRepository {
	return defaultUserAvatarRepository{}
}

func (r defaultUserAvatarRepository) CreateOrUpdate(userID string, blob io.Reader) error {
	// This is where we upload our blob and make sure it has a unique id so we
	// can override it if it already exists.
	return nil
}
