package main

import (
	"advent_of_code/day3"
	"advent_of_code/utils"
	"fmt"
)

func main() {
	lines := utils.ReadFile("../input.txt")
	// part 1
	fmt.Println(day3.GetConsumption(lines))
	//part 2
	fmt.Println(day3.GetLifeSupportRating(lines))
}
