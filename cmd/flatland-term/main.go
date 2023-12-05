package main

import (
	"fmt"
	worldmap "golang/game-v1/internal"
	"math/rand"
	"time"
)

func getRandomDirection(exclusion string) (direction string, err error) {

	for {
		randomIndex := rand.Intn(4)
		switch randomIndex {
		case 0:
			direction = worldmap.UP
		case 1:
			direction = worldmap.DOWN
		case 2:
			direction = worldmap.LEFT
		case 3:
			direction = worldmap.RIGHT
		default:
			return "", fmt.Errorf("unknown direction index %d", randomIndex)
		}

		fmt.Printf("exclusion[%s] index[%d] direction[%s]", exclusion, randomIndex, direction)
		// run this puppy until it comes back with a direction that isn't the exclusion
		// Hacky as hell, but is simple for now
		if direction != exclusion {
			return direction, nil
		}
	}
}

func main() {
	// Create a new WorldMap instance with dimensions 40x10
	worldMap := worldmap.NewWorldMap(200, 60)

	ox, y := 0, 5
	worldMap.Put(ox, y, "*")

	// Move an element across the map
	var err error
	direction := worldmap.RIGHT
	var isEdge bool
	var edgeType string
	for x := ox; x < worldMap.Width; {

		isEdge, edgeType = worldMap.IsMapEdge(x, y)

		if isEdge {
			switch edgeType {
			case worldmap.TOP:
				direction, err = getRandomDirection(worldmap.UP)
			case worldmap.BOTTOM:
				direction, err = getRandomDirection(worldmap.DOWN)
			case worldmap.LEFT:
				direction, err = getRandomDirection(worldmap.LEFT)
			case worldmap.RIGHT:
				direction, err = getRandomDirection(worldmap.RIGHT)
			}
			if err != nil {
				fmt.Println(err)
				return
			}

			getRandomDirection(worldmap.LEFT) // test
		}

		x, y, err = worldMap.Move(x, y, direction)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}

		worldMap.PrintMap()
		time.Sleep(50 * time.Millisecond)

	}

	fmt.Println("Movement finished.")
}

func NewWorldMap(i1, i2 int) {
	panic("unimplemented")
}
