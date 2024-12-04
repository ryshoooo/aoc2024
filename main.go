package main

import (
	"fmt"

	no1 "github.com/ryshoooo/aoc2024/1"
	no2 "github.com/ryshoooo/aoc2024/2"
	no3 "github.com/ryshoooo/aoc2024/3"
	no4 "github.com/ryshoooo/aoc2024/4"
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
	printPretty(1, no1.Solve)
	printPretty(2, no2.Solve)
	printPretty(3, no3.Solve)
	printPretty(4, no4.Solve)
}
