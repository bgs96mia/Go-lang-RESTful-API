package controllers

import (
	"Golang-RESTful-APi/middleware"
	"Golang-RESTful-APi/models"
	"Golang-RESTful-APi/services"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type AddressController struct {
	Log            *logrus.Logger
	AddressService *services.AddressService
}

func NewAddressController(AddressService *services.AddressService, log *logrus.Logger) *AddressController {
	return &AddressController{
		Log:            log,
		AddressService: AddressService,
	}
}

func (c *AddressController) Create(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(models.CreateAddressRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return fiber.ErrBadRequest
	}

	request.Username = auth.Username
	request.ContactId = ctx.Params("contactId")

	response, err := c.AddressService.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to create address")
		return err
	}

	return ctx.JSON(models.WebResponse[*models.AddressResponse]{Data: response})
}

func (c *AddressController) List(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)
	contactId := ctx.Params("contactId")

	request := &models.ListAddressRequest{
		Username:  auth.Username,
		ContactId: contactId,
	}

	responses, err := c.AddressService.List(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to list addresses")
		return err
	}

	return ctx.JSON(models.WebResponse[[]models.AddressResponse]{Data: responses})
}

func (c *AddressController) Get(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)
	contactId := ctx.Params("contactId")
	addressId := ctx.Params("addressId")

	request := &models.GetAddressRequest{
		Username:  auth.Username,
		ContactId: contactId,
		ID:        addressId,
	}

	response, err := c.AddressService.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to get address")
		return err
	}

	return ctx.JSON(models.WebResponse[*models.AddressResponse]{Data: response})
}

func (c *AddressController) Update(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(models.UpdateAddressRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return fiber.ErrBadRequest
	}

	request.Username = auth.Username
	request.ContactId = ctx.Params("contactId")
	request.ID = ctx.Params("addressId")

	response, err := c.AddressService.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("failed to update address")
		return err
	}

	return ctx.JSON(models.WebResponse[*models.AddressResponse]{Data: response})
}

func (c *AddressController) Delete(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)
	contactId := ctx.Params("contactId")
	addressId := ctx.Params("addressId")

	request := &models.DeleteAddressRequest{
		Username:  auth.Username,
		ContactId: contactId,
		ID:        addressId,
	}

	if err := c.AddressService.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Error("failed to delete address")
		return err
	}

	return ctx.JSON(models.WebResponse[bool]{Data: true})
}
