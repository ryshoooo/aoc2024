package main

import (
	"fmt"

	no1 "github.com/ryshoooo/aoc2024/1"
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
}
