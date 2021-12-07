package day7

import (
	"advent_of_code/utils"
	"strings"
)

func AlignCrabs(input []string) int {
	positions := utils.StringsToInts(strings.Split(input[0], ","))

	min := positions[0]
	max := positions[0]
	mapped := make(map[int]int)

	for x := 0; x < len(positions); x++ {
		if positions[x] > max{
			max = positions[x]
		}
		if positions[x] < min{
			min = positions[x]
		}
		mapped[positions[x]] += 1
	}


	var fuelValues []int
	for p := min; p < max; p++ {
		fuel := 0
		for key, val := range mapped {
			steps := utils.Abs(key - p)
			realSteps := 0
			for x := 1; x <= steps; x++ {
				realSteps += x
			}
			fuel += realSteps * val
		}

		fuelValues = append(fuelValues, fuel)
	}

	return utils.MinValue(fuelValues)
}
