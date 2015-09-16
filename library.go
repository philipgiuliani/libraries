package main

type Library struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	TakenPlaces uint16 `json:"takenPlaces"`
	TotalPlaces uint16 `json:"totalPlaces"`
}

type Libraries []Library
