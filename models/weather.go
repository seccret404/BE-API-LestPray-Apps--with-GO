package models

import "fmt"

type WeatherResponse struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

type Location struct {
	Name      string  `json:"name"`
	Region    string  `json:"region"`
	Country   string  `json:"country"`
	Lat       float64 `json:"lat"`
	Lon       float64 `json:"lon"`
	TzID      string  `json:"tz_id"`
	Localtime string  `json:"localtime"`
}

type Current struct {
	TempC     float64   `json:"temp_c"`
	Condition Condition `json:"condition"`
	Humidity  int       `json:"humidity"`
	WindKph   float64   `json:"wind_kph"`
}

type Condition struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
}
type WeatherFormatted struct {
	Location    string `json:"location"`
	LocalTime   string `json:"localtime"`
	Temperature string `json:"temperature"`
	Condition   string `json:"condition"`
	Icon        string `json:"icon"`
	Humidity    string `json:"humidity"`
	Wind        string `json:"wind"`
}

func FormatWeatherResponse(w *WeatherResponse) *WeatherFormatted {
	return &WeatherFormatted{
		Location:    fmt.Sprintf("%s, %s", w.Location.Name, w.Location.Country),
		LocalTime:   w.Location.Localtime,
		Temperature: fmt.Sprintf("%.1f°C", w.Current.TempC),
		Condition:   w.Current.Condition.Text,
		Icon:        w.Current.Condition.Icon,
		Humidity:    fmt.Sprintf("%d%%", w.Current.Humidity),
		Wind:        fmt.Sprintf("%.1f kph", w.Current.WindKph),
	}
}
