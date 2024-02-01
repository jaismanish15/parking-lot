package main

import (
	"errors"
	"regexp"
)

type Car struct {
	RegistrationNumber string
	Color              Color
}

func NewCar(registrationNumber string, color Color) (*Car, error) {
	registrationNumberRegex := "^[A-Z]{2}\\d{2}[A-Z]{2}\\d{4}"
	matched, err := regexp.MatchString(registrationNumberRegex, registrationNumber)
	if err != nil {
		return nil, err
	}

	if !matched {
		return nil, errors.New("Invalid car registration number")
	}

	return &Car{RegistrationNumber: registrationNumber, Color: color}, nil
}
