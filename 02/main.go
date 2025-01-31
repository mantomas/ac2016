package main

import (
	"bufio"
	"fmt"
	"os"
)

var keypad1 = [3][3]string{
	{"1", "2", "3"},
	{"4", "5", "6"},
	{"7", "8", "9"},
}

var keypad2 = [5][5]string{
	{"", "", "1", "", ""},
	{"", "2", "3", "4", ""},
	{"5", "6", "7", "8", "9"},
	{"", "A", "B", "C", ""},
	{"", "", "D", "", ""}}

func setupArgparse() string {
	default_file := "input.txt"

	if len(os.Args) < 2 {
		// Default value
		fmt.Println("No file provided, defaulting to input.txt.")
		return default_file
	}
	// Get the first positional argument
	// check if it is a file
	arg := os.Args[1]
	if _, err := os.Stat(arg); os.IsNotExist(err) {
		fmt.Println("File does not exist, defaulting to input.txt.")
		return default_file
	}
	return arg
}

func get_lines(file_path string) []string {
	file, err := os.Open(file_path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func main() {
	file_path := setupArgparse()
	lines := get_lines(file_path)
	// start at 5 of keypad1
	x := 1
	y := 1
	output := ""
	for _, line := range lines {
		for _, c := range line {
			switch string(c) {
			case "U":
				if x > 0 {
					x--
				}
			case "D":
				if x < 2 {
					x++
				}
			case "L":
				if y > 0 {
					y--
				}
			case "R":
				if y < 2 {
					y++
				}
			}
		}
		output += keypad1[x][y]
	}
	fmt.Println("Day 2, part 1 answer: ", output)

	// start at 5 of keypad2
	x2 := 2
	y2 := 0
	output2 := ""
	for _, line := range lines {
		for _, c := range line {
			switch string(c) {
			case "U":
				if x2 > 0 && keypad2[x2-1][y2] != "" {
					x2--
				}
			case "D":
				if x2 < 4 && keypad2[x2+1][y2] != "" {
					x2++
				}
			case "L":
				if y2 > 0 && keypad2[x2][y2-1] != "" {
					y2--
				}
			case "R":
				if y2 < 4 && keypad2[x2][y2+1] != "" {
					y2++
				}
			}
		}
		output2 += keypad2[x2][y2]
	}
	fmt.Println("Day 2, part 2 answer: ", output2)
}
