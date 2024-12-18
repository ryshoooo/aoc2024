package no9

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findEmptyLocations(arr []string) ([]int, []int, error) {
	res := []int{}
	counts := []int{}
	sidx := 0
	for ind, x := range arr {
		xi, err := strconv.Atoi(x)
		if err != nil {
			return []int{}, []int{}, fmt.Errorf("could not convert %s to int: %w", x, err)
		}

		if ind%2 == 0 {
			sidx += xi
			counts = append(counts, xi)
		} else {
			sidx += xi
			res = append(res, xi)
		}
	}

	return res, counts, nil
}

func constructOptimal(emptyLocations []int, counts []int) []int {
	res := []int{}

	countsCopy := make([]int, len(counts))
	emptyLocationsCopy := make([]int, len(emptyLocations))
	copy(countsCopy, counts)
	copy(emptyLocationsCopy, emptyLocations)

	lidx := 0
	eidx := 0
	ridx := len(countsCopy) - 1

	for {
		if lidx >= len(countsCopy) {
			break
		}

		for range countsCopy[lidx] {
			res = append(res, lidx)
		}
		countsCopy[lidx] = 0
		lidx++

		if eidx >= len(emptyLocationsCopy) {
			continue
		}

		for range emptyLocationsCopy[eidx] {
			if ridx >= 0 && countsCopy[ridx] == 0 {
				for {
					if ridx < 0 {
						break
					}
					if countsCopy[ridx] != 0 {
						break
					}
					ridx--
				}
			}

			if ridx < 0 {
				res = append(res, 0)
			} else {
				if countsCopy[ridx] == 0 {
					panic("should not happen")
				}
				res = append(res, ridx)
				countsCopy[ridx]--
			}
		}
		emptyLocationsCopy[eidx] = 0
		eidx++

	}
	return res
}

type Node struct {
	Index int
	Count int
}

func Solve() error {
	datab, err := os.ReadFile("9/input1.txt")
	if err != nil {
		return fmt.Errorf("could not read input file: %w", err)
	}

	datas := strings.Split(string(datab), "")
	emptyLocs, counts, err := findEmptyLocations(datas)
	if err != nil {
		return fmt.Errorf("could not find empty locations: %w", err)
	}
	opt := constructOptimal(emptyLocs, counts)
	res := 0
	for ind, x := range opt {
		res += x * ind
	}

	fmt.Println("Result 1:", res)

	gcounts := make([]*Node, len(counts))
	idx := 0
	for ind, count := range counts {
		gcounts[ind] = &Node{Index: idx, Count: count}
		if len(emptyLocs) > ind {
			idx = idx + count + emptyLocs[ind]
		}
	}

	gemptylocs := make([]*Node, len(emptyLocs))
	for idx, emptyLoc := range emptyLocs {
		gemptylocs[idx] = &Node{Index: idx, Count: emptyLoc}
	}

	for i := len(gcounts) - 1; i >= 0; i-- {
		c := gcounts[i]
		for _, emptyLoc := range gemptylocs[:i] {
			if emptyLoc.Count >= c.Count {
				emptyLoc.Count -= c.Count
				currc := gcounts[emptyLoc.Index]
				c.Index = currc.Index + currc.Count
				emptyLoc.Index = i
				break
			}
		}
	}

	total := 0
	for idx, ucount := range gcounts {
		for j := ucount.Index; j < ucount.Index+ucount.Count; j++ {
			total += j * idx
		}
	}

	fmt.Println("Result 2:", total)
	return nil
}
