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

// WorldMap represents a 2D map with a grid of cells.
type WorldMap struct {
	Size   Point       // Size stores the width and height of the map.
	matrix [][]Element // matrix stores the grid of cells as Elements.
}

// NewWorldMap creates a new WorldMap instance with the specified dimensions.
// It preallocates memory for the map and initializes all cells with ".".
func NewWorldMap(width, height uint32) *WorldMap {
	matrix := make([][]Element, height)
	for i := range matrix {
		matrix[i] = make([]Element, width)
		for j := range matrix[i] {
			matrix[i][j] = *NewElement("background", ".", Point{uint32(j), uint32(i)})
		}
	}
	return &WorldMap{
		Size:   Point{width, height},
		matrix: matrix,
	}
}

// Put sets the value of a specific cell at coordinates (x, y).
// It checks if the coordinates are within bounds and returns an error if not.
func (wm *WorldMap) Put(el Element) error {

	if !el.hasPosition() {
		return fmt.Errorf("element %s has no position", el.Name)
	}

	if !wm.IsValidCoord(*el.Position) {
		return fmt.Errorf("coordinates (%d, %d) are out of bounds", el.Position.X, el.Position.Y)
	}

	wm.matrix[el.Position.Y][el.Position.Y] = el
	return nil
}

// Move moves an element from one cell to another in the specified direction.
// It checks if the movement is valid and returns an error if not.
func (wm *WorldMap) Move(element *Element, direction string) (position *Point, err error) {
	if !element.hasPosition() {
		return nil, fmt.Errorf("element %s has no position", element.Name)
	}

	originalElement := wm.GetElementCopyAt(element.Position.X, element.Position.Y)

	np := Point{element.Position.X, element.Position.Y}

	switch direction {
	case UP:
		np.Y++
	case DOWN:
		np.Y--
	case LEFT:
		np.X--
	case RIGHT:
		np.X++
	default:
		return nil, fmt.Errorf("unknown direction: %s", direction)
	}

	if !wm.IsValidCoord(np) {
		return nil, fmt.Errorf("invalid move to (%d, %d)", np.X, np.Y)
	}

	movedElement := element.SetPosition(np.X, np.Y)

	wm.Put(originalElement)
	wm.Put(*movedElement)

	return movedElement.Position, nil
}

// IsValidCoord checks if the given position is within bounds.
func (wm *WorldMap) IsValidCoord(pos Point) bool {
	return pos.X < uint32(wm.Size.X) && pos.Y < uint32(wm.Size.Y)
}

func (wm *WorldMap) IsMapEdge(p *Point) (isEdge bool, edgeType string) {
	topY := wm.Size.Y - 1
	bottomY := uint32(0)

	leftX := uint32(0)
	rightX := wm.Size.X - 1

	switch {
	case p.X == leftX:
		return true, LEFT
	case p.X == rightX:
		return true, RIGHT
	case p.Y == topY:
		return true, TOP
	case p.Y == bottomY:
		return true, BOTTOM
	default:
		return false, ""

	}
}

func (wm *WorldMap) GetElementCopyAt(x, y uint32) Element {
	return wm.matrix[y][x]
}

func (wm *WorldMap) GetElementAt(x, y uint32) *Element {
	return &wm.matrix[y][x]
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
			builder.WriteString(cell.String())
		}
		builder.WriteString("\n")
	}

	return builder.String()
}

// clearScreen clears the terminal screen using ANSI escape codes.
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
