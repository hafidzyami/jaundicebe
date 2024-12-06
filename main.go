package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/hafidzyami/jaundicebe/controller"
	repository "github.com/hafidzyami/jaundicebe/repository/impl"
	service "github.com/hafidzyami/jaundicebe/service/impl"
	"github.com/joho/godotenv"
	supa "github.com/nedpals/supabase-go"
	"github.com/spf13/viper"
	"github.com/gofiber/swagger"
	_ "github.com/hafidzyami/jaundicebe/docs"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Use Viper to read the API_URL and API_KEY from the environment
	viper.AutomaticEnv()

	API_URL := viper.GetString("SUPABASE_URL")
	API_KEY := viper.GetString("SUPABASE_API_KEY")

	// Initialize the Supabase client
	supabase := supa.CreateClient(API_URL, API_KEY)

	// repository
	articleRepository := repository.NewArticleRepository(supabase)
	hospitalRepository := repository.NewHospitalRepository(supabase)
	userRepository := repository.NewUserRepository(supabase)

	// service
	articleService := service.NewArticleService(&articleRepository)
	hospitalService := service.NewHospitalService(&hospitalRepository)
	userService := service.NewUserService(&userRepository)

	// controller
	ArticleController := controller.NewArticleController(&articleService)
	HospitalController := controller.NewHospitalController(&hospitalService)
	UserController := controller.NewUserController(&userService)
	

	// setup fiber
	app := fiber.New()
	
	// routing
	ArticleController.Route(app)
	HospitalController.Route(app)
	UserController.Route(app)

	// Swagger
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	// Listen on port 3000
	err = app.Listen(viper.GetString("SERVER_PORT"))
	if err != nil {
		log.Fatal(err)
	}
}
