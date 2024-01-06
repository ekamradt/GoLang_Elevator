# GoLang Elevator

### Introduction

The past 10 or so years I've been mostly a Java micro-serivces person. 
I do have a  bit of experience with GoLang. 
I decided to put that exerience to the test write a Interview Question in Go.

I was given this problem in two different interviews. 

### Problem
* Write an algorithm for an elevator.

### Technical Summary

All functionality is implemented within tests.

* The basic rules for this elevator is.
  * Finish in one direction, up or down, before switching directions.
  * At any point in time someone can call the elevator for any floor (that exists).

* To solve this I created an Elevator struct (class) that contains two Sorted Sets (SS).
    * One Sorted Set for each of the directions Up and Down.
    * We start at floor one (1) and set the Elevator's direction to Up.
    * When someone presses a button to call the elevator we add that floor to either the Up SS
        if the called floor is greater than the current floor, or else we add it to the Down SS.
        If the floor called is the current floor, we ignore the event.
    * We move the elevator up performing a Pop Min, moving from lowest to highest, from the Up SS until the Set is empty.
    * When the Up Set is empty we change the elevator;s direction to Down, and start 
        travelling to the Down SS by doing a Pop Max, and going from highest to lowest.
    * When the Down SS is empty we reverse the process and start travelling Up again.

### Tests

See tests/ directory and run via 
```azure
go test -v ./...
```

### Test Code
```azure
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

```