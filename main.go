package main

import (
	"fmt"

	no14 "github.com/ryshoooo/aoc2024/14"
)

func printPretty(no int, solver func() error) {
	fmt.Println("Day", no)
	fmt.Println("==========")
	err := solver()
	if err != nil {
		panic(err)
	}
	fmt.Println("==========")
}

func main() {
	// printPretty(1, no1.Solve)
	// printPretty(2, no2.Solve)
	// printPretty(3, no3.Solve)
	// printPretty(4, no4.Solve)
	// printPretty(5, no5.Solve)
	// printPretty(6, no6.Solve)
	// printPretty(7, no7.Solve)
	// printPretty(8, no8.Solve)
	// printPretty(9, no9.Solve)
	// printPretty(10, no10.Solve)
	// printPretty(11, no11.Solve)
	// printPretty(12, no12.Solve)
	// printPretty(13, no13.Solve)
	printPretty(14, no14.Solve)
}
