package context

import (
	"github.com/LukaszKokot/go-clean-template/pkg/ports/http"
	"github.com/LukaszKokot/go-clean-template/pkg/storage/db"
	"github.com/LukaszKokot/go-clean-template/pkg/storage/s3"
	"github.com/LukaszKokot/go-clean-template/pkg/usecases"
	"github.com/LukaszKokot/go-clean-template/pkg/usecases/user"

	"github.com/jmoiron/sqlx"
)

// NewUserUsecase returns a new instance of the user use case
func NewUserUsecase(database *sqlx.DB) usecases.UserUsecase {
	return user.NewDefaultUserUsecase(
		db.NewUserRepository(database),
	)
}

// NewUserAvatarUsecase returns a new instance of the user avatar use case
func NewUserAvatarUsecase(database *sqlx.DB) usecases.UserAvatarUsecase {
	return user.NewDefaultUserAvatarUsecase(
		db.NewUserRepository(database),
		s3.NewUserAvatarRepository(),
	)
}

// NewUserPort returns a new HTTP user port
func NewUserPort(database *sqlx.DB) http.UserPort {
	return http.NewDefaultUserPort(
		NewUserUsecase(database),
	)
}

// NewUserAvatarPort returns a new HTTP user avatar port
func NewUserAvatarPort(database *sqlx.DB) http.UserAvatarPort {
	return http.NewDefaultUserAvatarPort(
		NewUserAvatarUsecase(database),
	)
}
