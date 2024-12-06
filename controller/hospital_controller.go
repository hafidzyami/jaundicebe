package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hafidzyami/jaundicebe/middleware"
	"github.com/hafidzyami/jaundicebe/model"
	"github.com/hafidzyami/jaundicebe/service"
)

type HospitalController struct {
	service service.HospitalService
}

func NewHospitalController(service *service.HospitalService) *HospitalController {
	return &HospitalController{service: *service}
}

func (controller HospitalController) Route(app *fiber.App) {
	app.Post("/v1/api/hospital", middleware.JWTMiddleware, controller.Create)
	app.Put("/v1/api/hospital/:id", middleware.JWTMiddleware,controller.Update)
	app.Delete("/v1/api/hospital/:id", middleware.JWTMiddleware,controller.Delete)
	app.Get("/v1/api/hospital", controller.FindAll)
	app.Get("/v1/api/hospital/:id", controller.FindByID)
}

// @Summary Create hospital
// @Description Create a new hospital
// @Tags Hospitals
// @Accept json
// @Produce json
// @Param hospital body model.HospitalCreateOrUpdate true "hospital Data"
// @Param Authorization header string true "Authorization" default(Bearer <JWT_TOKEN>)
// @Success 200 {object} model.GeneralResponse
// @Failure 400 {object} map[string]string
// @Security JWT
// @Router /v1/api/hospital [post]
func (controller HospitalController) Create(c *fiber.Ctx) error {
	var request model.HospitalCreateOrUpdate
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	response := controller.service.Create(c.Context(), request)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    response,
	})
}

// @Summary Update hospital
// @Description Update an existing hospital
// @Tags Hospitals
// @Accept json
// @Produce json
// @Param id path int true "hospital ID"
// @Param hospital body model.HospitalCreateOrUpdate true "Updated hospital Data"
// @Param Authorization header string true "Authorization" default(Bearer <JWT_TOKEN>)
// @Success 200 {object} model.GeneralResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Security JWT
// @Router /v1/api/hospital/{id} [put]
func (controller HospitalController) Update(c *fiber.Ctx) error {
	var request model.HospitalCreateOrUpdate
	id := c.Params("id")
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	response := controller.service.Update(c.Context(), request, idInt)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    response,
	})
}

// @Summary Delete hospital
// @Description Delete an hospital by ID
// @Tags Hospitals
// @Produce json
// @Param id path int true "hospital ID"
// @Param Authorization header string true "Authorization" default(Bearer <JWT_TOKEN>)
// @Success 200 {object} model.GeneralResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Security JWT
// @Router /v1/api/hospital/{id} [delete]
func (controller HospitalController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	controller.service.Delete(c.Context(), idInt)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "Success",
	})
}

// @Summary Get hospital by ID
// @Description Get hospital details by its ID
// @Tags Hospitals
// @Accept json
// @Produce json
// @Param id path int true "hospital ID"
// @Success 200 {object} model.GeneralResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /v1/api/hospital/{id} [get]
func (controller HospitalController) FindByID(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	response, err := controller.service.FindByID(c.Context(), idInt)
	if err != nil {
		// Check if the error is a "not found" error and return 404
		return c.Status(fiber.StatusNotFound).JSON(model.GeneralResponse{
			Code:    fiber.StatusNotFound,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    response,
	})
}

// @Summary Get All Hospitals
// @Description Get a list of all Hospitals
// @Tags Hospitals
// @Produce json
// @Success 200 {object} model.GeneralResponse
// @Failure 400 {object} map[string]string
// @Router /v1/api/hospital [get]
func (controller HospitalController) FindAll(c *fiber.Ctx) error {
	response := controller.service.FindAll(c.Context())
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    response,
	})
}
