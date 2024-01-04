package tests

import (
	"github.com/stretchr/testify/assert"
	"golang_elevator/elevator"
	"testing"
)

// A More intricate test mimicking the pushing of various floor buttons
//	when the elevator is at the mid-stage of floors.

func TestUpThenDown(t *testing.T) {
	elev := elevator.NewElevator(NumberOfFloors)

	travelTheseFloors := []int{5, 3, 2}
	expectTheseFloors := []int{2, 3, 5}
	testUpThenDown(t, &elev, travelTheseFloors, expectTheseFloors)

	// Should be on floor 5 now.
	//	Now we should go all the way up and then come down
	travelTheseFloors = []int{7, 3, 9, 4}
	// Finish going up, then go back down.
	expectTheseFloors = []int{7, 9, 4, 3}
	testUpThenDown(t, &elev, travelTheseFloors, expectTheseFloors)
}

func testUpThenDown(t *testing.T, elev *elevator.Elevator, travelTheseFloors []int,
	expectTheseFloors []int) {

	var err error
	// Add the floors
	for i := 0; i < len(travelTheseFloors); i++ {
		err = elev.AddFloor(travelTheseFloors[i])
		assert.NoError(t, err)
	}

	nextFloor := -1
	// Validate we are travelling to the proper floors.
	for i := 0; i < len(expectTheseFloors); i++ {
		nextFloor, err = elev.NextFloor()
		assert.Equal(t, expectTheseFloors[i], nextFloor)
		assert.NoError(t, err)
	}
}
