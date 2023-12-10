package config

import (
	"Golang-RESTful-APi/controllers"
	"Golang-RESTful-APi/middleware"
	"Golang-RESTful-APi/repository"
	"Golang-RESTful-APi/routes"
	"Golang-RESTful-APi/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	userRepository := repository.NewUserRepository(config.Log)
	contactRepository := repository.NewContactRepository(config.Log)
	addressRepository := repository.NewAddressRepository(config.Log)

	userService := services.NewUserService(config.DB, config.Log, config.Validate, userRepository)
	contactService := services.NewContactService(config.DB, config.Log, config.Validate, contactRepository)
	addressService := services.NewAddressService(config.DB, config.Log, config.Validate, addressRepository, contactRepository)

	userController := controllers.NewUserController(userService, config.Log)
	contactController := controllers.NewContactController(contactService, config.Log)
	addressController := controllers.NewAddressController(addressService, config.Log)

	authMiddleware := middleware.NewAuth(userService)

	routeConfig := routes.RouteConfig{
		App:               config.App,
		UserController:    userController,
		ContactController: contactController,
		AddressController: addressController,
		AuthMiddleware:    authMiddleware,
	}
	routeConfig.Setup()
}
