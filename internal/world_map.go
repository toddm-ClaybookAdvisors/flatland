package worldmap

import (
	"fmt"
	"strings"
)

// Constants for directions
const (
	UP    = "U"
	DOWN  = "D"
	LEFT  = "L"
	RIGHT = "R"
)

// constants for map borders
const (
	TOP    = "T"
	BOTTOM = "B"
)

type Position struct {
}

// WorldMap represents a 2D map with a grid of cells.
type WorldMap struct {
	Width  int        // Width is the number of columns in the map.
	Height int        // Height is the number of rows in the map.
	matrix [][]string // matrix stores the grid of cells as strings.
}

// NewWorldMap creates a new WorldMap instance with the specified dimensions.
// It preallocates memory for the map and initializes all cells with ".".
func NewWorldMap(width, height int) *WorldMap {
	matrix := make([][]string, height)
	for i := range matrix {
		matrix[i] = make([]string, width)
		for j := range matrix[i] {
			matrix[i][j] = "."
		}
	}
	return &WorldMap{
		Width:  width,
		Height: height,
		matrix: matrix,
	}
}

// Put sets the value of a specific cell at coordinates (x, y).
// It checks if the coordinates are within bounds and returns an error if not.
func (wm *WorldMap) Put(x, y int, element string) error {
	if !wm.IsValidCoord(x, y) {
		return fmt.Errorf("coordinates (%d, %d) are out of bounds", x, y)
	}

	wm.matrix[y][x] = element
	return nil
}

// Move moves an element from one cell to another in the specified direction.
// It checks if the movement is valid and returns an error if not.
func (wm *WorldMap) Move(x, y int, direction string) (newX, newY int, err error) {
	newX, newY = x, y

	switch direction {
	case UP:
		newY++
	case DOWN:
		newY--
	case LEFT:
		newX--
	case RIGHT:
		newX++
	default:
		return -1, -1, fmt.Errorf("unknown direction: %s", direction)
	}

	if !wm.IsValidCoord(newX, newY) {
		return -1, -1, fmt.Errorf("invalid move to (%d, %d)", newX, newY)
	}

	oldElement := wm.matrix[newY][newX]
	currentElement := wm.matrix[y][x]

	wm.Put(newX, newY, currentElement)
	wm.Put(x, y, oldElement)

	return newX, newY, nil
}

// IsValidCoord checks if the given coordinates are within bounds.
func (wm *WorldMap) IsValidCoord(x, y int) bool {
	return x >= 0 && x < wm.Width && y >= 0 && y < wm.Height
}

func (wm *WorldMap) IsMapEdge(x, y int) (isEdge bool, edgeType string) {
	topY := wm.Height - 1
	bottomY := 0

	leftX := 0
	rightX := wm.Width - 1

	switch {
	case x == leftX:
		return true, LEFT
	case x == rightX:
		return true, RIGHT
	case y == topY:
		return true, TOP
	case y == bottomY:
		return true, BOTTOM
	default:
		return false, ""

	}
}

// PrintMap prints the current state of the WorldMap with element movement.
func (wm *WorldMap) PrintMap() {
	clearScreen()
	fmt.Print(wm.String())
}

// String returns a string representation of the WorldMap.
func (wm *WorldMap) String() string {
	var builder strings.Builder

	for _, row := range wm.matrix {
		for _, cell := range row {
			builder.WriteString(cell)
		}
		builder.WriteString("\n")
	}

	return builder.String()
}

// clearScreen clears the terminal screen using ANSI escape codes.
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
