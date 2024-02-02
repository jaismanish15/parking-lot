package main

import (
	"errors"
)

type ParkingLot struct {
	Capacity int
	Slots    map[int]*Car
	NextSlot int
	IsFull   bool
}

func NewParkingLot(capacity int) (*ParkingLot, error) {
	if capacity <= 0 {
		return nil, errors.New("Invalid number of lots")
	}
	return &ParkingLot{
		Capacity: capacity,
		Slots:    make(map[int]*Car),
		NextSlot: 1,
		IsFull:   false,
	}, nil
}

func (p *ParkingLot) IsCarParked(carToBeChecked *Car) bool {
	for _, car := range p.Slots {
		if car == carToBeChecked {
			return true
		}
	}
	return false
}

func (p *ParkingLot) CountCarsByColor(color Color) int {
	count := 0
	for _, car := range p.Slots {
		if car != nil && car.color == color {
			count++
		}
	}
	return count
}

func (p *ParkingLot) Park(car *Car) (int, error) {
	err := p.checkForSameCarParked(car)
	if err != nil {
		return 0, err
	}
	if p.IsFull {
		return 0, errors.New("Parking lot is full")
	}

	slotNumber := p.NextSlot
	p.Slots[slotNumber] = car
	p.NextSlot++

	if p.NextSlot > p.Capacity {
		p.IsFull = true
	}
	return slotNumber, nil
}

func (p *ParkingLot) checkForSameCarParked(car *Car) error {
	if p.IsCarParked(car) {
		return errors.New("Car is parked already")
	}
	return nil
}
func (pl *ParkingLot) Unpark(slot int, registrationNumber string) (*Car, error) {

	if slot <= 0 || slot > pl.Capacity {
		return nil, errors.New("Invalid slot number")
	}

	car := pl.Slots[slot]
	if car != nil && car.registrationNumber == registrationNumber {
		pl.Slots[slot] = nil
		return car, nil
	}

	return nil, errors.New("Car not found. Thus, cannot be unparked")
}
