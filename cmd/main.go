package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"weather-city-api/api"
	"weather-city-api/repository"
	"weather-city-api/services"
	"weather-city-api/util"
)

func main() {
	if err := godotenv.Load("../app.env"); err != nil {
		log.Println("Warning: .env file not found")
	}

	cfg, err := util.LoadConfig("..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	fmt.Printf("This is the api key in main func : %s\n", cfg.WeatherAPIKey)

	cityRepo := repository.NewCityRepository()

	cityService := services.NewCityService(cityRepo)
	weatherService := services.NewWeatherService(cfg.WeatherAPIKey)

	router := gin.Default()

	api.SetupRoutes(router, cityService, weatherService)

	log.Printf("Server starting on port %s\n", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
