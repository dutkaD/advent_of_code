package day11

import (
	"advent_of_code/utils"
	"fmt"
	"strings"
)

func GetFlashes(lines []string) int {
	var octopuses [][]int
	for _, line := range lines {
		octopuses = append(octopuses, utils.StringsToInts(strings.Split(line, "")))
	}
	step := 0
	sum := 0
	fmt.Println(octopuses)

	allFlash := false
	for !allFlash {
		step++

		for r, row := range octopuses {
			for c, _ := range row {
				octopuses[r][c] += 1
			}
		}

		for r, row := range octopuses {
			for c, _ := range row {
				if octopuses[r][c] >= 10 {
					octopuses = checkAdjacent(octopuses, r, c)
				}
			}
		}

		count := 0
		for r, row := range octopuses {
			for c, _ := range row {
				if octopuses[r][c] == 0 {
					sum += 1
					count += 1
				}
			}
		}

		if count == len(octopuses) * len(octopuses[0]) {
			fmt.Println(step)
			allFlash = true
		}
	}


	fmt.Println(sum)
	return sum
}

func checkAdjacent(octopuses [][]int, r, c int) [][]int {
	octopuses[r][c] = 0
	if c > 0 {
		octopuses = add(octopuses, r, c-1) // left
		if r > 0 {
			octopuses = add(octopuses, r-1, c-1)  // left top
		}
		if r < len(octopuses)-1 {
			octopuses = add(octopuses, r+1, c-1) // left bottom
		}
	}
	if c < len(octopuses[r])-1 {
		octopuses = add(octopuses, r, c+1) // right
		if r > 0 {
			octopuses = add(octopuses, r-1, c+1) // right top
		}
		if r < len(octopuses)-1 {
			octopuses = add(octopuses, r+1, c+1)  // right bottom
		}
	}

	if r > 0 { // top
		octopuses = add(octopuses, r-1, c)
	}

	if r < len(octopuses)-1 { // bottom
		octopuses = add(octopuses, r+1, c)
	}

	return octopuses
}

func add(octopuses [][]int, r, c int) [][]int {
	if octopuses[r][c] != 0 {
		octopuses[r][c] += 1 // right top
	}
	if octopuses[r][c] == 10 {
		octopuses = checkAdjacent(octopuses, r, c)
	}
	return octopuses
}
