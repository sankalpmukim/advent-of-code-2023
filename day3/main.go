package day3

import (
	"bufio"
	"log"
	"os"
	"unicode"
)

func Q1() {
	file, err := os.Open("day3/dummyinput.txt")
	if err != nil {
		log.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	finalAnswer := 0

	lines := [][]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, []rune(line))
	}

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if unicode.IsDigit(lines[i][j]) {
				// iterate and find the index of ending digit to get the whole number
				endIndex := j
				for k := j + 1; k < len(lines[i]); k++ {
					if !unicode.IsDigit(lines[i][k]) {
						break
					}
					endIndex = k
				}
				// convert the string to int
				num := 0
				for k := j; k <= endIndex; k++ {
					num = num*10 + int(lines[i][k]-'0')
				}
				// log.Println("Found number:", num)
				if isPartNumber(lines, i, j, endIndex) {
					finalAnswer += num
					// } else {
					// log.Println("Not part of a number:", num)
				}
				j = endIndex + 1
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading file:", err)
	}
	log.Println("Final answer:", finalAnswer)
}

func isPartNumber(lines [][]rune, i, j, k int) bool {
	// check i-1th to i+1th row
	for l := i - 1; l <= i+1; l++ {
		// check j-1th to k+1th column
		for m := j - 1; m <= k+1; m++ {
			if l < 0 || l >= len(lines) || m < 0 || m >= len(lines[l]) {
				continue
			}
			if l == i && m >= j && m <= k {
				continue
			}
			if !unicode.IsDigit(lines[l][m]) && lines[l][m] != '.' {
				return true
			}
		}
	}
	return false
}

func isAdjacentToGear(lines [][]rune, i, j, k int) (bool, gearPosition) {
	// gear is '*' symbol
	for l := i - 1; l <= i+1; l++ {
		// check j-1th to k+1th column
		for m := j - 1; m <= k+1; m++ {
			if l < 0 || l >= len(lines) || m < 0 || m >= len(lines[l]) {
				continue
			}
			if l == i && m >= j && m <= k {
				continue
			}
			if lines[l][m] == '*' {
				return true, gearPosition{l, m}
			}
		}
	}
	return false, gearPosition{}
}

type gearPosition struct {
	x int
	y int
}

func Q2() {
	file, err := os.Open("day3/main.txt")
	if err != nil {
		log.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	finalAnswer := 0

	lines := [][]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, []rune(line))
	}

	partNumsNearGear := make(map[gearPosition][]int)

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if unicode.IsDigit(lines[i][j]) {
				endIndex := j
				for k := j + 1; k < len(lines[i]); k++ {
					if !unicode.IsDigit(lines[i][k]) {
						break
					}
					endIndex = k
				}
				num := 0
				for k := j; k <= endIndex; k++ {
					num = num*10 + int(lines[i][k]-'0')
				}
				if isPartNumber(lines, i, j, endIndex) {
					isNearGear, gearPos := isAdjacentToGear(lines, i, j, endIndex)
					if isNearGear {
						partNumsNearGear[gearPos] = append(partNumsNearGear[gearPos], num)
						// } else {
						log.Println("Not near gear:", num)
					}
				}
				j = endIndex + 1
			}
		}
	}

	for k, nums := range partNumsNearGear {
		log.Println(k, nums)
		if len(nums) == 2 {
			finalAnswer += nums[0] * nums[1]
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading file:", err)
	}
	log.Println("Final answer:", finalAnswer)
}
