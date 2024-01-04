package tests

import (
	"github.com/stretchr/testify/assert"
	"golang_elevator/elevator"
	"testing"
)

const NumberOfFloors = 100
const TestFloor_01 = 70
const TestFloor_02 = 33

// This set of tests just tests basic Floor series movement

// Test NextFloor when there are no more floors.
func TestEmptyFloors(t *testing.T) {
	f := elevator.NewFloorList(NumberOfFloors, elevator.Up)
	floorCount := f.GetCount()
	_, err := f.NextFloor()

	assert.Error(t, err)
	assert.Equal(t, 0, floorCount)
}

// Test Bad input on AddFloor - Lower Range
func TestAddFloorLessThanOne(t *testing.T) {
	f := elevator.NewFloorList(NumberOfFloors, elevator.Up)
	err := f.AddFloor(0)
	floorCount := f.GetCount()
	assert.Error(t, err)
	assert.Equal(t, 0, floorCount)
}

// Test Bad input on AddFloor - Upper Range
func TestAddFloorHigherThanBuilding(t *testing.T) {
	f := elevator.NewFloorList(NumberOfFloors, elevator.Up)
	err := f.AddFloor(NumberOfFloors + 1)
	floorCount := f.GetCount()
	assert.Error(t, err)
	assert.Equal(t, 0, floorCount)
}

// Test the movement to a single floor
func TestOneFloor(t *testing.T) {
	f := elevator.NewFloorList(NumberOfFloors, elevator.Up)

	// Add a floor number, check we have just 1 fllor to go to,
	//	make sure we go to that floor
	f.AddFloor(TestFloor_01)

	floorCount := f.GetCount()
	assert.Equal(t, 1, floorCount)

	nextFloor, err := f.NextFloor()
	assert.NoError(t, err)
	assert.Equal(t, TestFloor_01, nextFloor)
}

// Test 2 floors going up, and check that floors are not duplicated.
//
//	Someone can pound and pound on a button, and the floor will only be called once.
func TestUpForNoDuplicateFloors(t *testing.T) {
	floorList := elevator.NewFloorList(NumberOfFloors, elevator.Up)
	testUpOrDownForNoDuplicateFloors(t, &floorList, []int{TestFloor_02, TestFloor_01})
}

// Test 2 floors going up, and check that floors are not duplicated.
//
//	Someone can pound and pound on a button, and the floor will only be called once.
func TestDownForNoDuplicateFloors(t *testing.T) {
	floorList := elevator.NewFloorList(NumberOfFloors, elevator.Down)
	testUpOrDownForNoDuplicateFloors(t, &floorList, []int{TestFloor_01, TestFloor_02})
}

func testUpOrDownForNoDuplicateFloors(t *testing.T, f *elevator.FloorList, expectedFloors []int) {

	for i := 0; i < 5; i++ {
		f.AddFloor(TestFloor_01)
		f.AddFloor(TestFloor_02)
	}

	floorCount := f.GetCount()
	assert.Equal(t, 2, floorCount)

	nextFloor, err := f.NextFloor()
	assert.NoError(t, err)
	assert.Equal(t, expectedFloors[0], nextFloor)

	floorCount = f.GetCount()
	assert.Equal(t, 1, floorCount)

	nextFloor, err = f.NextFloor()
	assert.NoError(t, err)
	assert.Equal(t, expectedFloors[1], nextFloor)

	floorCount = f.GetCount()
	assert.Equal(t, 0, floorCount)
}
