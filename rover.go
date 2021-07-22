package mars_rover

import (
	"errors"
	"fmt"
)

var directions = []string{"NORTH", "EAST", "SOUTH", "WEST"}

type rotate int

const (
	rotateLeft  rotate = 1
	rotateRight rotate = -1
)

type move int64

const (
	moveForward  move = 1
	moveBackward move = -1
)

type Rover struct {
	position  coordinate
	direction string
	obstacles []coordinate
}

type coordinate struct {
	x int64
	y int64
}

func NewRover(x int64, y int64, direction string, obstacles []coordinate) *Rover {
	return &Rover{position: coordinate{x: x, y: y}, direction: direction, obstacles: obstacles}
}

func (r *Rover) Command(cc string) string {
	var err error
	for i := 0; i < len(cc); i++ {
		c := string(cc[i])
		if c == "F" {
			err = r.move(moveForward)
			if err != nil {
				break
			}
		} else if c == "B" {
			err = r.move(moveBackward)
			if err != nil {
				break
			}
		} else if c == "L" {
			r.direction = r.rotate(rotateLeft)
		} else if c == "R" {
			r.direction = r.rotate(rotateRight)
		}
	}

	var errMsg string
	if err != nil {
		errMsg = fmt.Sprintf(" %s", err.Error())
	}
	return fmt.Sprintf("(%d, %d) %s%s", r.position.x, r.position.y, r.direction, errMsg)
}

func (r *Rover) rotate(direction rotate) string {
	for i, d := range directions {
		if d == r.direction {
			return directions[(i-int(direction)+len(directions))%len(directions)]
		}
	}
	return ""
}

func (r *Rover) move(direction move) error {
	nextPos := r.position
	switch r.direction {
	case "WEST":
		nextPos.x -= int64(direction)
	case "EAST":
		nextPos.x += int64(direction)
	case "NORTH":
		nextPos.y += int64(direction)
	case "SOUTH":
		nextPos.y -= int64(direction)
	}
	for _, o := range r.obstacles {
		if o.x == nextPos.x && o.y == nextPos.y {
			return errors.New("STOPPED")
		}
	}
	r.position = nextPos
	return nil
}
