package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"weather-city-api/services"
)

type WeatherHandler struct {
	weatherService services.WeatherService
}

func NewWeatherHandler(weatherService services.WeatherService) *WeatherHandler {
	return &WeatherHandler{
		weatherService: weatherService,
	}
}

func (h *WeatherHandler) GetWeather(c *gin.Context) {
	city := c.Query("city")
	if city == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "City parameter is required"})
		return
	}

	weatherData, err := h.weatherService.GetWeatherByCity(city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, weatherData)
}
