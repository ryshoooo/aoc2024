package no12

import (
	"fmt"
	"os"
	"strings"
)

type GardenNode struct {
	Value  string
	Region int
}

func readGarden(data string) [][]*GardenNode {
	res := [][]*GardenNode{}
	rows := strings.Split(data, "\n")
	for _, row := range rows {
		rv := []*GardenNode{}
		for _, val := range strings.Split(row, "") {
			rv = append(rv, &GardenNode{Value: val})
		}
		res = append(res, rv)
	}
	return res
}

func determineRegion(garden [][]*GardenNode, x, y, id int, val string) {
	garden[x][y].Region = id

	xp := x + 1
	xm := x - 1
	yp := y + 1
	ym := y - 1

	if xp < len(garden) && garden[xp][y].Value == val && garden[xp][y].Region == 0 {
		determineRegion(garden, xp, y, id, val)
	}

	if xm >= 0 && garden[xm][y].Value == val && garden[xm][y].Region == 0 {
		determineRegion(garden, xm, y, id, val)
	}

	if yp < len(garden[x]) && garden[x][yp].Value == val && garden[x][yp].Region == 0 {
		determineRegion(garden, x, yp, id, val)
	}

	if ym >= 0 && garden[x][ym].Value == val && garden[x][ym].Region == 0 {
		determineRegion(garden, x, ym, id, val)
	}
}

func contains(arr [][]int, val []int) bool {
	for _, a := range arr {
		if a[0] == val[0] && a[1] == val[1] {
			return true
		}
	}
	return false
}

func countCorners(regionIdexes [][]int, idx []int) int {
	l := []int{idx[0], idx[1] - 1}
	u := []int{idx[0] - 1, idx[1]}
	lud := []int{idx[0] - 1, idx[1] - 1}
	r := []int{idx[0], idx[1] + 1}
	urd := []int{idx[0] - 1, idx[1] + 1}
	d := []int{idx[0] + 1, idx[1]}
	drd := []int{idx[0] + 1, idx[1] + 1}
	dld := []int{idx[0] + 1, idx[1] - 1}

	counter := 0
	if !contains(regionIdexes, l) && !contains(regionIdexes, u) {
		counter++
	}
	if !contains(regionIdexes, u) && !contains(regionIdexes, r) {
		counter++
	}
	if !contains(regionIdexes, r) && !contains(regionIdexes, d) {
		counter++
	}
	if !contains(regionIdexes, d) && !contains(regionIdexes, l) {
		counter++
	}
	if contains(regionIdexes, l) && contains(regionIdexes, u) && !contains(regionIdexes, lud) {
		counter++
	}
	if contains(regionIdexes, u) && contains(regionIdexes, r) && !contains(regionIdexes, urd) {
		counter++
	}
	if contains(regionIdexes, r) && contains(regionIdexes, d) && !contains(regionIdexes, drd) {
		counter++
	}
	if contains(regionIdexes, d) && contains(regionIdexes, l) && !contains(regionIdexes, dld) {
		counter++
	}
	return counter
}

func getCost(garden [][]*GardenNode, region int) (int, int) {
	ridx := [][]int{}
	for x, row := range garden {
		for y, node := range row {
			if node.Region == region {
				ridx = append(ridx, []int{x, y})
			}
		}
	}

	fenceIdx := [][]int{}
	for _, idx := range ridx {
		if !contains(ridx, []int{idx[0] + 1, idx[1]}) {
			fenceIdx = append(fenceIdx, []int{idx[0] + 1, idx[1]})
		}
		if !contains(ridx, []int{idx[0] - 1, idx[1]}) {
			fenceIdx = append(fenceIdx, []int{idx[0] - 1, idx[1]})
		}
		if !contains(ridx, []int{idx[0], idx[1] + 1}) {
			fenceIdx = append(fenceIdx, []int{idx[0], idx[1] + 1})
		}
		if !contains(ridx, []int{idx[0], idx[1] - 1}) {
			fenceIdx = append(fenceIdx, []int{idx[0], idx[1] - 1})
		}
	}

	cornerCost := 0
	for _, entry := range ridx {
		cornerCost = cornerCost + countCorners(ridx, entry)
	}

	return len(fenceIdx) * len(ridx), cornerCost * len(ridx)
}

func Solve() error {
	datab, err := os.ReadFile("12/input1.txt")
	if err != nil {
		return fmt.Errorf("could not read data from file: %w", err)
	}

	datas := string(datab)
	garden := readGarden(datas)

	regionId := 1
	for x, row := range garden {
		for y, node := range row {
			if node.Region == 0 {
				determineRegion(garden, x, y, regionId, node.Value)
				regionId++
			}
		}
	}

	totalCost := 0
	discountedCost := 0
	for rgid := 1; rgid < regionId; rgid++ {
		t, d := getCost(garden, rgid)
		totalCost = totalCost + t
		discountedCost = discountedCost + d
	}

	fmt.Println("Result 1:", totalCost)
	fmt.Println("Result 2:", discountedCost)
	return nil
}
