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

type ContactService struct {
	DB                *gorm.DB
	Log               *logrus.Logger
	Validate          *validator.Validate
	ContactRepository *repository.ContactRepository
}

func NewContactService(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, contactRepository *repository.ContactRepository) *ContactService {
	return &ContactService{
		DB:                db,
		Log:               logger,
		Validate:          validate,
		ContactRepository: contactRepository,
	}

}

func (c *ContactService) Create(ctx context.Context, request *models.CreateContactRequest) (*models.ContactResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validation request body")
		return nil, fiber.ErrBadRequest
	}
	
	contact := &entities.Contact{
		ID:        uuid.New().String(),
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Phone:     request.Phone,
		Username:  request.Username,
	}

	if err := c.ContactRepository.Create(tx, contact); err != nil {
		c.Log.WithError(err).Error("error creating contact")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error creating contact")
		return nil, fiber.ErrInternalServerError
	}

	return converter.ContactToResponse(contact), nil
}

func (c *ContactService) Get(ctx context.Context, request *models.GetContactRequest) (*models.ContactResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validation request body")
		return nil, fiber.ErrBadRequest
	}

	contact := new(entities.Contact)
	if err := c.ContactRepository.FindByIdAndUsername(tx, contact, request.ID, request.Username); err != nil {
		c.Log.WithError(err).Error("error getting contact")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting contact")
		return nil, fiber.ErrInternalServerError
	}

	return converter.ContactToResponse(contact), nil
}

func (c *ContactService) Update(ctx context.Context, request *models.UpdateContactRequest) (*models.ContactResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	contact := new(entities.Contact)
	if err := c.ContactRepository.FindByIdAndUsername(tx, contact, request.ID, request.Username); err != nil {
		c.Log.WithError(err).Error("error getting contact")
		return nil, fiber.ErrNotFound
	}

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validation contact")
		return nil, fiber.ErrBadRequest
	}

	contact.FirstName = request.FirstName
	contact.LastName = request.LastName
	contact.Email = request.Email
	contact.Phone = request.Phone

	if err := c.ContactRepository.Update(tx, contact); err != nil {
		c.Log.WithError(err).Error("error updating contact")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error updating contact")
		return nil, fiber.ErrInternalServerError
	}

	return converter.ContactToResponse(contact), nil
}

func (c *ContactService) Delete(ctx context.Context, request *models.DeleteContactRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validation contact")
		return fiber.ErrBadRequest
	}

	contact := new(entities.Contact)
	if err := c.ContactRepository.FindByIdAndUsername(tx, contact, request.ID, request.Username); err != nil {
		c.Log.WithError(err).Error("error getting contact")
		return fiber.ErrNotFound
	}

	if err := c.ContactRepository.Delete(tx, contact); err != nil {
		c.Log.WithError(err).Error("error deleting contact")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error deleting contact")
		return fiber.ErrInternalServerError
	}

	return nil

}

func (c *ContactService) Search(ctx context.Context, request *models.SearchContactRequest) ([]models.ContactResponse, int64, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validation contact")
		return nil, 0, fiber.ErrBadRequest
	}

	contacts, total, err := c.ContactRepository.Search(tx, request)
	if err != nil {
		c.Log.WithError(err).Error("error getting contact")
		return nil, 0, fiber.ErrInternalServerError
	}

	if e := tx.Commit().Error; e != nil {
		c.Log.WithError(e).Error("error getting contact")
		return nil, 0, fiber.ErrInternalServerError
	}

	responses := make([]models.ContactResponse, len(contacts))
	for i, contact := range contacts {
		responses[i] = *converter.ContactToResponse(&contact)
	}

	return responses, total, nil
}
