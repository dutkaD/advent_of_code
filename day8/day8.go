package day8

import (
	"advent_of_code/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func GetEasyNumbers(input []string) int {
	total := 0
	for _, line := range input {
		splitLines := strings.Split(line, "|")
		inSegments := strings.Fields(splitLines[0])
		outSegments := strings.Fields(splitLines[1])
		numberMap := decodeEasyNumbers(inSegments)
		_ = outSegments
		//fmt.Println(inSegments)
		//fmt.Println(outSegments)
		numberMap = decodeOtherNumbers(inSegments, numberMap)

		segmentNumber := ""
		for _, seg := range outSegments{
			sorSeg := utils.SortString(seg)
			for num, code := range numberMap {
				sort.Strings(code)
				sortCode := strings.Join(code[:], "")
				if sorSeg == sortCode {
					segmentNumber += strconv.Itoa(num)
				}
			}
		}

		fmt.Println(segmentNumber)
		total += utils.ToInt(segmentNumber)


	}
	return total

}

func decodeEasyNumbers(inSegments []string) map[int][]string {
	numberMap := make(map[int][]string)

	for _, seg := range inSegments {
		switch {
		case len(seg) == 2:
			numberMap[1] = strings.Split(seg, "")
		case len(seg) == 4:
			numberMap[4] = strings.Split(seg, "")
		case len(seg) == 3:
			numberMap[7] = strings.Split(seg, "")
		case len(seg) == 7:
			numberMap[8] = strings.Split(seg, "")
		default:
			fmt.Print("")
		}

	}
	return numberMap
}

func decodeOtherNumbers(inSegments []string, numberMap map[int][]string) map[int][]string {
	for _, seg := range inSegments {
		current := strings.Split(seg, "")

		switch {
		case len(seg) == 6:
			eight := numberMap[8]
			diffWithEight := difference(eight, current)
			diffWithFour := difference(numberMap[4], current)
			if len(diffWithEight) == 1 {
				if utils.Contains(numberMap[4], diffWithEight[0]) && !utils.Contains(numberMap[1], diffWithEight[0]){
					numberMap[0] = current
				}

				if utils.Contains(numberMap[1], diffWithEight[0])   {
					numberMap[6] = current
				}
			}

			if len(diffWithFour) == 0 {
				numberMap[9] = current

			}

		default:
			fmt.Println("")
		}
	}

	for _, seg := range inSegments {
		current := strings.Split(seg, "")

		if len(seg) == 5 {
			differenceWithSix := difference(numberMap[6], current)
			differenceWithOne := difference(numberMap[1], current)
			if len(differenceWithOne) == 0 {
				numberMap[3] = current
			}else{
				if len(differenceWithSix) == 2{
					numberMap[2] =current
				}
			}
			if len(differenceWithSix) == 1{
				numberMap[5] = current
			}

		}

	}


	return numberMap
}

func difference(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}
