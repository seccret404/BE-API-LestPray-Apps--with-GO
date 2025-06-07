package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/seccret404/BE-API-LestPray-Apps--with-GO/models"
)

func GetWorshipService(lat, lon float64) ([]models.WorshipPlace, error){
	apiKey := os.Getenv("GEOFY_API")
	url := fmt.Sprintf(
		"https://api.geoapify.com/v2/places?categories=religion.place_of_worship&filter=circle:%f,%f,5000&bias=proximity:%f,%f&limit=20&apiKey=%s",
		lon, lat, lon, lat, apiKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("geory error: %s", string(body))

	}

	var geoResp models.GeoapifyResponse
	if err :=  json.NewDecoder(resp.Body).Decode(&geoResp); err != nil{
		return nil, err
	}

	var results []models.WorshipPlace

	for _, f := range geoResp.Features{
		results = append(results, models.WorshipPlace{
			Name: f.Properties.Name,
			Category: f.Properties.Category,
			Religion: f.Properties.Address.Religion,
			Lat: f.Properties.Lat,
			Lon: f.Properties.Lon,
		})
	}

	return results, nil



}