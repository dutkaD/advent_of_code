package main

import (
	"advent_of_code/day1"
	utils2 "advent_of_code/utils"
	"fmt"
)

func main() {
	lines := utils2.ReadFile("../input.txt")
	numbers := stringsToInts(lines)
	// part 1
	fmt.Println(day1.GetIncreased(numbers))
	//part 2
	fmt.Println(day1.GetWithWindows(numbers))

}

func stringsToInts(stringVals []string) []int {
	var nums []int
	for _, s := range stringVals {
		nums = append(nums, utils2.ToInt(s))
	}
	return nums
}
