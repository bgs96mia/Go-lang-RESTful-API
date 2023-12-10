package converter

import (
	"Golang-RESTful-APi/entities"
	"Golang-RESTful-APi/models"
)

func UserToResponse(user *entities.User) *models.UserResponse {
	return &models.UserResponse{
		Username:  user.Username,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func UserToTokenResponse(user *entities.User) *models.UserResponse {
	return &models.UserResponse{
		Token: user.Token,
	}
}
