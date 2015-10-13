package main

type Library struct {
	ID          uint    `json:"id"`
	MappingID   string  `json:"-"`
	Name        string  `json:"name"`
	TakenPlaces uint16  `json:"takenPlaces,omitempty"`
	TotalPlaces uint16  `json:"totalPlaces,omitempty"`
	Distance    float32 `json:"distance,omitempty"`
	Latitude    float32 `json:"latitude,omitempty"`
	Longitude   float32 `json:"longitude,omitempty"`
	Description string  `json:"description,omitempty"`
	City        string  `json:"city,omitempty"`
	CountryCode string  `json:"countryCode,omitempty"`
	Contact     string  `json:"contact,omitempty"`
}

type Libraries []Library
