package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func setupArgparse() string {
	default_file := "input.txt"

	if len(os.Args) < 2 {
		fmt.Println("No file provided, defaulting to input.txt.")
		return default_file
	}

	arg := os.Args[1]
	if _, err := os.Stat(arg); os.IsNotExist(err) {
		fmt.Println("File does not exist, defaulting to input.txt.")
		return default_file
	}
	return arg
}

func main() {
	file_path := setupArgparse()
	file, err := os.Open(file_path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner1 := bufio.NewScanner(file)
	column1 := []int{}
	column2 := []int{}
	column3 := []int{}
	good_triangles1 := 0
	for scanner1.Scan() {
		line := scanner1.Text()
		sides := strings.Fields(line)
		sides_ints := []int{}
		for i, side := range sides {
			number, _ := strconv.Atoi(side)
			sides_ints = append(sides_ints, number)
			switch i {
			case 0:
				column1 = append(column1, number)
			case 1:
				column2 = append(column2, number)
			case 2:
				column3 = append(column3, number)
			}
		}
		sort.Ints(sides_ints)
		if sides_ints[2] < sides_ints[0]+sides_ints[1] {
			good_triangles1++
		}

	}
	fmt.Println("Day 3, part 1 answer: ", good_triangles1)

	// join the columns
	column1 = append(column1, column2...)
	column1 = append(column1, column3...)
	good_triangles2 := 0
	for i := 0; i < len(column1); i += 3 {
		current := []int{column1[i], column1[i+1], column1[i+2]}
		sort.Ints(current)
		if current[2] < current[0]+current[1] {
			good_triangles2++
		}
	}
	fmt.Println("Day 3, part 2 answer: ", good_triangles2)
}
