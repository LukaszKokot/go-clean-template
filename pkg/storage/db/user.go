package db

import (
	"github.com/LukaszKokot/go-clean-template/pkg/models"
	"github.com/LukaszKokot/go-clean-template/pkg/storage"

	"github.com/jmoiron/sqlx"
)

type defaultUserRepository struct {
	db sqlx.Ext
}

// NewUserRepository returns a new instance of the user repository
func NewUserRepository(db sqlx.Ext) storage.UserRepository {
	return defaultUserRepository{db}
}

func (r defaultUserRepository) Create(user models.User) error {
	return nil
}

func (r defaultUserRepository) FindByID(id string) (*models.User, error) {
	return nil, nil
}

func (r defaultUserRepository) Update(user models.User) error {
	return nil
}
