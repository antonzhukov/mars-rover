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
	x         int64
	y         int64
	direction string
	obstacles []coordinate
}

type coordinate struct {
	x int64
	y int64
}

func NewRover(x int64, y int64, direction string, obstacles []coordinate) *Rover {
	return &Rover{x: x, y: y, direction: direction, obstacles: obstacles}
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
		fmt.Println(r.x, r.y, r.direction)
	}

	var errMsg string
	if err != nil {
		errMsg = fmt.Sprintf(" %s", err.Error())
	}
	return fmt.Sprintf("(%d, %d) %s%s", r.x, r.y, r.direction, errMsg)
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
	x := r.x
	y := r.y
	switch r.direction {
	case "WEST":
		x -= int64(direction)
	case "EAST":
		x += int64(direction)
	case "NORTH":
		y += int64(direction)
	case "SOUTH":
		y -= int64(direction)
	}
	for _, o := range r.obstacles {
		if o.x == x && o.y == y {
			return errors.New("STOPPED")
		}
	}
	r.x = x
	r.y = y
	return nil
}
