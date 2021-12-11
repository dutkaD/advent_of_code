package main

import (
	"advent_of_code/day9"
	"advent_of_code/utils"
	"fmt"
)

func main() {
	lines := utils.ReadFile("input.txt")
	_ = lines
	fmt.Println(day9.GetLowestPoint(lines))
}
