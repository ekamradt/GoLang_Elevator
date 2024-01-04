package tests

import (
	"github.com/stretchr/testify/assert"
	direction "golang_elevator/elevator"
	"testing"
)

// Test the simple toggle between Up and Down.
func TestDirectionToggle(t *testing.T) {
	dir := direction.Up
	oppositeDir := dir.Toggle()
	assert.Equal(t, dir, direction.Up)
	assert.Equal(t, oppositeDir, direction.Down)
}
