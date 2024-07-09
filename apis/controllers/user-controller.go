package controllers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/hazitgi/graphql-with-fiber/apis/common"
	"github.com/hazitgi/graphql-with-fiber/apis/services"
	"github.com/hazitgi/graphql-with-fiber/apis/utils"
	"github.com/hazitgi/graphql-with-fiber/models"
)

type UserController interface {
	GetUser(ctx *fiber.Ctx) error
	CreateUser(ctx *fiber.Ctx) error
	ListUser(ctx *fiber.Ctx) error
	DeleteUser(ctx *fiber.Ctx) error
	UpdateUser(ctx *fiber.Ctx) error
}

type userControllerImpl struct {
	// Implement user logic here
	userDatabaseService *services.UserService
}

func NewUserControllerImpl() UserController {
	userService := services.NewUserService()
	return &userControllerImpl{
		userDatabaseService: userService,
	}
}

func (userCtrl userControllerImpl) GetUser(ctx *fiber.Ctx) error {
	log.Printf("Getting user...")
	id := ctx.Params("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("invalid user id: %s", id)
		response := utils.ErrorHTTPResponse{
			Status:  fiber.StatusNotFound,
			Message: "User not found",
			Errors:  fmt.Sprintf("invalid user id: %s", id),
		}
		return response.Send(ctx)
	}
	user, err := userCtrl.userDatabaseService.GetUserByID(uint(userId))
	if err != nil {
		fmt.Printf("failed to get user: %v", err)
		response := utils.ErrorHTTPResponse{
			Status:  fiber.StatusNotFound,
			Message: "User not found",
			Errors:  fmt.Sprintf("failed to get user: %v", err),
		}
		return response.Send(ctx)
	}

	response := utils.ValidHTTPResponse{
		Status:  fiber.StatusOK,
		Message: "User retrieved successfully",
		Data:    user,
	}
	return response.Send(ctx)
}

func (userCtrl userControllerImpl) DeleteUser(ctx *fiber.Ctx) error {
	log.Printf("Deleting  user...")
	id := ctx.Params("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("invalid user id: %s", id)
		response := utils.ErrorHTTPResponse{
			Status:  fiber.StatusNotFound,
			Message: "User not found",
			Errors:  fmt.Sprintf("invalid user id: %s", id),
		}
		return response.Send(ctx)
	}
	err = userCtrl.userDatabaseService.DeleteUser(uint(userId))
	if err != nil {
		fmt.Printf("failed to Delete user: %v", err)
		response := utils.ErrorHTTPResponse{
			Status:  fiber.StatusNotFound,
			Message: "User not found",
			Errors:  fmt.Sprintf("failed to get user: %v", err),
		}
		return response.Send(ctx)
	}

	response := utils.ValidHTTPResponse{
		Status:  fiber.StatusOK,
		Message: "User deleted successfully",
	}
	return response.Send(ctx)
}

func (userCtrl userControllerImpl) ListUser(ctx *fiber.Ctx) error {
	log.Printf("Getting all users...")
	pagination := &common.Pagination{}
	if err := ctx.QueryParser(pagination); err != nil {
		log.Fatalf("pagination error: %v", err)
	}
	pagination.Setup()

	// Implement pagination
	_, err := userCtrl.userDatabaseService.FindAllUsers(pagination)
	if err != nil {
		fmt.Printf("failed to get user: %v", err)
		response := utils.ErrorHTTPResponse{
			Status:  fiber.StatusNotFound,
			Message: "User not found",
			Errors:  fmt.Sprintf("failed to get user: %v", err),
		}
		return response.Send(ctx)
	}
	response := utils.ValidHTTPResponse{
		Status:  fiber.StatusOK,
		Message: "User retrieved successfully",
		Data:    pagination,
	}
	return response.Send(ctx)
}

func (userCtrl *userControllerImpl) CreateUser(ctx *fiber.Ctx) error {
	log.Printf("Creating user...")
	userInput := ctx.Locals("input").(*common.CreateUserInput)
	// check is user already exists
	_, err := userCtrl.userDatabaseService.FindUserByField("email", userInput.Email)
	if err == nil {
		fmt.Printf("user with email '%s' already exists", userInput.Email)
		response := utils.ErrorHTTPResponse{
			Status:  fiber.StatusBadRequest,
			Message: "User already exists",
			Errors:  fmt.Sprintf("user with email '%s' already exists", userInput.Email),
		}
		return response.Send(ctx)
	}

	hashPassword, err := utils.HashPassword(userInput.Password)
	if err != nil {
		log.Printf("failed to hash password: %v", err)
		response := utils.ErrorHTTPResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Internal Server Error",
			Errors:  err.Error(), // Include error details for debugging purposes, remove in production code.
		}
		return response.Send(ctx)
	}
	userInput.Password = hashPassword
	user := models.User{
		FullName:    userInput.FullName,
		CompanyName: userInput.CompanyName,
		CountryID:   userInput.CountryID,
		StateID:     userInput.StateID,
		Email:       userInput.Email,
		Location:    userInput.Location,
		Address:     userInput.Address,
		Password:    userInput.Password,
	}
	insertedUser, err := userCtrl.userDatabaseService.InsertUser(&user)
	if err != nil {
		response := utils.ErrorHTTPResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Internal Server Error",
			Errors:  err.Error(), // Include error details for debugging purposes, remove in production code.
		}
		return response.Send(ctx)
	}
	response := utils.ValidHTTPResponse{
		Status:  fiber.StatusCreated,
		Message: "User created successfully",
		Data:    insertedUser,
	}
	return response.Send(ctx)
}

func (userCtrl *userControllerImpl) UpdateUser(ctx *fiber.Ctx) error {
	log.Printf("Creating user...")
	id := ctx.Params("id")
	userId, err := strconv.Atoi(id)
	fmt.Println(userId)
	userInput := ctx.Locals("input").(*common.CreateUserInput)
	// check is user already exists
	_, err = userCtrl.userDatabaseService.FindUserByField("email", userInput.Email)
	if err == nil {
		fmt.Printf("user with email '%s' already exists", userInput.Email)
		response := utils.ErrorHTTPResponse{
			Status:  fiber.StatusBadRequest,
			Message: "User already exists",
			Errors:  fmt.Sprintf("user with email '%s' already exists", userInput.Email),
		}
		return response.Send(ctx)
	}

	hashPassword, err := utils.HashPassword(userInput.Password)
	if err != nil {
		log.Printf("failed to hash password: %v", err)
		response := utils.ErrorHTTPResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Internal Server Error",
			Errors:  err.Error(), // Include error details for debugging purposes, remove in production code.
		}
		return response.Send(ctx)
	}
	userInput.Password = hashPassword

	updatedUser, err := userCtrl.userDatabaseService.UpdateUser(uint(userId), *userInput)
	if err != nil {
		response := utils.ErrorHTTPResponse{
			Status:  fiber.StatusInternalServerError,
			Message: "Internal Server Error",
			Errors:  err.Error(), // Include error details for debugging purposes, remove in production code.
		}
		return response.Send(ctx)
	}
	response := utils.ValidHTTPResponse{
		Status:  fiber.StatusCreated,
		Message: "User updated successfully",
		Data:    updatedUser,
	}
	return response.Send(ctx)
}
