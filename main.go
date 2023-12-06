package main

import (
	"fmt"
	"time"

	"github.com/sankalpmukim/aoc-2023-go/day6"
)

func main() {
	// start stopwatch
	startTime := time.Now()
	day6.Q2()
	elapsedTime := time.Since(startTime)
	fmt.Println("Time elapsed:", elapsedTime)
}
