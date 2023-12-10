package services

import (
	"Golang-RESTful-APi/entities"
	"Golang-RESTful-APi/models"
	"Golang-RESTful-APi/models/converter"
	"Golang-RESTful-APi/repository"
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AddressService struct {
	DB                *gorm.DB
	Log               *logrus.Logger
	Validate          *validator.Validate
	AddressRepository *repository.AddressRepository
	ContactRepository *repository.ContactRepository
}

func NewAddressService(DB *gorm.DB, Log *logrus.Logger, Validate *validator.Validate, addressRepository *repository.AddressRepository, ContactRepository *repository.ContactRepository) *AddressService {
	return &AddressService{
		DB:                DB,
		Log:               Log,
		Validate:          Validate,
		AddressRepository: addressRepository,
		ContactRepository: ContactRepository,
	}
}

func (c *AddressService) Create(ctx context.Context, request *models.CreateAddressRequest) (*models.AddressResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, fiber.ErrBadRequest
	}

	contact := new(entities.Contact)
	if err := c.ContactRepository.FindByIdAndUsername(tx, contact, request.ContactId, request.Username); err != nil {
		c.Log.WithError(err).Error("failed to find contact")
		return nil, fiber.ErrNotFound
	}

	address := &entities.Address{
		ID:         uuid.NewString(),
		ContactId:  contact.ID,
		Street:     request.Street,
		City:       request.City,
		Province:   request.Province,
		PostalCode: request.PostalCode,
		Country:    request.Country,
	}

	if err := c.AddressRepository.Create(tx, address); err != nil {
		c.Log.WithError(err).Error("failed to create address")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.AddressToResponse(address), nil

}

func (c *AddressService) Update(ctx context.Context, request *models.UpdateAddressRequest) (*models.AddressResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("failed to validate request body")
		return nil, fiber.ErrBadRequest
	}

	contact := new(entities.Contact)
	if err := c.ContactRepository.FindByIdAndUsername(tx, contact, request.ContactId, request.Username); err != nil {
		c.Log.WithError(err).Error("failed to find contact")
		return nil, fiber.ErrNotFound
	}

	address := new(entities.Address)
	if err := c.AddressRepository.FindByIdAndContactId(tx, address, request.ID, contact.ID); err != nil {
		c.Log.WithError(err).Error("failed to find address")
		return nil, fiber.ErrNotFound
	}

	address.Street = request.Street
	address.City = request.City
	address.Province = request.Province
	address.PostalCode = request.PostalCode
	address.Country = request.Country

	if err := c.AddressRepository.Update(tx, address); err != nil {
		c.Log.WithError(err).Error("failed to update address")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.AddressToResponse(address), nil
}

func (c *AddressService) Get(ctx context.Context, request *models.GetAddressRequest) (*models.AddressResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	contact := new(entities.Contact)
	if err := c.ContactRepository.FindByIdAndUsername(tx, contact, request.ContactId, request.Username); err != nil {
		c.Log.WithError(err).Error("failed to find contact")
		return nil, fiber.ErrNotFound
	}

	address := new(entities.Address)
	if err := c.AddressRepository.FindByIdAndContactId(tx, address, request.ID, request.ContactId); err != nil {
		c.Log.WithError(err).Error("failed to find address")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.AddressToResponse(address), nil
}

func (c *AddressService) Delete(ctx context.Context, request *models.DeleteAddressRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	contact := new(entities.Contact)
	if err := c.ContactRepository.FindByIdAndUsername(tx, contact, request.ContactId, request.Username); err != nil {
		c.Log.WithError(err).Error("failed to find contact")
		return fiber.ErrNotFound
	}

	address := new(entities.Address)
	if err := c.AddressRepository.FindByIdAndContactId(tx, address, request.ID, request.ContactId); err != nil {
		c.Log.WithError(err).Error("failed to find address")
		return fiber.ErrNotFound
	}

	if err := c.AddressRepository.Delete(tx, address); err != nil {
		c.Log.WithError(err).Error("failed to delete address")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return fiber.ErrInternalServerError
	}

	return nil
}

func (c *AddressService) List(ctx context.Context, request *models.ListAddressRequest) ([]models.AddressResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	contact := new(entities.Contact)
	if err := c.ContactRepository.FindByIdAndUsername(tx, contact, request.ContactId, request.Username); err != nil {
		c.Log.WithError(err).Error("failed to find contact")
		return nil, fiber.ErrNotFound
	}

	addresses, err := c.AddressRepository.FindAllByContactId(tx, contact.ID)
	if err != nil {
		c.Log.WithError(err).Error("failed to find addresses")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	responses := make([]models.AddressResponse, len(addresses))
	for i, address := range addresses {
		responses[i] = *converter.AddressToResponse(&address)
	}

	return responses, nil
}
