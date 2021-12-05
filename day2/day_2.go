package day2

import (
	"fmt"
	"strconv"
	"strings"
)

func CalculateLocation(input []string) int {
	moves := splitBySpace(input)

	horizontal := 0
	vertical := 0
	for a := 0; a < len(moves); a++ {
		direction := moves[a][0]
		steps, _ := strconv.Atoi(moves[a][1])
		switch {
		case direction == "forward":
			horizontal += steps
		case direction == "down":
			vertical += steps
		case direction == "up":
			vertical -= steps
		default:
			fmt.Println("nowhere")
		}

	}
	return vertical * horizontal
}

func CalculateLocationWithAim(input []string) int {
	moves := splitBySpace(input)
	horizontal := 0
	aim := 0
	depth := 0
	for a := 0; a < len(moves); a++ {
		direction := moves[a][0]
		steps, _ := strconv.Atoi(moves[a][1])
		switch {
		case direction == "forward":
			horizontal += steps
			if aim != 0 {
				depth += steps * aim
			}
		case direction == "down":
			aim += steps
		case direction == "up":
			aim -= steps
		default:
			fmt.Println("nope")
		}
	}

	return depth * horizontal
}

func splitBySpace(lines []string) [][]string {
	var splitLines [][]string

	for _, s := range lines {
		splitLines = append(splitLines, strings.Split(s, " "))
	}
	return splitLines
}
