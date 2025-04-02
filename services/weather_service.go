package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WeatherService interface {
	GetWeatherByCity(city string) (map[string]interface{}, error)
}

type weatherService struct {
	apiKey string
}

func NewWeatherService(apiKey string) WeatherService {
	return &weatherService{
		apiKey: apiKey,
	}
}

func (s *weatherService) GetWeatherByCity(city string) (map[string]interface{}, error) {
	fmt.Printf("This is the city : %s\n", city)
	fmt.Printf("This is the api key: %s\n", s.apiKey)

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, s.apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("weather API returned status code %d", resp.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}
