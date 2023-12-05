package worldmap

type Point struct {
	X uint32
	Y uint32
}

type Element struct {
	Name           string
	Representation string
	Position       *Point
}

func NewElement(name, representation string, position Point) *Element {
	return &Element{
		Name:           name,
		Representation: representation,
		Position:       &position,
	}
}

func (el *Element) SetPosition(x, y uint32) *Element {
	el.Position.X = x
	el.Position.Y = y
	return el
}

func (el *Element) hasPosition() bool {
	if el.Position == nil {
		return false
	}
	return true
}

func (el *Element) String() string {
	return el.Representation
}
