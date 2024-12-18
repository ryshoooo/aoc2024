package no11

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var memory map[string]int = make(map[string]int)

func blink(x, n int) int {
	if n == 0 {
		return 1
	}
	key := fmt.Sprintf("%d-%d", x, n)
	if memory[key] != 0 {
		return memory[key]
	}

	if x == 0 {
		memory[key] = blink(1, n-1)
		return memory[key]
	}

	xs := fmt.Sprintf("%d", x)
	if len(xs)%2 == 0 {
		rhs, _ := strconv.Atoi(xs[:len(xs)/2])
		lhs, _ := strconv.Atoi(xs[len(xs)/2:])
		res := blink(rhs, n-1) + blink(lhs, n-1)
		memory[key] = res
		return res
	}

	memory[key] = blink(x*2024, n-1)
	return memory[key]
}

func Solve() error {
	datab, err := os.ReadFile("11/input1.txt")
	if err != nil {
		return fmt.Errorf("could not read file: %w", err)
	}

	datas := string(datab)
	vals := strings.Split(datas, " ")
	total := 0
	for _, v := range vals {
		nv, err := strconv.Atoi(v)
		if err != nil {
			return fmt.Errorf("could not convert %s to int: %w", v, err)
		}
		total += blink(nv, 25)
	}

	fmt.Println("Result 1:", total)

	total = 0
	for _, v := range vals {
		nv, err := strconv.Atoi(v)
		if err != nil {
			return fmt.Errorf("could not convert %s to int: %w", v, err)
		}
		total += blink(nv, 75)
	}
	fmt.Println("Result 2:", total)

	return nil
}
