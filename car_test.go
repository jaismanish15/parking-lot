package main

import (
	"testing"
)

func TestCarIsCreated(t *testing.T) {
	car, err := NewCar("TN12AB1234", BLACK)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if car == nil {
		t.Error("Expected a car to be created, got nil")
	}
}

func TestInvalidRegistrationCarThrowsException(t *testing.T) {
	_, err := NewCar("12CD", RED)
	if err == nil {
		t.Error("Expected an error for invalid registration, got nil")
	}
}
