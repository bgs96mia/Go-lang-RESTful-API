package controllers

import (
	"Golang-RESTful-APi/middleware"
	"Golang-RESTful-APi/models"
	"Golang-RESTful-APi/services"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	Log         *logrus.Logger
	UserService *services.UserService
}

func NewUserController(userService *services.UserService, Log *logrus.Logger) *UserController {
	return &UserController{
		Log:         Log,
		UserService: userService,
	}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	request := new(models.RegisterUserRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, e := c.UserService.Create(ctx.UserContext(), request)
	if e != nil {
		c.Log.Warnf("Failed to register user : %+v", e)
		return e
	}

	return ctx.JSON(models.WebResponse[*models.UserResponse]{Data: response})
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	request := new(models.LoginUserRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, e := c.UserService.Login(ctx.UserContext(), request)
	if e != nil {
		c.Log.Warnf("Failed to login user : %+v", e)
		return e
	}

	return ctx.JSON(models.WebResponse[*models.UserResponse]{Data: response})
}

func (c *UserController) Current(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := &models.GetUserRequest{
		Username: auth.Username,
	}

	response, e := c.UserService.Current(ctx.UserContext(), request)
	if e != nil {
		c.Log.WithError(e).Warnf("Failed to logout user")
		return e
	}

	return ctx.JSON(models.WebResponse[*models.UserResponse]{Data: response})

}

func (c *UserController) Logout(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := &models.LogoutUserRequest{
		Username: auth.Username,
	}

	response, err := c.UserService.Logout(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to logout user")
		return err
	}

	return ctx.JSON(models.WebResponse[bool]{Data: response})
}

func (c *UserController) Update(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(models.UpdateUserRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	request.Username = auth.Username
	response, err := c.UserService.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Warnf("Failed to update user")
		return err
	}

	return ctx.JSON(models.WebResponse[*models.UserResponse]{Data: response})
}
