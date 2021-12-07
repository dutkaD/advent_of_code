package main

import (
	"advent_of_code/day1"
	utils2 "advent_of_code/utils"
	"fmt"
)

func main() {
	lines := utils2.ReadFile("../input.txt")
	numbers := utils2.StringsToInts(lines)
	// part 1
	fmt.Println(day1.GetIncreased(numbers))
	//part 2
	fmt.Println(day1.GetWithWindows(numbers))

}

