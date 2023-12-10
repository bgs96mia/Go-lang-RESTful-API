package repository

import (
	"Golang-RESTful-APi/entities"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AddressRepository struct {
	Repository[entities.Address]
	Log *logrus.Logger
}

func NewAddressRepository(log *logrus.Logger) *AddressRepository {
	return &AddressRepository{
		Log: log,
	}
}

func (r *AddressRepository) FindByIdAndContactId(tx *gorm.DB, address *entities.Address, id string, contactId string) error {
	return tx.Where("id = ? AND contact_id = ?", id, contactId).First(address).Error
}

func (r *AddressRepository) FindAllByContactId(tx *gorm.DB, contactId string) ([]entities.Address, error) {
	var addresses []entities.Address
	if err := tx.Where("contact_id = ?", contactId).Find(&addresses).Error; err != nil {
		return nil, err
	}
	return addresses, nil
}
