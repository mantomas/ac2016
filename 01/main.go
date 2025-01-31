package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Location struct {
	X int
	Y int
}

func get_facing(facing string, direction string) string {
	if direction == "R" {
		switch facing {
		case "N":
			return "E"
		case "E":
			return "S"
		case "S":
			return "W"
		case "W":
			return "N"
		}
	} else {
		switch facing {
		case "N":
			return "W"
		case "W":
			return "S"
		case "S":
			return "E"
		case "E":
			return "N"
		}
	}
	return facing
}

func find_crossing(locations []Location) *Location {
	visited := []Location{}
	prev := Location{0, 0}          // starting point
	visited = append(visited, prev) // Add starting point to visited

	// Go through all vertices until intersection found
	for _, curr := range locations {
		// for each move exclude starting point
		if curr.X == prev.X {
			y1, y2 := prev.Y, curr.Y
			if y1 > y2 {
				// negative y direction
				for y := y1 - 1; y >= y2; y-- {
					if is_in(visited, Location{curr.X, y}) {
						return &Location{curr.X, y}
					}
					visited = append(visited, Location{curr.X, y})
				}
			} else {
				for y := y1 + 1; y <= y2; y++ {
					if is_in(visited, Location{curr.X, y}) {
						return &Location{curr.X, y}
					}
					visited = append(visited, Location{curr.X, y})
				}
			}
		} else {
			x1, x2 := prev.X, curr.X
			if x1 > x2 {
				// negative x direction
				for x := x1 - 1; x >= x2; x-- {
					if is_in(visited, Location{x, curr.Y}) {
						return &Location{x, curr.Y}
					}
					visited = append(visited, Location{x, curr.Y})
				}
			} else {
				for x := x1 + 1; x <= x2; x++ {
					if is_in(visited, Location{x, curr.Y}) {
						return &Location{x, curr.Y}
					}
					visited = append(visited, Location{x, curr.Y})
				}
			}
		}

		prev = curr // Move to next vertex
	}
	return nil // No intersection

}

func is_in(locations []Location, location Location) bool {
	for _, loc := range locations {
		if loc == location {
			return true
		}
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func decide_location(X *int, Y *int, facing string, distance int) Location {
	switch facing {
	case "N":
		*Y += distance
	case "S":
		*Y -= distance
	case "E":
		*X += distance
	case "W":
		*X -= distance
	}
	return Location{*X, *Y}
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	text := string(content)
	parts := strings.Split(text, ",")
	X, Y := 0, 0
	facing := "N"
	locations := []Location{}
	for _, part := range parts {
		part = strings.TrimSpace(part)
		turn := string(string(part[0]))
		distance, _ := strconv.Atoi(string(part[1:]))
		facing = get_facing(facing, turn)
		new_location := decide_location(&X, &Y, facing, distance)
		// store locations for part two
		locations = append(locations, new_location)

	}
	// part one result from the end location
	fmt.Println("Day 1, part 1 answer: ", abs(X)+abs(Y))
	// walk the path again to find the first crossing
	right_spot := find_crossing(locations)
	fmt.Println("Day 1, part 2 answer: ", abs(right_spot.X)+abs(right_spot.Y))
}
