package user

import (
	"io"

	"github.com/LukaszKokot/go-clean-template/pkg/storage"
	"github.com/LukaszKokot/go-clean-template/pkg/usecases"
)

type defaultUserAvatarUsecase struct {
	user       storage.UserRepository
	userAvatar storage.UserAvatarRepository
}

func NewDefaultUserAvatarUsecase(
	user storage.UserRepository,
	userAvatar storage.UserAvatarRepository,
) usecases.UserAvatarUsecase {
	return defaultUserAvatarUsecase{
		user:       user,
		userAvatar: userAvatar,
	}
}
func (u defaultUserAvatarUsecase) CreateOrUpdate(userID string, blob io.Reader) error {
	// First we upload the blob and get the public URL
	avatarURL, err := u.userAvatar.CreateOrUpdate(userID, blob)
	if err != nil {
		return err
	}

	// Then we find the user
	user, err := u.user.FindByID(userID)
	if err != nil {
		return err
	}

	// ...and update the info
	user.AvatarURL = *avatarURL
	err = u.user.Update(*user)
	if err != nil {
		return err
	}

	return nil
}
