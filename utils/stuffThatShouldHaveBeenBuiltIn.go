package utils

import (
	"sort"
	"strconv"
	"strings"
)

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

func MinValue(numbers []int) (i int) {
	if len(numbers) > 0 {
		min := numbers[0]
		for x := 0; x < len(numbers); x++ {
			if numbers[x] < min {
				min = numbers[x]
			}
		}
		return min
	}
	return
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

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}