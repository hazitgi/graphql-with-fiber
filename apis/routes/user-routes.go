package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/hazitgi/loom-erp/apis/common"
	"github.com/hazitgi/loom-erp/apis/controllers"
	"github.com/hazitgi/loom-erp/apis/middleware"
)

type UserRoutes struct {
	apiGroup       string
	userController controllers.UserController
	validate       *validator.Validate
}

func NewUserRoutes() *UserRoutes {
	controller := controllers.NewUserControllerImpl()
	return &UserRoutes{
		"api/users",
		controller,
		validator.New(),
	}
}

func (routes *UserRoutes) RegisterUserRoutes(app *fiber.App) {
	api := app.Group(routes.apiGroup)

	api.Get("/", routes.userController.ListUser)
	api.Get("/:id", routes.userController.GetUser)
	api.Delete("/:id", routes.userController.DeleteUser)
	api.Post("/", middleware.ValidateRequest(new(common.CreateUserInput), common.UserValidationMessages), routes.userController.CreateUser)
	api.Put("/:id", middleware.ValidateRequest(new(common.CreateUserInput), common.UserValidationMessages), routes.userController.UpdateUser)
}
