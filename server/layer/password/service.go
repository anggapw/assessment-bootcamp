package password

import (
	"errors"
	"fmt"
	"password-app/entity"
	"password-app/helper"
	"time"
)

type PasswordService interface {
	CreateNewPassword(userID int, password entity.CreateWebsitePassword) (entity.WebsitePassword, error)
	GetAllPassword() ([]entity.WebsitePassword, error)
	GetPasswordByID(id string) (entity.WebsitePassword, error)
	UpdatePasswordByID(id string, dataInput entity.CreateWebsitePassword) (entity.WebsitePassword, error)
	DeletePasswordByID(id string) (interface{}, error)
}

type passwordService struct {
	repository PasswordRepository
}

func NewService(repository PasswordRepository) *passwordService {
	return &passwordService{repository}
}

func (s *passwordService) CreateNewPassword(userID int, password entity.CreateWebsitePassword) (entity.WebsitePassword, error) {
	var newPass = entity.WebsitePassword{
		UserID:  userID,
		Website: password.Website,
		Title:   password.Title,
	}

	createPassword, err := s.repository.CreatePassword(newPass)

	if err != nil {
		return createPassword, err
	}

	return createPassword, nil
}

func (s *passwordService) GetAllPassword() ([]entity.WebsitePassword, error) {
	passwords, err := s.repository.FindAllPassword()

	if err != nil {
		return passwords, err
	}

	return passwords, nil
}

func (s *passwordService) GetPasswordByID(id string) (entity.WebsitePassword, error) {
	if err := helper.ValidateIDNumber(id); err != nil {
		return entity.WebsitePassword{}, err
	}

	password, err := s.repository.GetOnePassword(id)

	if err != nil {
		return entity.WebsitePassword{}, err
	}

	if password.ID == 0 {
		newError := fmt.Sprintf("password id %s is not found", id)
		return entity.WebsitePassword{}, errors.New(newError)
	}

	return password, nil
}

func (s *passwordService) UpdatePasswordByID(id string, dataInput entity.CreateWebsitePassword) (entity.WebsitePassword, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(id); err != nil {
		return entity.WebsitePassword{}, err
	}

	password, err := s.repository.GetOnePassword(id)

	if err != nil {
		return entity.WebsitePassword{}, err
	}

	if password.ID == 0 {
		newError := fmt.Sprintf("password id %s is not found", id)
		return entity.WebsitePassword{}, errors.New(newError)
	}

	if dataInput.Website != "" || len(dataInput.Website) != 0 {
		dataUpdate["website"] = dataInput.Website
	}
	if dataInput.Title != "" || len(dataInput.Title) != 0 {
		dataUpdate["title"] = dataInput.Title
	}

	dataUpdate["updated_at"] = time.Now()

	passwordUpdated, err := s.repository.UpdatePassword(id, dataUpdate)

	if err != nil {
		return entity.WebsitePassword{}, err
	}

	return passwordUpdated, nil
}

func (s *passwordService) DeletePasswordByID(id string) (interface{}, error) {
	if err := helper.ValidateIDNumber(id); err != nil {
		return nil, err
	}

	password, err := s.repository.GetOnePassword(id)

	if err != nil {
		return nil, err
	}

	if password.ID == 0 {
		newError := fmt.Sprintf("password id %s is not found", id)
		return nil, errors.New(newError)
	}

	status, err := s.repository.DeletePassword(id)

	if err != nil {
		panic(err)
	}

	if status == "error" {
		return nil, errors.New("error delete in internal server")
	}

	msg := fmt.Sprintf("password id %s success deleted", id)

	formatDelete := FormatDelete(msg)

	return formatDelete, nil
}
