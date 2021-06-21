package user

import (
	"assessment/entity"

	"gorm.io/gorm"
)

type Repository interface {
	RShowAllUser() ([]entity.User, error)
	RRegisterUser(user entity.User) (entity.User, error)
	RFindUserByEmail(email string) (entity.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) RShowAllUser() ([]entity.User, error) {
	var user []entity.User

	err := r.db.Find(&user).Error
	if err != nil {
		return user, nil
	}

	return user, nil
}

func (r *repository) RRegisterUser(user entity.User) (entity.User, error) {

	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindUserByEmail(email string) (entity.User, error) {
	var user entity.User

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
