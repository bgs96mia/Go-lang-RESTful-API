package converter

import (
	"Golang-RESTful-APi/entities"
	"Golang-RESTful-APi/models"
)

func AddressToResponse(address *entities.Address) *models.AddressResponse {
	return &models.AddressResponse{
		ID:         address.ID,
		Street:     address.Street,
		City:       address.City,
		Province:   address.Province,
		PostalCode: address.PostalCode,
		Country:    address.Country,
		CreatedAt:  address.CreatedAt,
		UpdatedAt:  address.UpdatedAt,
	}
}
