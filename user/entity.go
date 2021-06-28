package user

import (
	"database/sql"
	"time"
)

type User struct {
	ID             int
	Name           string
	Occupation     string
	Email          string
	PasswordHash   string
	AvatarFileName string
	Role           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type UserDB struct {
	ID             sql.NullInt64  `db:"id"`
	Name           sql.NullString `db:"name"`
	Occupation     sql.NullString `db:"occupation"`
	Email          sql.NullString `db:"email"`
	PasswordHash   sql.NullString `db:"password_hash"`
	AvatarFileName sql.NullString `db:"avatarfilename"`
	Role           sql.NullString `db:"role"`
	CreatedAt      time.Time	`db:"created_at"`
	UpdatedAt      time.Time	`db:"updated_at"`
}

