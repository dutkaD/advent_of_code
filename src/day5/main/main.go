package main

import (
	"advent_of_code/src/day5"
	"advent_of_code/src/utils"
	"fmt"
)

func main() {
	lines := utils.ReadFile("../input.txt")

	// part 1
	fmt.Println(day5.GetOverlappingLines(lines, false))
	// part 2
	fmt.Println(day5.GetOverlappingLines(lines, true))
}
