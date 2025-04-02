package api

import (
	"github.com/gin-gonic/gin"

	"weather-city-api/api/handlers"
	"weather-city-api/services"
)

func SetupRoutes(router *gin.Engine, cityService services.CityService, weatherService services.WeatherService) {
	cityHandler := handlers.NewCityHandler(cityService)
	weatherHandler := handlers.NewWeatherHandler(weatherService)

	router.GET("/api/weather", weatherHandler.GetWeather)
	router.GET("/api/cities", cityHandler.GetAllCities)
	router.GET("/api/cities/:id", cityHandler.GetCityByID)
	router.POST("/api/cities", cityHandler.CreateCity)
	router.PUT("/api/cities/:id", cityHandler.UpdateCity)
	router.DELETE("/api/cities/:id", cityHandler.DeleteCity)
}
