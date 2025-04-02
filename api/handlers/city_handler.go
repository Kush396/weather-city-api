package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"weather-city-api/models"
	"weather-city-api/services"
)

type CityHandler struct {
	cityService services.CityService
}

func NewCityHandler(cityService services.CityService) *CityHandler {
	return &CityHandler{
		cityService: cityService,
	}
}

func (h *CityHandler) GetAllCities(c *gin.Context) {
	cities := h.cityService.GetAllCities()
	c.JSON(http.StatusOK, cities)
}

type getCityByIDRequest struct {
	ID int `uri:"id" binding:"required,min=1"`
}

func (h *CityHandler) GetCityByID(c *gin.Context) {
	var req getCityByIDRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid city ID"})
		return
	}

	fmt.Printf("This is the id: %d\n", req.ID)

	city, err := h.cityService.GetCityByID(req.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "City not found"})
		return
	}

	c.JSON(http.StatusOK, city)
}

func (h *CityHandler) CreateCity(c *gin.Context) {
	var arg models.CityRequest
	if err := c.ShouldBindJSON(&arg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	city := models.City{
		ID:          1,
		Name:        arg.Name,
		Country:     arg.Country,
		Description: arg.Description,
	}

	createdCity, err := h.cityService.CreateCity(city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdCity)
}

func (h *CityHandler) UpdateCity(c *gin.Context) {
	var req getCityByIDRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid city ID"})
		return
	}

	var arg models.CityRequest
	if err := c.ShouldBindJSON(&arg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	city := models.City{
		ID:          req.ID,
		Name:        arg.Name,
		Country:     arg.Country,
		Description: arg.Description,
	}

	updatedCity, err := h.cityService.UpdateCity(req.ID, city)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedCity)
}

func (h *CityHandler) DeleteCity(c *gin.Context) {
	var req getCityByIDRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid city ID"})
		return
	}

	if err := h.cityService.DeleteCity(req.ID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "City deleted successfully"})
}
