package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hafidzyami/jaundicebe/middleware"
	"github.com/hafidzyami/jaundicebe/model"
	"github.com/hafidzyami/jaundicebe/service"
)

type ArticleController struct {
	service service.ArticleService
}

func NewArticleController(service *service.ArticleService) *ArticleController {
	return &ArticleController{service: *service}
}

func (controller ArticleController) Route(app *fiber.App) {
	app.Post("/v1/api/article", middleware.JWTMiddleware, controller.Create)
	app.Put("/v1/api/article/:id", middleware.JWTMiddleware, controller.Update)
	app.Delete("/v1/api/article/:id", middleware.JWTMiddleware, controller.Delete)
	app.Get("/v1/api/article", controller.FindAll)
	app.Get("/v1/api/article/:id", controller.FindByID)
}

// @Summary Create Article
// @Description Create a new article
// @Tags Articles
// @Accept json
// @Produce json
// @Param article body model.ArticleCreateOrUpdate true "Article Data"
// @Param Authorization header string true "Authorization" default(Bearer <JWT_TOKEN>)
// @Success 200 {object} model.GeneralResponse
// @Failure 400 {object} map[string]string
// @Security JWT
// @Router /v1/api/article [post]
func (controller ArticleController) Create(c *fiber.Ctx) error {
	var request model.ArticleCreateOrUpdate
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

// @Summary Update Article
// @Description Update an existing article
// @Tags Articles
// @Accept json
// @Produce json
// @Param id path int true "Article ID"
// @Param article body model.ArticleCreateOrUpdate true "Updated Article Data"
// @Param Authorization header string true "Authorization" default(Bearer <JWT_TOKEN>)
// @Success 200 {object} model.GeneralResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Security JWT
// @Router /v1/api/article/{id} [put]
func (controller ArticleController) Update(c *fiber.Ctx) error {
	var request model.ArticleCreateOrUpdate
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

// @Summary Delete Article
// @Description Delete an article by ID
// @Tags Articles
// @Produce json
// @Param id path int true "Article ID"
// @Param Authorization header string true "Authorization" default(Bearer <JWT_TOKEN>)
// @Success 200 {object} model.GeneralResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Security JWT
// @Router /v1/api/article/{id} [delete]
func (controller ArticleController) Delete(c *fiber.Ctx) error {
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

// @Summary Get Article by ID
// @Description Get article details by its ID
// @Tags Articles
// @Accept json
// @Produce json
// @Param id path int true "Article ID"
// @Success 200 {object} model.GeneralResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /v1/api/article/{id} [get]
func (controller ArticleController) FindByID(c *fiber.Ctx) error {
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

// @Summary Get All Articles
// @Description Get a list of all articles
// @Tags Articles
// @Produce json
// @Success 200 {object} model.GeneralResponse
// @Failure 400 {object} map[string]string
// @Router /v1/api/article [get]
func (controller ArticleController) FindAll(c *fiber.Ctx) error {
	response := controller.service.FindAll(c.Context())
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    fiber.StatusOK,
		Message: "Success",
		Data:    response,
	})
}
