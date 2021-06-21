package user

import "assessment/entity"

type Repository interface {
	RShowAllUser() ([]entity.User, error)
	RRegisterUser(user entity.User) (entity.User, error)
	RFindUserByEmail(email string) (entity.User, error)
}
