package no13

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readEquations(data string) ([][][]int, error) {
	res := [][][]int{}

	dspl := strings.Split(data, "\n")
	var ax, ay, bx, by, xp, yp int
	var err error
	for ind, d := range dspl {
		if ind%4 == 0 {
			axr := strings.Split(strings.Split(strings.Split(d, " ")[2], "+")[1], ",")[0]
			ayr := strings.Split(strings.Split(d, " ")[3], "+")[1]
			ax, err = strconv.Atoi(axr)
			if err != nil {
				return nil, fmt.Errorf("could not convert ax: %w", err)
			}
			ay, err = strconv.Atoi(ayr)
			if err != nil {
				return nil, fmt.Errorf("could not convert ay: %w", err)
			}
		}
		if ind%4 == 1 {
			bxr := strings.Split(strings.Split(strings.Split(d, " ")[2], "+")[1], ",")[0]
			byr := strings.Split(strings.Split(d, " ")[3], "+")[1]
			bx, err = strconv.Atoi(bxr)
			if err != nil {
				return nil, fmt.Errorf("could not convert bx: %w", err)
			}
			by, err = strconv.Atoi(byr)
			if err != nil {
				return nil, fmt.Errorf("could not convert by: %w", err)
			}
		}
		if ind%4 == 2 {
			xpr := strings.Split(strings.Split(strings.Split(d, " ")[1], "=")[1], ",")[0]
			ypr := strings.Split(strings.Split(d, " ")[2], "=")[1]
			xp, err = strconv.Atoi(xpr)
			if err != nil {
				return nil, fmt.Errorf("could not convert xp: %w", err)
			}
			yp, err = strconv.Atoi(ypr)
			if err != nil {
				return nil, fmt.Errorf("could not convert yp: %w", err)
			}

		}
		if ind%4 == 3 {
			res = append(res, [][]int{{ax, bx, xp}, {ay, by, yp}})
		}
	}
	return res, nil
}

func SolveEquation(eq [][]int, priceIncrease int) int {
	eq[0][2] += priceIncrease
	eq[1][2] += priceIncrease
	a1 := eq[0][0]
	a2 := eq[1][0]
	b := eq[0][1]*a2 - eq[1][1]*a1
	p := eq[0][2]*a2 - eq[1][2]*a1
	if p%b != 0 {
		return 0
	}
	bsolv := p / b
	if bsolv < 0 {
		return 0
	}
	a := eq[0][2] - eq[0][1]*bsolv
	if a%eq[0][0] != 0 {
		return 0
	}
	asolv := a / eq[0][0]
	if asolv < 0 {
		return 0
	}
	return asolv*3 + bsolv
}

func Solve() error {
	datab, err := os.ReadFile("13/input1.txt")
	if err != nil {
		return fmt.Errorf("could not read file: %w", err)
	}

	datas := string(datab)
	equations, err := readEquations(datas)
	if err != nil {
		return fmt.Errorf("could not read equations: %w", err)
	}

	tokens := 0
	tokens2 := 0
	for _, eq := range equations {
		tokens += SolveEquation(eq, 0)
		tokens2 += SolveEquation(eq, 10000000000000)
	}

	fmt.Println("Result 1:", tokens)
	fmt.Println("Result 2:", tokens2)
	return nil
}
