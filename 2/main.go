package no2

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func parseReport(report string) ([]int, error) {
	levelsS := strings.Split(report, " ")
	var levels []int
	for _, levelS := range levelsS {
		level, err := strconv.Atoi(levelS)
		if err != nil {
			return nil, fmt.Errorf("could not parse level: %w", err)
		}
		levels = append(levels, level)
	}
	return levels, nil
}

func isSafe(report []int) bool {
	x := report[0]
	y := report[1]
	isInc := y > x
	for _, y := range report[1:] {
		diff := math.Abs(float64(y - x))
		if diff < 1 || diff > 3 {
			return false
		}
		if (y > x && !isInc) || (y < x && isInc) {
			return false
		}
		x = y
	}
	return true
}

func remove(slice []int, s int) []int {
	var nr []int
	for ind, val := range slice {
		if ind != s {
			nr = append(nr, val)
		}
	}
	return nr
}

func Solve() error {
	data, err := os.ReadFile("2/input1.txt")
	if err != nil {
		return fmt.Errorf("could not read input file: %w", err)
	}

	datas := string(data)
	reportsRaw := strings.Split(datas, "\n")
	var nSafe int
	for _, report := range reportsRaw {
		if report == "" {
			continue
		}

		reportI, err := parseReport(report)
		if err != nil {
			return fmt.Errorf("could not parse report: %w", err)
		}
		if isSafe(reportI) {
			nSafe++
		}
	}

	fmt.Println("Result 1:", nSafe)

	nSafe = 0
	for _, report := range reportsRaw {
		if report == "" {
			continue
		}

		reportI, err := parseReport(report)
		if err != nil {
			return fmt.Errorf("could not parse report: %w", err)
		}
		if isSafe(reportI) {
			nSafe++
		} else {
			for i := range reportI {
				nrepo := remove(reportI, i)
				if isSafe(nrepo) {
					nSafe++
					break
				}
			}
		}
	}

	fmt.Println("Result 2:", nSafe)
	return nil
}
