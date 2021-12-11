package main

import (
	"advent_of_code/day8"
	"advent_of_code/utils"
	"fmt"
)

func main() {
	lines := utils.ReadFile("input.txt")
	_ = lines
	fmt.Println(day8.GetEasyNumbers(lines))
}
