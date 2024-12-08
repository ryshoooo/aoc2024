package no8

import (
	"fmt"
	"os"
	"strings"
)

func Solve() error {
	datab, err := os.ReadFile("8/input1.txt")
	if err != nil {
		return fmt.Errorf("could not read file: %v", err)
	}

	datas := string(datab)
	antennas := make(map[string][][]int)
	xbound := len(strings.Split(datas, "\n"))
	ybound := len(strings.Split(datas, "\n")[0])
	for x, line := range strings.Split(datas, "\n") {
		for y, val := range strings.Split(line, "") {
			if val != "." {
				antennas[val] = append(antennas[val], []int{x, y})
			}
		}
	}

	antinodes := make(map[string]int)
	for _, positions := range antennas {
		for ind, p1 := range positions {
			for _, p2 := range positions[ind+1:] {
				vec := []int{p2[0] - p1[0], p2[1] - p1[1]}
				vec180 := []int{-vec[0], -vec[1]}
				a1 := []int{p2[0] + vec[0], p2[1] + vec[1]}
				a2 := []int{p1[0] + vec180[0], p1[1] + vec180[1]}

				if a1[0] >= 0 && a1[0] < xbound && a1[1] >= 0 && a1[1] < ybound {
					antinodes[fmt.Sprintf("%d-%d", a1[0], a1[1])]++
				}
				if a2[0] >= 0 && a2[0] < xbound && a2[1] >= 0 && a2[1] < ybound {
					antinodes[fmt.Sprintf("%d-%d", a2[0], a2[1])]++
				}
			}
		}
	}

	fmt.Println("Result 1:", len(antinodes))

	antinodes2 := make(map[string]int)
	for _, positions := range antennas {
		for ind, p1 := range positions {
			for _, p2 := range positions[ind+1:] {
				vec := []int{p2[0] - p1[0], p2[1] - p1[1]}
				vec180 := []int{-vec[0], -vec[1]}

				antinodes2[fmt.Sprintf("%d-%d", p1[0], p1[1])]++
				antinodes2[fmt.Sprintf("%d-%d", p2[0], p2[1])]++

				as := []int{p2[0], p2[1]}
				for {
					a := []int{as[0] + vec[0], as[1] + vec[1]}
					if a[0] >= 0 && a[0] < xbound && a[1] >= 0 && a[1] < ybound {
						antinodes2[fmt.Sprintf("%d-%d", a[0], a[1])]++
						as[0] = a[0]
						as[1] = a[1]
					} else {
						break
					}
				}

				bs := []int{p1[0], p1[1]}
				for {
					b := []int{bs[0] + vec180[0], bs[1] + vec180[1]}
					if b[0] >= 0 && b[0] < xbound && b[1] >= 0 && b[1] < ybound {
						antinodes2[fmt.Sprintf("%d-%d", b[0], b[1])]++
						bs[0] = b[0]
						bs[1] = b[1]
					} else {
						break
					}
				}
			}
		}
	}
	fmt.Println("Result 2:", len(antinodes2))
	return nil
}
