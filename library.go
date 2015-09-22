package main

type Library struct {
	ID          uint   `json:"id"`
	MappingID   string `json:"-"`
	Name        string `json:"name"`
	TakenPlaces uint16 `json:"takenPlaces,omitempty"`
	TotalPlaces uint16 `json:"totalPlaces,omitempty"`
}

type Libraries []Library
