package user

import (
	"github.com/LukaszKokot/go-clean-template/pkg/models"
	"github.com/LukaszKokot/go-clean-template/pkg/storage"
	"github.com/LukaszKokot/go-clean-template/pkg/usecases"
)

type defaultUserUsecase struct {
	user storage.UserRepository
}

func NewDefaultUserUsecase(user storage.UserRepository) usecases.UserUsecase {
	return defaultUserUsecase{
		user: user,
	}
}

func (u defaultUserUsecase) Login(email, password string) (*models.User, error) {
	user, err := u.user.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	// Here we should hash the "password" and check it against
	// the hashed password stored in "user"

	// .. and if it matches, we ca return that user, otherwise we would return an
	// error stating that credentials are incorrect.
	return user, nil
}

func (u defaultUserUsecase) Register(newUser models.NewUser) (*models.User, error) {
	user := models.User{
		ID:           "", // generate some random uuid here
		Email:        newUser.Email,
		FirstName:    newUser.FirstName,
		LastName:     newUser.LastName,
		PasswordHash: "", // hash the password and set it here
	}

	err := u.user.Create(user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
