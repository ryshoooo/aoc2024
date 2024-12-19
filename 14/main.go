package no14

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Robot struct {
	X  int
	Y  int
	Vx int
	Vy int
}

func (r *Robot) Move(xlim, ylim int) {
	r.X = (r.X + r.Vx) % xlim
	r.Y = (r.Y + r.Vy) % ylim
	if r.X < 0 {
		r.X = (r.X + xlim) % xlim
	}
	if r.Y < 0 {
		r.Y = (r.Y + ylim) % ylim
	}
}

func (r *Robot) GetQuadrant(xlim, ylim int) string {
	midx := (xlim - 1) / 2
	midy := (ylim - 1) / 2

	if r.X < midx && r.Y < midy {
		return "UL"
	}

	if r.X > midx && r.Y < midy {
		return "UR"
	}

	if r.X < midx && r.Y > midy {
		return "LL"
	}

	if r.X > midx && r.Y > midy {
		return "LR"
	}

	return "none"

}

func readRobots(data string) []*Robot {
	rr := strings.Split(data, "\n")
	res := make([]*Robot, len(rr))

	for ind, r := range rr {
		pv := strings.Split(r, " ")
		ps := strings.Split(strings.Split(pv[0], "=")[1], ",")
		x, _ := strconv.Atoi(ps[0])
		y, _ := strconv.Atoi(ps[1])
		vs := strings.Split(strings.Split(pv[1], "=")[1], ",")
		vx, _ := strconv.Atoi(vs[0])
		vy, _ := strconv.Atoi(vs[1])
		res[ind] = &Robot{X: x, Y: y, Vx: vx, Vy: vy}
	}
	return res
}

func contains(robots []*Robot, x, y int) bool {
	for _, r := range robots {
		if r.X == x && r.Y == y {
			return true
		}
	}
	return false
}

func maxConsecutive(robots []*Robot) int {
	max := 0
	for _, r := range robots {
		consecutive := 0
		for i := 1; i <= 30; i++ {
			if contains(robots, r.X, r.Y+i) {
				consecutive++
			} else {
				break
			}
		}
		if consecutive > max {
			max = consecutive
		}
	}
	return max
}

func Solve() error {
	xlim := 101
	ylim := 103
	// xlim := 11
	// ylim := 7
	datab, err := os.ReadFile("14/input1.txt")
	if err != nil {
		return fmt.Errorf("could not read file: %w", err)
	}

	datas := string(datab)
	robots := readRobots(datas)
	quadsRes := make(map[string]int)
	for _, robot := range robots {
		for range 100 {
			robot.Move(xlim, ylim)
		}
		quad := robot.GetQuadrant(xlim, ylim)
		if quad != "none" {
			quadsRes[quad]++
		}
	}

	total := 1
	for _, v := range quadsRes {
		total *= v
	}
	fmt.Println("Result 1:", total)

	robots = readRobots(datas)
	it := 1
	for {
		for _, robot := range robots {
			robot.Move(xlim, ylim)
		}
		m := maxConsecutive(robots)
		if m >= 30 {
			fmt.Println("Result 2:", it)
			return nil
		}
		it++
	}
}
