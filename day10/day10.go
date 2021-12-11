package day10

import (
	"fmt"
	"sort"
	"strings"
)

func GetSyntxErrors(input []string) {
	sum := 0
	var scores []int
	for _, line := range input {
		var stack []string
		chars := strings.Split(line, "")
		stack = append(stack, chars[0])
		stop := false
		for c := 1; c < len(chars); c++ {
			if stop {
				break
			}
			if isOpening(chars[c]) {
				stack = append(stack, chars[c])
			} else {
				switch chars[c] {
				case ")":
					if stack[len(stack)-1] == "(" {
						stack = stack[:len(stack)-1]

					} else {
						stop = true
						sum += 3
						break
					}
				case "}":
					if stack[len(stack)-1] == "{" {
						stack = stack[:len(stack)-1]

					} else {
						stop = true
						sum += 1197
						break
					}
				case ">":
					if stack[len(stack)-1] == "<" {
						stack = stack[:len(stack)-1]

					} else {
						stop = true
						sum += 25137
						break
					}

				case "]":
					if stack[len(stack)-1] == "[" {
						stack = stack[:len(stack)-1]

					} else {
						stop = true
						sum += 57
						break
					}
				default:
					break
				}

			}
		}
		if !stop {
			score := 0
			for v := len(stack) - 1; v >= 0; v-- {
				score *= 5
				switch stack[v] {
				case "(":
					score += 1
				case "[":
					score += 2
				case "{":
					score += 3
				case "<":
					score += 4
				default:
				}

			}
			scores = append(scores, score)
		}
	}
	sort.Slice(scores, func(i, j int) bool {
		return scores[i] < scores[j]
	})

	fmt.Println("error score", sum)

	// part 2
	fmt.Println("mid", scores[len(scores) / 2])
}

func isOpening(c string) bool {
	if c == "[" || c == "(" || c == "{" || c == "<" {
		return true
	}
	return false
}
