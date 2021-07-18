package user

import "BWA/rpcp"

type UserFormater struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
	ImageURL   string `json:"image_url"`
}

func (f *UserFormater) ToDataProto(data *rpcp.RegisterUserData) {
	data.ID = int32(f.ID)
	data.Name = f.Name
	data.Occupation = f.Occupation
	data.Email = f.Email
	data.Token = f.Token
	data.ImageURL = f.ImageURL
}

func FormatUser(user User, token string) UserFormater {
	formatter := UserFormater{
		ID:         user.ID,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
		ImageURL:   user.AvatarFileName,
	}

	return formatter
}
