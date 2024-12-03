package no3

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func mul(s string) (int, error) {
	tuple := strings.ReplaceAll(strings.ReplaceAll(s, "mul(", ""), ")", "")
	xys := strings.Split(tuple, ",")
	x, err := strconv.Atoi(xys[0])
	if err != nil {
		return 0, fmt.Errorf("could not convert %s to int: %w", xys[0], err)
	}
	y, err := strconv.Atoi(xys[1])
	if err != nil {
		return 0, fmt.Errorf("could not convert %s to int: %w", xys[1], err)
	}

	return x * y, nil
}

func findClosestIndex(startingIndex int, arr [][]int) int {
	result := -1
	for _, tup := range arr {
		if tup[0] > startingIndex {
			return result
		} else {
			result = tup[0]
		}
	}
	return result
}

func Solve() error {
	datab, err := os.ReadFile("3/input1.txt")
	if err != nil {
		return fmt.Errorf("could not read data from file: %w", err)
	}

	datas := string(datab)
	ret := regexp.MustCompile(`mul\(\d+,\d+\)`)
	results := ret.FindAllString(datas, -1)

	var fr int
	for _, res := range results {
		m, err := mul(res)
		if err != nil {
			return fmt.Errorf("could not multiply: %w", err)
		}
		fr += m
	}

	fmt.Println("Result 1:", fr)

	fr = 0
	redos := regexp.MustCompile(`do\(\)`)
	redonts := regexp.MustCompile(`don\'t\(\)`)
	mulidx := ret.FindAllIndex(datab, -1)
	doidx := redos.FindAllIndex(datab, -1)
	dontsidx := redonts.FindAllIndex(datab, -1)

	for _, mti := range mulidx {
		startIndex := mti[0]
		endIndex := mti[1]
		latestDoIndex := findClosestIndex(startIndex, doidx)
		latestDontIndex := findClosestIndex(startIndex, dontsidx)

		if (latestDoIndex == -1 && latestDontIndex == -1) || (latestDoIndex > latestDontIndex) {
			r, err := mul(string(datab[startIndex:endIndex]))
			if err != nil {
				return fmt.Errorf("could not multiply: %w", err)
			}
			fr += r
		}

	}

	fmt.Println("Result 2:", fr)
	return nil
}
