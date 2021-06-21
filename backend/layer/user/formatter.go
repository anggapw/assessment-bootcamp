package user

import "assessment/entity"

type UserFormat struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Address  string `json:"address"`
	Email    string `json:"email"`
}

func FormatUser(user entity.User) UserFormat {
	var formatuser = UserFormat{
		ID:       user.ID,
		FullName: user.FullName,
		Address:  user.Address,
		Email:    user.Email,
	}

	return formatuser
}
