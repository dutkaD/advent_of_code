package day9

import (
	"advent_of_code/utils"
	"sort"
	"strings"
)

func GetLowestPoint(input []string) int {
	var table [][]int
	visited := make(map[int]map[int]bool)
	var basins []int

	for _, line := range input {
		numbers := strings.Split(line, "")
		var nums []int
		for _, height := range numbers {
			nums = append(nums, utils.ToInt(height))
		}
		table = append(table, nums)

	}

	riskLevel := 0
	for r, row := range table {
		for c, point := range row {
			if isLowest(r, c, table) {
				riskLevel += 1 + point

				s, v := getBasin(r, c, table, visited)

				basins = append(basins, s)
				visited = v
			}
		}
	}
	sort.Slice(basins, func(i, j int) bool {
		return basins[i] > basins[j]
	})

	return basins[0] * basins[1] * basins[2]
}

func isLowest(r, c int, table [][]int) bool {
	lowest := true
	if c != 0 {
		if table[r][c-1] <= table[r][c] {
			lowest = false
		}
	}
	if c != len(table[r])-1 {
		if table[r][c+1] <= table[r][c] {
			lowest = false
		}
	}

	if r != 0 {
		if table[r-1][c] <= table[r][c] {
			lowest = false
		}
	}

	if r != len(table)-1 {
		if table[r+1][c] <= table[r][c] {
			lowest = false
		}
	}
	return lowest
}

func getBasin(r, c int, table [][]int, visited map[int]map[int]bool) (int, map[int]map[int]bool) {
	count := 0
	if table[r][c] == 9 {
		return count, visited
	}
	if visited[r][c] == true {
		return count, visited
	}

	if visited[r] == nil {
		visited[r] = make(map[int]bool)
	}
	visited[r][c] = true
	count += 1

	if c != 0 {
		if table[r][c-1]-table[r][c] == 1 {
			s, v := getBasin(r, c-1, table, visited)
			count += s
			visited = v
		}
	}
	if c != len(table[r])-1 {
		if table[r][c+1]-table[r][c] == 1 {
			s, v := getBasin(r, c+1, table, visited)
			count += s
			visited = v
		}
	}
	if r != 0 {
		if table[r-1][c]-table[r][c] == 1 {
			s, v := getBasin(r-1, c, table, visited)
			count += s
			visited = v
		}
	}
	if r != len(table)-1 {
		if table[r+1][c]-table[r][c] == 1 {
			s, v := getBasin(r+1, c, table, visited)
			count += s
			visited = v
		}
	}

	return count, visited

}
