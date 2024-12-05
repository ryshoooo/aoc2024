package no5

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func makeOrderMap(orders []string) (map[int][]int, error) {
	unres := make(map[int][]int)

	for _, o := range orders {
		pages := strings.Split(o, "|")
		lh := pages[0]
		rg := pages[1]

		lhi, err := strconv.Atoi(lh)
		if err != nil {
			return unres, fmt.Errorf("could not convert %s to int: %w", lh, err)
		}

		rhi, err := strconv.Atoi(rg)
		if err != nil {
			return unres, fmt.Errorf("could not convert %s to int: %w", rg, err)
		}

		unres[rhi] = append(unres[rhi], lhi)
	}
	return unres, nil
}

func parseUpdate(update string) ([]int, error) {
	vals := strings.Split(update, ",")
	res := make([]int, len(vals))
	for ind, v := range vals {
		vali, err := strconv.Atoi(v)
		if err != nil {
			return res, fmt.Errorf("could not convert %s to int: %w", v, err)
		}
		res[ind] = vali
	}
	return res, nil
}

func contains(arr []int, val int) bool {
	for _, x := range arr {
		if val == x {
			return true
		}
	}
	return false
}

func getValidMiddle(update []int, notAllowed map[int][]int) int {
	for ind, uv := range update {
		valsToCheck := update[ind+1:]
		for _, vtc := range valsToCheck {
			if contains(notAllowed[uv], vtc) {
				return 0
			}
		}
	}

	return update[len(update)/2]
}

func fixUpdate(update []int, notAllowed map[int][]int) int {
	cidx := 0
	for {
		if len(update) == cidx+1 {
			break
		}
		val := update[cidx]
		valsToCheck := update[cidx+1:]
		tryAgain := false
		for _idx, vtc := range valsToCheck {
			if contains(notAllowed[val], vtc) {
				update[cidx] = vtc
				update[_idx+cidx+1] = val
				tryAgain = true
				break
			}
		}
		if tryAgain {
			continue
		}
		cidx++
	}
	return update[len(update)/2]
}

func Solve() error {
	datab, err := os.ReadFile("5/input1.txt")
	if err != nil {
		return fmt.Errorf("could not read file: %w", err)
	}

	data := string(datab)
	alld := strings.Split(data, "\n")
	sidx := 0
	for ind, d := range alld {
		if len(d) == 0 {
			sidx = ind
			break
		}
	}
	orders := alld[:sidx]
	updates := alld[sidx+1:]

	notAllowedOrderMap, err := makeOrderMap(orders)
	if err != nil {
		return fmt.Errorf("could not make order map: %w", err)
	}

	res := 0
	res2 := 0
	for _, update := range updates {
		us, err := parseUpdate(update)
		if err != nil {
			return fmt.Errorf("could not parse update: %w", err)
		}

		mid := getValidMiddle(us, notAllowedOrderMap)
		if mid == 0 {
			res2 += fixUpdate(us, notAllowedOrderMap)
		}
		res += mid
	}

	fmt.Println("Result 1:", res)
	fmt.Println("Result 2:", res2)
	return nil
}
