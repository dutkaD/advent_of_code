package main

import (
	"advent_of_code/day2"
	"advent_of_code/utils"
	"fmt"
)

func main() {

	lines := utils.ReadFile("../input.txt")
	// part 1
	fmt.Println(day2.CalculateLocation(lines))
	// part 2
	fmt.Println(day2.CalculateLocationWithAim(lines))

}
