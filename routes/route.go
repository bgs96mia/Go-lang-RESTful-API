package routes

import (
	"Golang-RESTful-APi/controllers"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App               *fiber.App
	UserController    *controllers.UserController
	ContactController *controllers.ContactController
	AddressController *controllers.AddressController
	AuthMiddleware    fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	router := c.App.Group("/api")
	router.Post("/users", c.UserController.Register)
	router.Post("/users/_login", c.UserController.Login)
}

func (c *RouteConfig) SetupAuthRoute() {
	router := c.App.Group("/api")
	router.Use(c.AuthMiddleware)
	router.Get("users/_current", c.UserController.Current)
	router.Delete("/users", c.UserController.Logout)
	router.Patch("/users/_current", c.UserController.Update)

	router.Get("/contacts", c.ContactController.List)
	router.Post("/contacts", c.ContactController.Create)
	router.Get("/contacts/:contactId", c.ContactController.Get)
	router.Put("/contacts/:contactId", c.ContactController.Update)
	router.Delete("/contacts/:contactId", c.ContactController.Delete)

	router.Get("/contacts/:contactId/addresses", c.AddressController.List)
	router.Post("/contacts/:contactId/addresses", c.AddressController.Create)
	router.Put("/contacts/:contactId/addresses/:addressId", c.AddressController.Update)
	router.Get("/contacts/:contactId/addresses/:addressId", c.AddressController.Get)
	router.Delete("/contacts/:contactId/addresses/:addressId", c.AddressController.Delete)

}
