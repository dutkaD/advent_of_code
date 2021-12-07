package main

import (
	"advent_of_code/day7"
	"advent_of_code/utils"
	"fmt"
)

func main() {
	lines := utils.ReadFile("input.txt")
	_ = lines
	fmt.Println(day7.AlignCrabs(lines))
}
