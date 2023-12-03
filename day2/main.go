package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Q1() {
	file, err := os.Open("day2/main.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	possibleCount := 0
	maxRanges := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	i := 1
	for scanner.Scan() {

		// ["3 blue, 4 red","1 red, 2 green"]
		gameSets := strings.Split(
			strings.Split(
				scanner.Text(),
				": ",
			)[1],
			"; ",
		)

		isValidFlag := true
		for _, gameSet := range gameSets {
			// ["3 blue", "4 red"]
			isInnerValid := true
			games := strings.Split(gameSet, ", ")
			for _, game := range games {
				// "3 blue"
				gameInfo := strings.Split(game, " ")
				count, err := strconv.Atoi(gameInfo[0])
				if err != nil {
					fmt.Println("Error converting string to int:", err)
					return
				}
				if count > maxRanges[gameInfo[1]] {
					isInnerValid = false
					break
				}
			}
			if !isInnerValid {
				isValidFlag = false
				break
			}
		}
		if isValidFlag {
			possibleCount += i
		}
		i += 1
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println(possibleCount)
}

func Q2() {
	file, err := os.Open("day2/main.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	finalSum := 0

	for scanner.Scan() {
		// ["3 blue, 4 red","1 red, 2 green"]
		games := strings.Split(
			strings.Split(
				scanner.Text(),
				": ",
			)[1],
			"; ",
		)
		// initialize min counts
		minCounts := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		for _, game := range games {

			drops := strings.Split(game, ", ")
			for _, drop := range drops {
				gameInfo := strings.Split(drop, " ")
				count, err := strconv.Atoi(gameInfo[0])
				if err != nil {
					fmt.Println("Error converting string to int:", err)
				}
				if count > minCounts[gameInfo[1]] {
					minCounts[gameInfo[1]] = count
				}
			}
		}
		power := 1
		for _, minCount := range minCounts {
			if minCount != 0 {
				power *= minCount
			}
		}
		finalSum += power
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println(finalSum)
}
