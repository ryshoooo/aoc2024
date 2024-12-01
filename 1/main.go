package no1

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Solve() error {
	f, err := os.ReadFile("1/input1.txt")
	if err != nil {
		return fmt.Errorf("could not read file: %w", err)
	}

	input := string(f)
	entries := strings.Split(input, "\n")

	var ll, rl []int
	for _, entry := range entries {
		if entry == "" {
			continue
		}

		nos := strings.Split(entry, "   ")
		if len(nos) != 2 {
			return fmt.Errorf("invalid entry: %s (%s)", entry, nos)
		}

		le, err := strconv.Atoi(nos[0])
		if err != nil {
			return fmt.Errorf("invalid entry: %s (%s)", entry, nos[0])
		}

		re, err := strconv.Atoi(nos[1])
		if err != nil {
			return fmt.Errorf("invalid entry: %s (%s)", entry, nos[1])
		}

		ll = append(ll, le)
		rl = append(rl, re)
	}

	slices.Sort(ll)
	slices.Sort(rl)

	var result int
	for i, lv := range ll {
		x := rl[i] - lv
		if x < 0 {
			x = -x
		}
		result += x
	}

	fmt.Println(fmt.Sprintf("Result 1: %d", result))

	counter := make(map[int]int)
	for _, rv := range rl {
		counter[rv]++
	}

	result = 0
	for _, lv := range ll {
		result += (lv * counter[lv])
	}

	fmt.Println(fmt.Sprintf("Result 2: %d", result))
	return nil
}
