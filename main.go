package main

import (
	"fmt"

	no1 "github.com/ryshoooo/aoc2024/1"
	no10 "github.com/ryshoooo/aoc2024/10"
	no11 "github.com/ryshoooo/aoc2024/11"
	no12 "github.com/ryshoooo/aoc2024/12"
	no2 "github.com/ryshoooo/aoc2024/2"
	no3 "github.com/ryshoooo/aoc2024/3"
	no4 "github.com/ryshoooo/aoc2024/4"
	no5 "github.com/ryshoooo/aoc2024/5"
	no6 "github.com/ryshoooo/aoc2024/6"
	no7 "github.com/ryshoooo/aoc2024/7"
	no8 "github.com/ryshoooo/aoc2024/8"
	no9 "github.com/ryshoooo/aoc2024/9"
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
	printPretty(5, no5.Solve)
	printPretty(6, no6.Solve)
	printPretty(7, no7.Solve)
	printPretty(8, no8.Solve)
	printPretty(9, no9.Solve)
	printPretty(10, no10.Solve)
	printPretty(11, no11.Solve)
	printPretty(12, no12.Solve)
}
