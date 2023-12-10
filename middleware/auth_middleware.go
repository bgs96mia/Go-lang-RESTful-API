package middleware

import (
	"Golang-RESTful-APi/models"
	"Golang-RESTful-APi/services"
	"github.com/gofiber/fiber/v2"
)

func NewAuth(userService *services.UserService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		request := &models.VerifyUserRequest{Token: ctx.Get("Authorization", "NOT_FOUND")}
		userService.Log.Debugf("Authorization : %s", request.Token)

		auth, err := userService.Verify(ctx.UserContext(), request)
		if err != nil {
			userService.Log.Warnf("Failed find user by token : %+v", err)
			return fiber.ErrUnauthorized
		}

		userService.Log.Debugf("User : %+v", auth.Username)
		ctx.Locals("auth", auth)
		return ctx.Next()
	}
}

func GetUser(ctx *fiber.Ctx) *models.Auth {
	return ctx.Locals("auth").(*models.Auth)
}
