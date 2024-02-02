package main

import (
	"testing"
)

func TestParkingLotIsCreated(t *testing.T) {
	parkingLot, err := NewParkingLot(12)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if parkingLot == nil {
		t.Error("Parking should be created, got nil")
	}
}

func TestParkingLotWithZeroThrowsException(t *testing.T) {
	_, err := NewParkingLot(0)
	if err == nil {
		t.Error("Expected an error for zero lots, got nil")
	}
}

func TestParkingLotWithNegative2ThrowsException(t *testing.T) {
	_, err := NewParkingLot(-2)
	if err == nil {
		t.Error("Expected an error for -2 lots, got nil")
	}
}

func TestACarIsNotPresentInLot(t *testing.T) {
	parkingLot := &ParkingLot{Capacity: 1}
	car, err := NewCar("AB12BC1235", RED)
	if err != nil {
		t.Fatalf("Error creating car: %v", err)
	}
	actual := parkingLot.IsCarParked(car)

	if actual {
		t.Error("Expected the car not to be present in the parking lot")
	}
}

func TestParkCarInFullCapacityParkingLot(t *testing.T) {
	parkingLot, err := NewParkingLot(1)
	if err != nil {
		t.Fatalf("Error creating parking lot: %v", err)
	}

	car1, err := NewCar("AB12BC1236", BLUE)
	if err != nil {
		t.Fatalf("Error creating car: %v", err)
	}

	_, err = parkingLot.Park(car1)
	if err != nil {
		t.Fatalf("Error parking car: %v", err)
	}

	car2, err := NewCar("AB12BC1234", BLACK)
	if err != nil {
		t.Fatalf("Error creating car: %v", err)
	}

	_, err = parkingLot.Park(car2)
	if err == nil {
		t.Error("Expected an error, but parking was successful.")
	} else if err.Error() != "Parking lot is full" {
		t.Errorf("Expected 'Parking lot is full' error, got: %v", err)
	}
}

func TestParkCar(t *testing.T) {
	parkingLot, err := NewParkingLot(12)
	if err != nil {
		t.Fatalf("Error creating lot: %v", err)
	}
	car, err := NewCar("AB12CD3456", RED)
	if err != nil {
		t.Fatalf("Error creating car: %v", err)
	}

	slotNumber, err := parkingLot.Park(car)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if slotNumber <= 0 || slotNumber > parkingLot.Capacity {
		t.Errorf("Invalid slot number: %d", slotNumber)
	}

	parkedCar := parkingLot.Slots[slotNumber]
	if parkedCar == nil || parkedCar != car {
		t.Errorf("Car is not parked correctly")
	}
}

func TestUnparkCar(t *testing.T) {
	parkingLot, err := NewParkingLot(2)
	if err != nil {
		t.Fatalf("Error creating parking lot: %v", err)
	}
	// parking
	car1, err := NewCar("AB12BC1236", BLUE)
	if err != nil {
		t.Fatalf("Error creating car: %v", err)
	}
	slotNumber, err := parkingLot.Park(car1)
	if err != nil {
		t.Fatalf("Error parking car: %v", err)
	}

	// unparking same car
	_, err = parkingLot.Unpark(slotNumber, "AB12BC1236")
	if err != nil {
		t.Errorf("Unexpected error when unparking car: %v", err)
	}

	if parkingLot.Slots[slotNumber] != nil {
		t.Error("Parking slot should be empty after unparking.")
	}
}

func TestUnparkCarWithInvalidSlotNumber(t *testing.T) {
	parkingLot, err := NewParkingLot(2)
	if err != nil {
		t.Fatalf("Error creating parking lot: %v", err)
	}

	_, err = parkingLot.Unpark(0, "AB12BC1236")

	if err == nil {
		t.Error("Expected an error for invalid slot number, but got nil.")
	} else if err.Error() != "Invalid slot number" {
		t.Errorf("'Invalid slot number' error, got: %v", err)
	}
}

func TestParkSameCarTwiceWithoutUnparking(t *testing.T) {
	parkingLot, err := NewParkingLot(3)
	if err != nil {
		t.Fatalf("Error creating parking lot: %v", err)
	}

	car, err := NewCar("AB12BC1236", BLUE)
	if err != nil {
		t.Fatalf("Error creating car: %v", err)
	}

	_, err = parkingLot.Park(car)
	if err != nil {
		t.Fatalf("Error parking car: %v", err)
	}

	_, err = parkingLot.Park(car)

	if err.Error() != "Car is parked already" {
		t.Errorf("Expected 'Car is parked already' error, got: %v", err)
	}
}
