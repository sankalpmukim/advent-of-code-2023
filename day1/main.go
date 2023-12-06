package day1

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"unicode"
)

func Q1() {
	file, err := os.Open("day1/main.txt")
	if err != nil {
		log.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		var num int
		for _, r := range line {
			// find first digit
			if unicode.IsDigit(r) {
				num = int(r - '0')
				break
			}
		}
		runes := []rune(line)

		for i := len(runes) - 1; i >= 0; i-- {
			if unicode.IsDigit(runes[i]) {
				num *= 10
				num += int(runes[i] - '0')
				break
			}
		}
		sum += num
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading file:", err)
	}
	log.Println(sum)
}

func Q2() {
	file, err := os.Open("day1/main.txt")
	if err != nil {
		log.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		// log.Println(line)
		firstDigit, lastDigit := getDigits(line)
		// log.Println("firstDigit", firstDigit, "line", line)
		// log.Println("lastDigit", lastDigit, "line", line)
		sum += firstDigit*10 + lastDigit
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading file:", err)
	}
	log.Println(sum)
}

type IndexPair struct {
	minIndex int
	maxIndex int
}

func getDigits(line string) (int, int) {
	stringIndexes := make([]IndexPair, 0)
	numIndexes := make([]IndexPair, 0)

	// log.Println("length", len(stringIndexes))

	numStrings := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	numStringsMap := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for i := 0; i < 10; i++ {
		stringIndexes = append(stringIndexes,
			IndexPair{
				minIndex: strings.Index(line, numStrings[i]),
				maxIndex: strings.LastIndex(line, numStrings[i]),
			})
		numIndexes = append(numIndexes, IndexPair{
			minIndex: strings.Index(line, fmt.Sprint(i)),
			maxIndex: strings.LastIndex(line, fmt.Sprint(i)),
		})
	}

	// log.Println(stringIndexes)
	// log.Println(numIndexes)

	stringMin := math.MaxInt
	numMin := math.MaxInt
	stringMax := math.MinInt
	numMax := math.MinInt
	for i := 0; i < 10; i++ {
		if stringIndexes[i].minIndex != -1 && stringIndexes[i].minIndex < stringMin {
			stringMin = stringIndexes[i].minIndex
		}

		if numIndexes[i].minIndex != -1 && numIndexes[i].minIndex < numMin {
			numMin = numIndexes[i].minIndex
		}

		if stringIndexes[i].maxIndex != -1 && stringIndexes[i].maxIndex > stringMax {
			stringMax = stringIndexes[i].maxIndex
		}

		if numIndexes[i].maxIndex != -1 && numIndexes[i].maxIndex > numMax {
			numMax = numIndexes[i].maxIndex
		}
	}

	var rMin, rMax int

	if numMin != -1 && numMin != math.MaxInt && numMin < stringMin {
		rMin = int(line[numMin] - '0')
	} else {
		for stringKey := range numStringsMap {
			if strings.Index(line, stringKey) == stringMin {
				rMin = numStringsMap[stringKey]
			}
		}
	}

	if numMax != -1 && numMax != math.MinInt && numMax > stringMax {
		rMax = int(line[numMax] - '0')
	} else {
		for stringKey := range numStringsMap {
			if strings.LastIndex(line, stringKey) == stringMax {
				rMax = numStringsMap[stringKey]
			}
		}
	}

	return rMin, rMax
}
