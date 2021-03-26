package models

// User represents a single user
type User struct {
	ID           string `db:"id" json:"id"`
	Email        string `db:"email" json:"email"`
	FirstName    string `db:"first_name" json:"firstName"`
	LastName     string `db:"last_name" json:"lastName"`
	AvatarURL    string `db:"avatar_url" json:"avatarURL"`
	PasswordHash string `db:"password" json:"-"`
}

// NewUser is for registering a new user
type NewUser struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
}
