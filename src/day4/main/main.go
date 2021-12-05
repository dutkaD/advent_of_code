package main

import (
	"advent_of_code/src/day4"
	"advent_of_code/src/utils"
	"fmt"
)

func main() {
	lines := utils.ReadFile("../input.txt")

	// part 1
	fmt.Println(day4.GetBingoWinner(lines))
	// part 2
	fmt.Println(day4.GetLosingBoard(lines))

}
