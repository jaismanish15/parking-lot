package main

import (
	"errors"
	"regexp"
)

type Car struct {
	registrationNumber string
	color              Color
}

func NewCar(registrationNumber string, color Color) (*Car, error) {
	if err := validateRegistrationNumber(registrationNumber); err != nil {
		return nil, err
	}

	return &Car{registrationNumber: registrationNumber, color: color}, nil
}

func validateRegistrationNumber(registrationNumber string) error {
	registrationNumberRegex := "^[A-Z]{2}\\d{2}[A-Z]{2}\\d{4}"
	matched, err := regexp.MatchString(registrationNumberRegex, registrationNumber)
	if err != nil {
		return err
	}

	if !matched {
		return errors.New("Invalid car registration number")
	}

	return nil
}
