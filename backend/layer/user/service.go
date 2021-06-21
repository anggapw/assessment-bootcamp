package user

import (
	"assessment/entity"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	SShowAllUser() ([]UserFormat, error)
	SRegisterUser(user entity.User) (UserFormat, error)
	SLoginUser(input entity.LoginUserInput) (UserFormat, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) SShowAllUser() ([]UserFormat, error) {
	user, err := s.repository.RShowAllUser()
	var formatUsers []UserFormat

	for _, user := range user {
		formatUser := FormatUser(user)
		formatUsers = append(formatUsers, formatUser)
	}

	if err != nil {
		return formatUsers, err
	}

	return formatUsers, nil
}

func (s *service) SRegisterUser(user entity.User) (UserFormat, error) {
	genPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)

	if err != nil {
		return UserFormat{}, err
	}

	var newUser = entity.User{
		FullName: user.FullName,
		Address:  user.Address,
		Email:    user.Email,
		Password: string(genPassword),
	}

	createUser, err := s.repository.RRegisterUser(newUser)
	formatUser := FormatUser(createUser)

	if err != nil {
		return formatUser, err
	}

	return formatUser, err
}

func (s *service) SLoginUser(input entity.LoginUserInput) (entity.User, error) {
	user, err := s.repository.RFindUserByEmail(input.Email)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user id %v not found", user.ID)
		return user, errors.New(newError)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return user, errors.New("password salah")
	}

	return user, nil
}
