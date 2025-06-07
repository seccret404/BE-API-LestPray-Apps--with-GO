package models


type GeoapifyResponse struct {
	Features []GeoapifyFeature `json:"features"`
}

type GeoapifyFeature struct {
	Properties GeoapifyProperties `json:"properties"`
}

type GeoapifyProperties struct {
	Name     string          `json:"name"`
	Category string          `json:"category"`
	Address  GeoapifyAddress `json:"address"`
	Lat      float64         `json:"lat"`
	Lon      float64         `json:"lon"`
}

type GeoapifyAddress struct {
	Religion string `json:"religion"`
}

//struktur response ke view

type WorshipPlace struct {
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Religion string  `json:"religion"`
	Lat      float64 `json:"lat"`
	Lon      float64 `json:"lon"`
}