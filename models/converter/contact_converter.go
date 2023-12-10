package converter

import (
	"Golang-RESTful-APi/entities"
	"Golang-RESTful-APi/models"
)

func ContactToResponse(contact *entities.Contact) *models.ContactResponse {
	return &models.ContactResponse{
		ID:        contact.ID,
		FirstName: contact.FirstName,
		LastName:  contact.LastName,
		Email:     contact.Email,
		Phone:     contact.Phone,
		CreatedAt: contact.CreatedAt,
		UpdatedAt: contact.UpdatedAt,
	}
}
