package no4

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func find(val string, m [][]string) [][]int {
	var res [][]int
	for x, row := range m {
		for y, col := range row {
			if col == val {
				res = append(res, []int{x, y})
			}
		}
	}
	return res
}

func whatIsRel(c1, c2 []int) (bool, string) {
	xd := c1[0] - c2[0]
	yd := c1[1] - c2[1]
	if math.Abs(float64(xd)) > 1 || math.Abs(float64(yd)) > 1 {
		return false, ""
	}
	switch xd {
	case 0:
		switch yd {
		case 1:
			return true, "R"
		case -1:
			return true, "L"
		}
	case 1:
		switch yd {
		case 0:
			return true, "D"
		case 1:
			return true, "DR"
		case -1:
			return true, "DL"
		}
	case -1:
		switch yd {
		case 0:
			return true, "U"
		case 1:
			return true, "UR"
		case -1:
			return true, "UL"
		}
	}
	panic("Unknown direction")
}

func isACenter(ac []int, mpositions, spositions [][]int) bool {
	counter := 0
	possiblePositions := [][]int{{ac[0] - 1, ac[1] - 1}, {ac[0] + 1, ac[1] + 1}, {ac[0] - 1, ac[1] + 1}, {ac[0] + 1, ac[1] - 1}}
	for _, pmc := range possiblePositions {
		for _, mc := range mpositions {
			if pmc[0] == mc[0] && pmc[1] == mc[1] {
				_, direction := whatIsRel(mc, ac)
				for _, smc := range possiblePositions {
					for _, sc := range spositions {
						if smc[0] == sc[0] && smc[1] == sc[1] {
							_, directions := whatIsRel(ac, sc)
							if direction == directions {
								counter++
							}
						}
					}
				}
			}
		}
	}
	return counter == 2
}

func Solve() error {
	datab, err := os.ReadFile("4/input1.txt")
	if err != nil {
		return fmt.Errorf("could not read data from file: %w", err)
	}

	rows := strings.Split(string(datab), "\n")
	var m [][]string

	for _, row := range rows {
		vals := strings.Split(row, "")
		m = append(m, vals)
	}

	xpositions := find("X", m)
	mpositions := find("M", m)
	apositions := find("A", m)
	spositions := find("S", m)

	var res int

	for _, xc := range xpositions {
		for _, mc := range mpositions {
			ok, direction := whatIsRel(xc, mc)
			if ok {
				for _, ac := range apositions {
					ok, directiona := whatIsRel(mc, ac)
					if ok && direction == directiona {
						for _, sc := range spositions {
							ok, directions := whatIsRel(ac, sc)
							if ok && directiona == directions {
								res++
							}
						}
					}
				}
			}
		}
	}

	fmt.Println("Result 1:", res)

	res = 0
	for _, ac := range apositions {
		if isACenter(ac, mpositions, spositions) {
			res++
		}
	}

	fmt.Println("Result 2:", res)
	return nil
}
