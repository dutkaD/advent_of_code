package utils

import "strconv"

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func ToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func Sum(ints []int) int {
	sum := 0
	for n := 0; n < len(ints); n++ {
		sum += ints[n]
	}
	return sum
}


func StringsToInts(stringVals []string) []int {
	var nums []int
	for _, s := range stringVals {
		nums = append(nums, ToInt(s))
	}
	return nums
}