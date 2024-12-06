package controller

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/hafidzyami/jaundicebe/middleware"
	"github.com/hafidzyami/jaundicebe/model"
	"github.com/hafidzyami/jaundicebe/service"
	"github.com/hafidzyami/jaundicebe/utils"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService *service.UserService) UserController {
	return UserController{UserService: *userService}
}

func (controller UserController) Route(app *fiber.App) {
	app.Post("/v1/auth/register", controller.Create)
	app.Post("/v1/auth/login", controller.Login)
	app.Put("/v1/auth/change-password", middleware.JWTMiddleware, controller.ChangePassword)
}

// @Summary Create User
// @Description Create a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body model.UserCreateOrUpdate true "User Data"
// @Success 200 {object} model.GeneralResponse
// @Failure 400 {object} map[string]string
// @Router /v1/auth/register [post]
func (controller UserController) Create(c *fiber.Ctx) error {
	var request model.UserCreateOrUpdate
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(400).JSON(model.GeneralResponse{
			Code:    400,
			Message: err.Error(),
			Data:    nil,
		})
	}
	response, err := controller.UserService.Create(c.Context(), request)
	if err != nil {
		// Handle specific errors
		if err.Error() == "username already exists" {
			return c.Status(fiber.StatusConflict).JSON(model.GeneralResponse{
				Code:    fiber.StatusConflict,
				Message: "Username already exists",
				Data:    nil,
			})
		}

		// Handle other unexpected errors
		return c.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    fiber.StatusInternalServerError,
			Message: "Internal server error",
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    response,
	})
}

// @Summary Login User
// @Description Login a user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body model.UserCreateOrUpdate true "User Data"
// @Success 200 {object} model.GeneralResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /v1/auth/login [post]
func (controller UserController) Login(c *fiber.Ctx) error {
	var request model.UserCreateOrUpdate
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(400).JSON(model.GeneralResponse{
			Code:    400,
			Message: err.Error(),
			Data:    nil,
		})
	}
	response, err := controller.UserService.Login(c.Context(), request)
	if err != nil {
		// Handle specific errors
		if err.Error() == "invalid credentials" {
			return c.Status(fiber.StatusUnauthorized).JSON(model.GeneralResponse{
				Code:    fiber.StatusUnauthorized,
				Message: "Invalid credentials",
				Data:    nil,
			})
		}

		// Handle other unexpected errors
		return c.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    fiber.StatusInternalServerError,
			Message: "Internal server error",
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    response,
	})
}

// @Summary Change Password
// @Description Change user's password
// @Tags Users
// @Accept json
// @Produce json
// @Param user body model.ChangePassword true "User Data"
// @Param Authorization header string true "Authorization" default
// @Success 200 {object} model.GeneralResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Security JWT
// @Router /v1/auth/change-password [put]
func (controller UserController) ChangePassword(c *fiber.Ctx) error {
	// Get the Authorization header
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(model.GeneralResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "Authorization header missing",
			Data:    nil,
		})
	}

	// Extract the token by splitting the "Bearer " prefix
	token := strings.TrimPrefix(authHeader, "Bearer ")

	// Get user ID from the token
	userId, err := utils.GetUserIDFromToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(model.GeneralResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "Unauthorized",
			Data:    nil,
		})
	}

	// Parse the body to get the new password details
	var request model.ChangePassword
	err = c.BodyParser(&request)
	if err != nil {
		return c.Status(400).JSON(model.GeneralResponse{
			Code:    400,
			Message: err.Error(),
			Data:    nil,
		})
	}

	// Call service to change the password
	response, err := controller.UserService.ChangePassword(c.Context(), userId, request)
	if err != nil {
		// Handle specific errors
		if err.Error() == "invalid credentials" {
			return c.Status(fiber.StatusUnauthorized).JSON(model.GeneralResponse{
				Code:    fiber.StatusUnauthorized,
				Message: "Invalid credentials",
				Data:    nil,
			})
		}

		if err.Error() == "invalid old password" {
			return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
				Code:    fiber.StatusBadRequest,
				Message: "Invalid old password",
				Data:    nil,
			})
		}

		// Handle other unexpected errors
		return c.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Code:    fiber.StatusInternalServerError,
			Message: "Internal server error",
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    response,
	})
}
