package main

import "errors"

type ParkingStrategy interface {
	GetNextSlotAvailable(parkingLot *ParkingLot) (int, error)
}

type NearestSlotFirstStrategy struct{}

func (nsf *NearestSlotFirstStrategy) GetNextSlotAvailable(parkingLot *ParkingLot) (int, error) {
	for slot := 1; slot <= parkingLot.Capacity; slot++ {
		if parkingLot.Slots[slot] == nil {
			return slot, nil
		}
	}
	return -1, errors.New("No available slot found")
}

type FarthestSlotFirstStrategy struct{}

func (fsf *FarthestSlotFirstStrategy) GetNextSlotAvailable(parkingLot *ParkingLot) (int, error) {
	for slot := parkingLot.Capacity; slot > 0; slot-- {
		if parkingLot.Slots[slot] == nil {
			return slot, nil
		}
	}
	return -1, errors.New("No available slot found")
}
