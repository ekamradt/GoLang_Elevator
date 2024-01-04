package elevator

type Direction int

const (
	Up Direction = iota
	Down
)

func (d Direction) Toggle() Direction {
	if d == Up {
		return Down
	} else {
		return Up
	}
}

func (d Direction) String() string {
	switch d {
	case 0:
		return "Up"
	case 1:
		return "Down"
	}
	return "Sideways"
}
