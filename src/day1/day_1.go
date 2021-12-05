package day1

func GetIncreased(numbers []int) int {

	sum := 0
	for a := 1; a < len(numbers); a++ {
		if numbers[a] > numbers[a-1] {
			sum++
		}
	}
	return sum
}

func GetWithWindows(numbers []int) int {
	var windowSums []int
	for a := 0; a < len(numbers)-2; a++ {
		window := numbers[a : a+3]
		windowSums = append(windowSums, sumOfSlice(window))
	}
	return GetIncreased(windowSums)
}

func sumOfSlice(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
