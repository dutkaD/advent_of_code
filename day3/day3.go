package day3

import (
	"fmt"
	"strconv"
)

func GetConsumption(reports []string) int {
	stats := getStats(reports)
	return toDecimal(getGamma(stats)) * toDecimal(getEpsilon(stats))
}

func GetLifeSupportRating(reports []string) int {
	return toDecimal(getOxygen(reports)) * toDecimal(getScrubber(reports))
}

func getStats(reports []string) map[int]map[int]int {
	stats := make(map[int]map[int]int)

	for a := 0; a < len(reports); a++ {
		for b := 0; b < len(reports[a]); b++ {
			s := reports[a]
			num, _ := strconv.Atoi(s[b : b+1])
			if stats[b] == nil {
				stats[b] = make(map[int]int)
			}
			stats[b][num] += 1
		}
	}
	return stats
}

func toDecimal(binary string) int {
	if i, err := strconv.ParseInt(binary, 2, 64); err != nil {
		fmt.Println(err)
	} else {
		return int(i)
	}
	return -1
}

func getGamma(stats map[int]map[int]int) string {
	gamma := ""
	for pos := 0; pos < len(stats); pos++ {
		if stats[pos][0] > stats[pos][1] {
			gamma += "0"
		} else {
			gamma += "1"
		}
	}
	return gamma
}

func getEpsilon(stats map[int]map[int]int) string {
	epsilon := ""
	for pos := 0; pos < len(stats); pos++ {
		if stats[pos][0] < stats[pos][1] {
			epsilon += "0"
		} else {
			epsilon += "1"
		}
	}
	return epsilon
}

func filterBy(reports []string, pos int, filter int) []string {
	var filtered []string
	for a := 0; a < len(reports); a++ {
		num, _ := strconv.Atoi(reports[a][pos : pos+1])
		if num == filter {
			filtered = append(filtered, reports[a])
		}
	}
	return filtered
}

func getOxygen(reports []string) string {
	for a := 0; a < len(reports[0]) && len(reports) > 1; a++ {
		newStats := getStats(reports)
		if newStats[a][0] > newStats[a][1] {
			reports = filterBy(reports, a, 0)
		}
		if newStats[a][0] < newStats[a][1] {
			reports = filterBy(reports, a, 1)
		}
		if newStats[a][0] == newStats[a][1] {
			reports = filterBy(reports, a, 1)
		}

	}
	return reports[0]
}

func getScrubber(reports []string) string {
	for a := 0; a < len(reports[0]) && len(reports) > 1; a++ {
		newStats := getStats(reports)
		if newStats[a][0] < newStats[a][1] {
			reports = filterBy(reports, a, 0)
		}
		if newStats[a][0] > newStats[a][1] {
			reports = filterBy(reports, a, 1)
		}
		if newStats[a][0] == newStats[a][1] {
			reports = filterBy(reports, a, 0)
		}

	}
	return reports[0]
}
