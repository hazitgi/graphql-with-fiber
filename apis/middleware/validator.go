package middleware

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/hazitgi/loom-erp/apis/utils"
)

type validationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Middleware for validating incoming requests Fiber
func ValidateRequest(data interface{}, customValidationMessages map[string]string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		errors := make([]validationError, 0)
		if err := ctx.BodyParser(&data); err != nil {
			response := utils.ErrorHTTPResponse{
				Status:  fiber.StatusForbidden,
				Message: "Invalid request payload",
				Errors:  err.Error(),
			}
			return response.Send(ctx)
		}
		validate := validator.New()
		err := validate.Struct(data)

		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, validationErr := range validationErrors {
				field := validationErr.Field()
				tag := validationErr.Tag()
				key := field + "." + tag

				customMessage := customValidationMessages[key]
				if customMessage == "" {
					customMessage = validationErr.Error()
				}

				errors = append(errors, validationError{
					Field:   field,
					Message: customMessage,
				})
			}
			response := utils.ErrorHTTPResponse{
				Status:  fiber.StatusUnprocessableEntity,
				Message: "Validation errors",
				Errors:  errors,
			}
			return response.Send(ctx)
		} else {
			fmt.Println("Validation errors: ", errors)
			ctx.Locals("input", data)
			return ctx.Next()
		}
	}
}
