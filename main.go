package main

import (
	"fmt"
	"time"

	"github.com/sankalpmukim/aoc-2023-go/day5"
)

func main() {
	// start stopwatch
	startTime := time.Now()
	// day5.Q1()
	day5.Q2()
	elapsedTime := time.Since(startTime)
	fmt.Println("Time elapsed:", elapsedTime)
}
