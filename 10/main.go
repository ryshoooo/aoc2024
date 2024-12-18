package no10

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TrailNode struct {
	Height        int
	ReachablePeks map[string]string
	Trails        map[string]string
}

func makeMap(data string) ([][]*TrailNode, error) {
	res := [][]*TrailNode{}

	rows := strings.Split(data, "\n")
	rowIdx := 0
	for _, row := range rows {
		vals := strings.Split(row, "")
		resrow := make([]*TrailNode, len(vals))
		for ind, val := range vals {
			v, err := strconv.Atoi(val)
			if err != nil {
				return res, fmt.Errorf("could not convert %s to int: %w", val, err)
			}
			rp := make(map[string]string)
			trails := make(map[string]string)
			if v == 9 {
				rp[fmt.Sprintf("%d-%d", rowIdx, ind)] = ""
			}
			resrow[ind] = &TrailNode{Height: v, ReachablePeks: rp, Trails: trails}
		}
		res = append(res, resrow)
		rowIdx++
	}
	return res, nil
}

func findHeight(m [][]*TrailNode, height int) [][]int {
	res := [][]int{}
	for x, row := range m {
		for y, tn := range row {
			if tn.Height == height {
				res = append(res, []int{x, y})
			}
		}
	}
	return res
}

func populate(m [][]*TrailNode, x, y, comingFromHeight int, reachablePeaks map[string]string, trails map[string]string) {
	if x < 0 || y < 0 || x >= len(m) || y >= len(m[x]) {
		return
	}

	tn := m[x][y]
	if comingFromHeight != tn.Height+1 {
		return
	}

	rp := make(map[string]string)
	ts := make(map[string]string)
	for k, v := range reachablePeaks {
		rp[k] = v
	}
	for k, v := range tn.ReachablePeks {
		rp[k] = v
	}
	for k, v := range trails {
		ts[fmt.Sprintf("%s,%d-%d", k, x, y)] = v
	}
	for k, v := range tn.Trails {
		ts[k] = v
	}
	tn.ReachablePeks = rp
	tn.Trails = ts

	if tn.Height == 0 {
		return
	}

	populate(m, x-1, y, tn.Height, rp, ts)
	populate(m, x+1, y, tn.Height, rp, ts)
	populate(m, x, y-1, tn.Height, rp, ts)
	populate(m, x, y+1, tn.Height, rp, ts)

	return
}

func Solve() error {
	datab, err := os.ReadFile("10/input1.txt")
	if err != nil {
		return fmt.Errorf("could not read input file: %w", err)
	}

	datas := string(datab)

	trailmap, err := makeMap(datas)
	if err != nil {
		return fmt.Errorf("could not make map: %w", err)
	}

	peaks := findHeight(trailmap, 9)
	for _, peak := range peaks {
		trails := make(map[string]string)
		trails[fmt.Sprintf("%d-%d", peak[0], peak[1])] = ""
		populate(trailmap, peak[0], peak[1], 10, trailmap[peak[0]][peak[1]].ReachablePeks, trails)
	}

	starts := findHeight(trailmap, 0)
	total := 0
	total2 := 0
	for _, start := range starts {
		total += len(trailmap[start[0]][start[1]].ReachablePeks)
		total2 += len(trailmap[start[0]][start[1]].Trails)
	}
	fmt.Println("Result 1:", total)
	fmt.Println("Result 2:", total2)
	return nil
}
