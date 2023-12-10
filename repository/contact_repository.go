package repository

import (
	"Golang-RESTful-APi/entities"
	"Golang-RESTful-APi/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ContactRepository struct {
	Repository[entities.Contact]
	Log *logrus.Logger
}

func NewContactRepository(log *logrus.Logger) *ContactRepository {
	return &ContactRepository{
		Log: log,
	}
}

func (r *ContactRepository) FindByIdAndUsername(db *gorm.DB, contact *entities.Contact, id string, username string) error {
	return db.Where("id = ? AND username = ?", id, username).Take(contact).Error
}

func (r *ContactRepository) Search(db *gorm.DB, request *models.SearchContactRequest) ([]entities.Contact, int64, error) {
	var contacts []entities.Contact
	if err := db.Scopes(r.FilterContact(request)).Offset((request.Page - 1) * request.Size).Limit(request.Size).Find(&contacts).Error; err != nil {
		return nil, 0, err
	}

	var total int64 = 0
	if err := db.Model(&entities.Contact{}).Scopes(r.FilterContact(request)).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return contacts, total, nil
}

func (r *ContactRepository) FilterContact(request *models.SearchContactRequest) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		tx = tx.Where("username = ?", request.Username)

		if name := request.Name; name != "" {
			name = "%" + name + "%"
			tx = tx.Where("first_name LIKE ? OR last_name LIKE ?", name, name)
		}

		if phone := request.Phone; phone != "" {
			phone = "%" + phone + "%"
			tx = tx.Where("phone LIKE ?", phone)
		}

		if email := request.Email; email != "" {
			email = "%" + email + "%"
			tx = tx.Where("email LIKE ?", email)
		}

		return tx
	}
}
