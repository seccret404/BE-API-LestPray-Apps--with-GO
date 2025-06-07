package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/seccret404/BE-API-LestPray-Apps--with-GO/models"
)

func GetWeatherService(lat, lon float64) (*models.WeatherResponse, error) {
	apiKey := os.Getenv("WEATHER_API")
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%f,%f", apiKey, lat, lon)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to get weather data, status: %d", resp.StatusCode)
	}

	var weatherResp models.WeatherResponse
	
	err = json.NewDecoder(resp.Body).Decode(&weatherResp)
	if err != nil {
		return nil, err
	}

	return &weatherResp, nil
}