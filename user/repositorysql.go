package user

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type repoSql struct {
	DB *sqlx.DB
}

type RepoSQL interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindByID(ID int) (User, error)
	UpdateUser(user User) (User, error)
}

func NewRepositorySQL(DB *sqlx.DB) *repoSql {
	return &repoSql{DB}
}

func (r *repoSql) Save(user User) (User, error) {
	querry := `
	INSERT INTO 
		users
		(
			name,
			email,
			occupation,
			password_hash,
			role,
			created_at,
			updated_at
		) 
		VALUES
		(
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7
		)
	
	`

	_, err := r.DB.Exec(querry, user.Name, user.Email, user.Occupation, user.PasswordHash, user.Role, time.Now(), time.Now())

	if err != nil {
		return User{}, err
	}

	var userdb UserDB

	return User{
		ID:             int(userdb.ID.Int64),
		Name:           userdb.Name.String,
		Occupation:     userdb.Occupation.String,
		Email:          userdb.Email.String,
		PasswordHash:   userdb.PasswordHash.String,
		AvatarFileName: userdb.AvatarFileName.String,
		Role:           userdb.Role.String,
	}, nil
}

func (r *repoSql) FindByEmail(email string) (User, error) {
	querry := ` 
	SELECT 
		id,
		name,
		email,
		password,
		created_at,
		avatar_file_name
	WHERE 
		email = $1
	`

	user := UserDB{}
	err := r.DB.Get(&user, querry, email)
	if err != nil {
		return User{}, err
	}

	return User{
		ID:             int(user.ID.Int64),
		Name:           user.Name.String,
		Occupation:     user.Occupation.String,
		Email:          user.Email.String,
		PasswordHash:   user.PasswordHash.String,
		AvatarFileName: user.AvatarFileName.String,
		Role:           user.Role.String,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
	}, nil
}

func (r *repoSql) FindByID(ID int) (User, error) {
	querry := `
	SELECT
		id,
		name,
		occupation,
		email,
		created_at,
		updated_at,
		avatar_file_name
	FROM
		user
	WHERE 
		id = 1$
	`
	var user2 UserDB
	err := r.DB.Get(&user2, querry, ID)

	if err != nil {
		return User{}, err
	}

	return User{
		ID:             int(user2.ID.Int64),
		Name:           user2.Name.String,
		Occupation:     user2.Occupation.String,
		Email:          user2.Email.String,
		CreatedAt:      user2.CreatedAt,
		UpdatedAt:      user2.UpdatedAt,
		AvatarFileName: user2.AvatarFileName.String,
	}, nil

}

func (r *repoSql) UpdateUser(user User) (User, error) {
	querry := `
			UPDATE
				users
			SET 
				avatar_file_name = $1
			WHERE
				id = $2
		`
	var userdb UserDB
	_, err := r.DB.Exec(querry, user.AvatarFileName, user.ID)
	if err != nil {
		return User{}, err
	}

	return User{
		ID:             int(userdb.ID.Int64),
		AvatarFileName: userdb.AvatarFileName.String,
		Email:          userdb.Email.String,
	}, nil

}
