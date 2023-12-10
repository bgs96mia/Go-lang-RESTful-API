package controllers

import (
	"Golang-RESTful-APi/middleware"
	"Golang-RESTful-APi/models"
	"Golang-RESTful-APi/services"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"math"
)

type ContactController struct {
	Log            *logrus.Logger
	ContactService *services.ContactService
}

func NewContactController(contactService *services.ContactService, Log *logrus.Logger) *ContactController {
	return &ContactController{
		Log:            Log,
		ContactService: contactService,
	}
}

func (c *ContactController) Create(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(models.CreateContactRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("error parsing request body")
		return fiber.ErrBadRequest
	}

	request.Username = auth.Username

	response, err := c.ContactService.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error creating contact")
		return err
	}

	return ctx.JSON(models.WebResponse[*models.ContactResponse]{Data: response})
}

func (c *ContactController) Get(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := &models.GetContactRequest{
		Username: auth.Username,
		ID:       ctx.Params("contactId"),
	}

	response, err := c.ContactService.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error creating contact")
		return err
	}

	return ctx.JSON(models.WebResponse[*models.ContactResponse]{Data: response})
}

func (c *ContactController) Update(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := new(models.UpdateContactRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("error parsing request body")
		return fiber.ErrBadRequest
	}

	request.Username = auth.Username
	request.ID = ctx.Params("contactId")

	response, err := c.ContactService.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error updating contact")
		return err
	}

	return ctx.JSON(models.WebResponse[*models.ContactResponse]{Data: response})
}

func (c *ContactController) Delete(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	contactId := ctx.Params("contactId")

	request := &models.DeleteContactRequest{
		Username: auth.Username,
		ID:       contactId,
	}

	if err := c.ContactService.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Error("error deleting contact")
		return err
	}

	return ctx.JSON(models.WebResponse[bool]{Data: true})
}

func (c *ContactController) List(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)

	request := &models.SearchContactRequest{
		Username: auth.Username,
		Name:     ctx.Query("name", ""),
		Phone:    ctx.Query("phone", ""),
		Email:    ctx.Query("email", ""),
		Page:     ctx.QueryInt("page", 1),
		Size:     ctx.QueryInt("size", 10),
	}

	responses, total, err := c.ContactService.Search(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error searching contact")
		return err
	}

	paging := &models.PageMetadata{
		Page:      request.Page,
		Size:      request.Size,
		TotalItem: total,
		TotalPage: int64(math.Ceil(float64(total) / float64(request.Size))),
	}

	return ctx.JSON(models.WebResponse[[]models.ContactResponse]{Data: responses, Paging: paging})
}
