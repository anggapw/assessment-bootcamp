package password

import (
	"password-app/entity"

	"gorm.io/gorm"
)

type PasswordRepository interface {
	CreatePassword(password entity.WebsitePassword) (entity.WebsitePassword, error)
	FindAllPassword() ([]entity.WebsitePassword, error)
	GetOnePassword(id string) (entity.WebsitePassword, error)
	UpdatePassword(id string, dataUpdate map[string]interface{}) (entity.WebsitePassword, error)
	DeletePassword(id string) (string, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) CreatePassword(password entity.WebsitePassword) (entity.WebsitePassword, error) {
	if err := r.db.Create(&password).Error; err != nil {
		return password, err
	}

	return password, nil

}

func (r *Repository) FindAllPassword() ([]entity.WebsitePassword, error) {
	var passwords []entity.WebsitePassword

	if err := r.db.Find(&passwords).Error; err != nil {
		return passwords, err
	}

	return passwords, nil
}

func (r *Repository) GetOnePassword(id string) (entity.WebsitePassword, error) {
	var password entity.WebsitePassword

	if err := r.db.Where("id = ?", id).Find(&password).Error; err != nil {
		return password, err
	}

	return password, nil
}

func (r *Repository) UpdatePassword(id string, dataUpdate map[string]interface{}) (entity.WebsitePassword, error) {
	var password entity.WebsitePassword

	if err := r.db.Model(&password).Where("id = ?", id).Updates(dataUpdate).Error; err != nil {
		return password, err
	}

	if err := r.db.Where("id = ?", id).Find(&password).Error; err != nil {
		return password, err
	}

	return password, nil
}

func (r *Repository) DeletePassword(id string) (string, error) {
	if err := r.db.Where("id = ?", id).Delete(&entity.WebsitePassword{}).Error; err != nil {
		return "error", err
	}

	return "success", nil
}
