package day5

import (
	"advent_of_code/utils"
	"strconv"
	"strings"
)

type line struct {
	pointA point
	pointB point
}

type point struct {
	x int
	y int
}

func (line *line) is45Degrees() bool {
	if utils.Abs(line.pointA.x-line.pointB.x) == utils.Abs(line.pointA.y-line.pointB.y) {
		return true
	}
	return false
}

func GetOverlappingLines(textByLines []string, withDiagonal bool) int {
	ventMap := make(map[int]map[int]int)
	lines := readVentLines(textByLines)
	overlapping := 0

	for l := 0; l < len(lines); l++ {
		pointA := lines[l].pointA
		pointB := lines[l].pointB

		// part 1
		if pointA.x == pointB.x {
			for p := utils.Min(pointA.y, pointB.y); p <= utils.Max(pointA.y, pointB.y); p++ {
				if ventMap[pointA.x] == nil {
					ventMap[pointA.x] = make(map[int]int)
				}
				if ventMap[pointA.x][p] == 1 {
					overlapping += 1
				}
				ventMap[pointA.x][p] += 1
			}
		}
		if pointA.y == pointB.y {
			for p := utils.Min(pointA.x, pointB.x); p <= utils.Max(pointA.x, pointB.x); p++ {
				if ventMap[p] == nil {
					ventMap[p] = make(map[int]int)
				}
				if ventMap[p][pointA.y] == 1 {
					overlapping += 1
				}
				ventMap[p][pointA.y] += 1
			}
		}

		// part 2
		if withDiagonal {
			if lines[l].is45Degrees() {
				diff := utils.Abs(pointA.x - pointB.x)
				stepX := 1
				stepY := 1
				if pointA.x-pointB.x > 0 {
					stepX = -stepX
				}
				if pointA.y-pointB.y > 0 {
					stepY = -stepY
				}

				for n := 0; n <= diff; n++ {
					x := pointA.x + stepX*n
					y := pointA.y + stepY*n
					if ventMap[x] == nil {
						ventMap[x] = make(map[int]int)
					}
					if ventMap[x][y] == 1 {
						overlapping += 1
					}
					ventMap[x][y] += 1
				}
			}
		}
	}

	return overlapping
}

func readVentLines(textByLines []string) []line {
	var lines []line
	for l := 0; l < len(textByLines); l++ {
		splitLine := strings.Split(textByLines[l], " ")
		pointA := parseCoordinates(splitLine[0])
		pointB := parseCoordinates(splitLine[2])
		lines = append(lines, line{pointA, pointB})
	}
	return lines
}

func parseCoordinates(xy string) point {
	points := strings.Split(xy, ",")
	x, _ := strconv.Atoi(points[0])
	y, _ := strconv.Atoi(points[1])
	return point{x, y}
}
