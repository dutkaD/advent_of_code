package main

import (
	"advent_of_code/src/day1"
	"advent_of_code/src/utils"
	"fmt"
)

func main() {
	lines := utils.ReadFile("../input.txt")
	numbers := stringsToInts(lines)
	// part 1
	fmt.Println(day1.GetIncreased(numbers))
	//part 2
	fmt.Println(day1.GetWithWindows(numbers))

}

func stringsToInts(stringVals []string) []int {
	var nums []int
	for _, s := range stringVals {
		nums = append(nums, utils.ToInt(s))
	}
	return nums
}
