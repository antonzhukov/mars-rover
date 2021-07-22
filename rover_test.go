package mars_rover

import "testing"

func TestRover_Command(t *testing.T) {
	tests := []struct {
		name      string
		x         int64
		y         int64
		direction string
		command   string
		obstacles []coordinate
		want      string
	}{
		{
			"one move forward",
			4,
			2,
			"EAST",
			"F",
			nil,
			"(5, 2) EAST",
		},
		{
			"one move backward",
			4,
			2,
			"EAST",
			"B",
			nil,
			"(3, 2) EAST",
		},
		{
			"turn left",
			4,
			2,
			"EAST",
			"L",
			nil,
			"(4, 2) NORTH",
		},
		{
			"turn right",
			4,
			2,
			"EAST",
			"R",
			nil,
			"(4, 2) SOUTH",
		},
		{
			"sequence of moves",
			4,
			2,
			"EAST",
			"FLFFFRFLB",
			nil,
			"(6, 4) NORTH",
		},
		{
			"obstacle met",
			6,
			4,
			"NORTH",
			"FLFFFRFLB",
			[]coordinate{{1, 4}, {3, 5}, {7, 4}},
			"(4, 5) WEST STOPPED",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRover(tt.x, tt.y, tt.direction, tt.obstacles)
			got := r.Command(tt.command)
			if tt.want != got {
				t.Errorf("Command failed, expected %s, got %s", tt.want, got)
			}
		})
	}
}

func TestRover_rotate(t *testing.T) {
	tests := []struct {
		name           string
		roverDirection string
		direction      rotate
		want           string
	}{
		{
			"turn left",
			"NORTH",
			rotateLeft,
			"WEST",
		},
		{
			"turn right",
			"NORTH",
			rotateRight,
			"EAST",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rover{
				direction: tt.roverDirection,
			}
			if got := r.rotate(tt.direction); got != tt.want {
				t.Errorf("rotate() = %v, want %v", got, tt.want)
			}
		})
	}
}
