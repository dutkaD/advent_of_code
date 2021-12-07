package day6

import (
	"advent_of_code/utils"
	"strings"
)

func GetFishCount(input []string) int {
	fish := utils.StringsToInts(strings.Split(input[0], ","))

	sum  := 0

	// part 1
	sum = countFishInTheMostInefficientWayPossible(fish, 18, 18, 0)

	// part 2
	fishSorted := make(map[int]int)

	for k := 0; k < len(fish); k++ {
		fishSorted[fish[k]] += 1
	}
	for key, val := range fishSorted {
		fishy := countFish(key, 255)
		sum += fishy * val
	}

	return sum
}

func countFish(start int, daysLeft int) int {
	fishCount := 1
	daysLeft -= start
	for daysLeft >= 0 {
		fishCount += countFish(9, daysLeft)
		daysLeft -= 7
	}
	return fishCount
}

func countFishInTheMostInefficientWayPossible(fish []int, daysLeft int, daysAll int, newFish int) int {
	var result int

	if daysLeft == 0 {
		return len(fish)
	}

	zeros := 0
	for zeros == 0 {
		//fmt.Println("DAY", daysAll-daysLeft)
		daysLeft -= 1

		for f := 0; f < len(fish); f++ {
			if fish[f] == 0 {
				fish[f] = 6
			} else {
				fish[f] -= 1
				if fish[f] == 0 {
					zeros += 1
				}
			}
		}
		for n:=0; n < newFish; n++ {
			fish = append(fish, 8)
		}
		newFish = 0
	}


	for n:=0; n < newFish; n++ {
		fish = append(fish, 8)
	}

	result = countFishInTheMostInefficientWayPossible(fish, daysLeft, daysAll, zeros)
	return result
}