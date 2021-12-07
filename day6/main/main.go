package main

import (
	"advent_of_code/day6"
	"advent_of_code/utils"
	"fmt"
)

func main() {
	lines := utils.ReadFile("input.txt")
	// part 1 / 2
	fmt.Println(day6.GetFishCount(lines))
}
