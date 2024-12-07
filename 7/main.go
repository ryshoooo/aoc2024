package no7

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseEq(eq string) (int, []int, error) {
	eqc := strings.Split(eq, ":")
	res, err := strconv.Atoi(eqc[0])
	if err != nil {
		return 0, []int{}, fmt.Errorf("could not parse result %s: %v", eqc[0], err)
	}

	vals := make([]int, 0)
	for _, v := range strings.Split(eqc[1], " ") {
		if v == "" {
			continue
		}
		vi, err := strconv.Atoi(v)
		if err != nil {
			return 0, []int{}, fmt.Errorf("could not parse value %s: %v", v, err)
		}
		vals = append(vals, vi)
	}

	return res, vals, nil
}

func isPossible(target int, start int, followups []int, inclConcat bool) bool {
	if len(followups) == 0 {
		return start == target
	}
	if start > target {
		return false
	}
	if isPossible(target, start+followups[0], followups[1:], inclConcat) {
		return true
	} else if isPossible(target, start*followups[0], followups[1:], inclConcat) {
		return true
	} else if inclConcat {
		v, err := strconv.Atoi(fmt.Sprintf("%d%d", start, followups[0]))
		if err != nil {
			panic("failed to convert concatenated string to int" + err.Error())
		}
		return isPossible(target, v, followups[1:], inclConcat)
	} else {
		return false
	}
}

func Solve() error {
	datab, err := os.ReadFile("7/input1.txt")
	if err != nil {
		return fmt.Errorf("could not read file: %v", err)
	}

	datas := string(datab)
	eqs := strings.Split(datas, "\n")
	res1 := 0
	res2 := 0
	for _, eq := range eqs {
		res, ins, err := parseEq(eq)
		if err != nil {
			return fmt.Errorf("could not parse equation %s: %v", eq, err)
		}
		if isPossible(res, ins[0], ins[1:], false) {
			res1 += res
		}
		if isPossible(res, ins[0], ins[1:], true) {
			res2 += res
		}
	}
	fmt.Println("Result 1: ", res1)
	fmt.Println("Result 2: ", res2)
	return nil
}
