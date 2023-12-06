package day6

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Q1() {
	bytesContent, err := os.ReadFile("day6/main.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	content := strings.Split(strings.TrimSpace(string(bytesContent)), "\n")
	timesStr := strings.Split(strings.TrimSpace(strings.Split(content[0], ":")[1]), " ")
	distsStr := strings.Split(strings.TrimSpace(strings.Split(content[1], ":")[1]), " ")
	times := make([]int, 0)
	for i := 0; i < len(timesStr); i++ {
		if timesStr[i] == "" {
			continue
		}
		intVal, err := strconv.Atoi(timesStr[i])
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			return
		}
		times = append(times, intVal)
	}
	dists := make([]int, 0)
	for i := 0; i < len(distsStr); i++ {
		if distsStr[i] == "" {
			continue
		}
		intVal, err := strconv.Atoi(distsStr[i])
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			return
		}
		dists = append(dists, intVal)
	}
	finalAnswer := 1
	for i := 0; i < len(times); i++ {
		for x := 1; x <= times[i]-1; x++ {
			realDist := x * (times[i] - x)
			if realDist > dists[i] {
				maxPossible := times[i] - x*2 + 1
				fmt.Println("Max Possible:", maxPossible)
				finalAnswer *= maxPossible
				break
			}
		}
	}
	fmt.Println("Final Answer:", finalAnswer)
}

func Q2() {
	bytesContent, err := os.ReadFile("day6/main.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	content := strings.Split(strings.TrimSpace(string(bytesContent)), "\n")
	timesStr := strings.Split(strings.TrimSpace(strings.Split(content[0], ":")[1]), " ")
	distsStr := strings.Split(strings.TrimSpace(strings.Split(content[1], ":")[1]), " ")
	times := make([]int, 0)
	for i := 0; i < len(timesStr); i++ {
		if timesStr[i] == "" {
			continue
		}
		intVal, err := strconv.Atoi(timesStr[i])
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			return
		}
		times = append(times, intVal)
	}
	dists := make([]int, 0)
	for i := 0; i < len(distsStr); i++ {
		if distsStr[i] == "" {
			continue
		}
		intVal, err := strconv.Atoi(distsStr[i])
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			return
		}
		dists = append(dists, intVal)
	}
	timesCum := ""
	distsCum := ""
	for i := 0; i < len(times); i++ {
		timesCum += strconv.Itoa(times[i])
		distsCum += strconv.Itoa(dists[i])
	}
	timesInt, err := strconv.Atoi(timesCum)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return
	}
	distsInt, err := strconv.Atoi(distsCum)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return
	}
	fmt.Println("Times:", timesInt)
	fmt.Println("Dists:", distsInt)
	times = []int{timesInt}
	dists = []int{distsInt}
	finalAnswer := 1
	for i := 0; i < len(times); i++ {
		for x := 1; x <= times[i]-1; x++ {
			realDist := x * (times[i] - x)
			if realDist > dists[i] {
				maxPossible := times[i] - x*2 + 1
				fmt.Println("Max Possible:", maxPossible)
				finalAnswer *= maxPossible
				break
			}
		}
	}
	fmt.Println("Final Answer:", finalAnswer)
}
