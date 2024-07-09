package utils

import "github.com/gofiber/fiber/v2"

type ValidHTTPResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (data ValidHTTPResponse) Send(ctx *fiber.Ctx) error {
	return ctx.Status(data.Status).JSON(data)
}

type ErrorHTTPResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors,omitempty"`
}

func (data ErrorHTTPResponse) Send(ctx *fiber.Ctx) error {
	return ctx.Status(data.Status).JSON(data)
}
