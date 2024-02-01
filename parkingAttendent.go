package main

import "errors"

type ParkingAttendant struct {
	assignedParkingLots []*ParkingLot
}

func NewParkingAttendant() *ParkingAttendant {
	return &ParkingAttendant{
		assignedParkingLots: make([]*ParkingLot, 0),
	}
}

func (pa *ParkingAttendant) AssignParkingLot(parkingLot *ParkingLot) {
	pa.assignedParkingLots = append(pa.assignedParkingLots, parkingLot)

}

func (pa *ParkingAttendant) Park(car *Car) (int, error) {
	for _, parkingLot := range pa.assignedParkingLots {
		slotNumber, err := parkingLot.Park(car)
		if err == nil {
			return slotNumber, nil
		}
	}

	return 0, errors.New("All assigned parking lots are full")
}
