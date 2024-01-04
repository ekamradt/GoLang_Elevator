package elevator

import (
	"errors"
	"fmt"
)

type Elevator struct {
	floors       int
	direction    Direction
	currentFloor int
	upQueue      FloorList
	downQueue    FloorList
}

func NewElevator(numFloors int) Elevator {
	return Elevator{
		floors:       numFloors,
		direction:    Up,
		currentFloor: 1,
		upQueue:      NewFloorList(numFloors, Up),
		downQueue:    NewFloorList(numFloors, Down),
	}
}

func (e Elevator) CurrentFloor() int {
	return e.currentFloor
}

func (e Elevator) Direction() Direction {
	return e.direction
}

func (e Elevator) AddFloor(floor int) error {
	if floor < 1 {
		return errors.New(fmt.Sprintf("Floors below 1 are not impleted (floor %d)", floor))
	}
	if floor > e.floors {
		return errors.New(fmt.Sprintf("Floor limit is %d, can't reach floor %d", e.floors, floor))
	}
	if floor == e.currentFloor {
		return errors.New(fmt.Sprintf("Already on floor %d", floor))
	}

	if floor < e.currentFloor {
		e.downQueue.AddFloor(floor)
		return nil
	}
	e.upQueue.AddFloor(floor)
	return nil
}

func (e *Elevator) NextFloor() (int, error) {
	err := errors.New("")
	if e.direction == Up {
		if e.upQueue.GetCount() > 0 {
			e.currentFloor, err = e.upQueue.NextFloor()
			return e.currentFloor, err
		}
		e.direction = Down
	}

	if e.downQueue.GetCount() > 0 {
		e.currentFloor, err = e.downQueue.NextFloor()
		return e.currentFloor, err
	} else {
		e.direction = Up
		if e.downQueue.GetCount() == 0 {
			return -1, errors.New("No next floor, try pushing a button.")
		}
	}
	return e.NextFloor()
}
