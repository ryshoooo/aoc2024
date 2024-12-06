package no6

import (
	"fmt"
	"os"
	"strings"
)

func parseMap(m string) ([]int, [][]int, int, int) {
	startingPosition := make([]int, 2)
	obstacles := [][]int{}
	rows := strings.Split(m, "\n")
	var cbound int
	for indx, row := range rows {
		els := strings.Split(row, "")
		cbound = len(els)
		for indy, col := range els {
			if col == "^" {
				startingPosition[0] = indx
				startingPosition[1] = indy
			} else if col == "#" {
				obstacles = append(obstacles, []int{indx, indy})
			}
		}
	}

	return startingPosition, obstacles, len(rows), cbound
}

func contains(arr [][]int, el []int) bool {
	for _, al := range arr {
		if al[0] == el[0] && al[1] == el[1] {
			return true
		}
	}
	return false
}

func hasLoop(spos []int, obs [][]int, rb, cb int) bool {
	combs := [4][]int{{0, -1}, {1, 1}, {0, 1}, {1, -1}}
	cidx := 0
	isReturn := false
	visitedLocations := map[string][]int{}
	for {
		if spos[0] < 0 || spos[1] < 0 || spos[0] >= rb || spos[1] >= cb {
			break
		}
		if v, ok := visitedLocations[fmt.Sprintf("%d-%d", spos[0], spos[1])]; ok {
			if v[0] == combs[cidx][0] && v[1] == combs[cidx][1] {
				return true
			}
		}
		if contains(obs, spos) {
			spos[combs[cidx][0]] -= combs[cidx][1]
			isReturn = true
			cidx = (cidx + 1) % 4
			continue
		}
		if !isReturn {
			visitedLocations[fmt.Sprintf("%d-%d", spos[0], spos[1])] = combs[cidx]
		}
		isReturn = false
		spos[combs[cidx][0]] = spos[combs[cidx][0]] + combs[cidx][1]
	}

	return false
}

func Solve() error {
	datab, err := os.ReadFile("6/input1.txt")
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	datas := string(datab)
	spos, obs, rb, cb := parseMap(datas)
	combs := [4][]int{{0, -1}, {1, 1}, {0, 1}, {1, -1}}
	cidx := 0
	isReturn := false

	visitedLocations := map[string]int{}
	for {
		if spos[0] < 0 || spos[1] < 0 || spos[0] >= rb || spos[1] >= cb {
			break
		}
		if contains(obs, spos) {
			spos[combs[cidx][0]] -= combs[cidx][1]
			isReturn = true
			cidx = (cidx + 1) % 4
			continue
		}
		if !isReturn {
			visitedLocations[fmt.Sprintf("%d-%d", spos[0], spos[1])] += 1
		}
		isReturn = false
		spos[combs[cidx][0]] = spos[combs[cidx][0]] + combs[cidx][1]
	}

	fmt.Println("Result 1:", len(visitedLocations))
	spos, obs, rb, cb = parseMap(datas)
	res := 0
	for x := range rb {
		for y := range cb {
			newObs := []int{x, y}
			if contains(obs, newObs) {
				continue
			}

			if x == spos[0] && y == spos[1] {
				continue
			}

			obsCopy := make([][]int, len(obs))
			sposCopy := make([]int, len(spos))
			copy(obsCopy, obs)
			copy(sposCopy, spos)
			obsCopy = append(obsCopy, newObs)
			if hasLoop(sposCopy, obsCopy, rb, cb) {
				res++
			}
		}
	}

	fmt.Println("Result 2:", res)
	return nil
}
