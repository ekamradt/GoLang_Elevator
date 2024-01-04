package elevator

import (
	"errors"
	"fmt"
	"github.com/wangjia184/sortedset"
)

// Wrap a Sorted Set within a bounded number of floors and
//
//	a specific Direction (Up, Down)
type FloorList struct {
	floors           *sortedset.SortedSet
	number_of_floors int
	direction        Direction
}

// Move to the next floor in the Set depending on Direction.
//
//	Throw an error if no next floor.
func (f FloorList) NextFloor() (int, error) {
	if f.floors.GetCount() == 0 {
		return 0, errors.New("No next floor.")
	}
	nextFloor := 0
	if f.direction == Up {
		nextFloor = f.floors.PopMin().Value.(int)
	} else {
		nextFloor = f.floors.PopMax().Value.(int)
	}
	return nextFloor, nil
}

// Return the number of floors in the Queue (Set)
func (f FloorList) GetCount() int {
	return f.floors.GetCount()
}

// Add a new Fllor to this Queue (Set)
func (f FloorList) AddFloor(floor int) error {
	if floor < 1 {
		return errors.New(fmt.Sprintf("Floors below 1 are not impleted (floor %d)", floor))
	}
	if floor > f.number_of_floors {
		return errors.New(fmt.Sprintf("Floor limit is %d, can't reach floor %d", f.number_of_floors, floor))
	}
	key := fmt.Sprintf("%06d", floor)
	score := sortedset.SCORE(floor)
	value := int(floor)
	f.floors.AddOrUpdate(key, score, value)
	return nil
}

func NewFloorList(numFloors int, direction Direction) FloorList {
	return FloorList{
		floors:           sortedset.New(),
		number_of_floors: numFloors,
		direction:        direction,
	}
}
