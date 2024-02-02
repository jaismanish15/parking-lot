package main

import (
	"testing"
)

func TestCreateParkingAttendant(t *testing.T) {
	parkingAttendant := NewParkingAttendant()

	if parkingAttendant == nil {
		t.Error("Failed to create a parking attendant")
	}
}

func TestUnassignedAttendantCannotParkCar(t *testing.T) {
	unassignedAttendant := NewParkingAttendant()

	car := &Car{
		registrationNumber: "AB12BC1234",
		color:              RED,
	}

	_, err := unassignedAttendant.Park(car)

	expectedErrorMsg := "All assigned parking lots are full"
	if err == nil {
		t.Error("Expected an error for an unassigned attendant parking a car, but got nil.")
	} else if err.Error() != expectedErrorMsg {
		t.Errorf("Expected '%s' error, got: %v", expectedErrorMsg, err)
	}
}

func TestParkingAttendantCanParkCar(t *testing.T) {
	parkingAttendant := NewParkingAttendant()
	parkingLot, _ := NewParkingLot(5)
	parkingAttendant.AssignParkingLot(parkingLot)

	car := &Car{
		registrationNumber: "AB12BC1234",
		color:              RED,
	}

	slotNumber, err := parkingAttendant.Park(car)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if slotNumber <= 0 || slotNumber > parkingLot.Capacity {
		t.Errorf("Invalid slot number: %d", slotNumber)
	}
}

func TestParkingAttendantCanUnparkCar(t *testing.T) {
	ParkingAttendant := NewParkingAttendant()
	ParkingLot, _ := NewParkingLot(5)
	ParkingAttendant.AssignParkingLot(ParkingLot)

	car := &Car{
		registrationNumber: "AB12BC1234",
		color:              RED,
	}

	slotNumber, err := ParkingAttendant.Park(car)
	if err != nil {
		t.Fatalf("Error parking car: %v", err)
	}

	unparkedCar, err := ParkingAttendant.Unpark(slotNumber, "AB12BC1234")
	if err != nil {
		t.Fatalf("Error unparking car: %v", err)
	}

	if unparkedCar != car {
		t.Error("Unparked car does not match the parked car.")
	}
}
