package day4

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Q1() {
	file, err := os.Open("day4/main.txt")
	if err != nil {
		log.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	finalAnswer := 0

	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, ": ")[1]
		winningNumbersStr := strings.Split(strings.Split(numbers, " | ")[0], " ")
		numbersIHaveStr := strings.Split(strings.Split(numbers, " | ")[1], " ")
		winningNumbers := []int{}
		numbersIHave := []int{}
		// log.Println(winningNumbersStr, numbersIHaveStr)
		for _, num := range winningNumbersStr {
			if num == "" {
				continue
			}
			val, err := strconv.Atoi(num)
			if err != nil {
				log.Println("error converting string to int", err)
				return
			}
			winningNumbers = append(winningNumbers, val)
		}
		for _, num := range numbersIHaveStr {
			if num == "" {
				continue
			}
			val, err := strconv.Atoi(num)
			if err != nil {
				log.Println("error converting string to int", err)
				return
			}
			numbersIHave = append(numbersIHave, val)
		}
		// log.Println(winningNumbers, numbersIHave)

		// count how many numbersIHave are in winningNumbers
		count := 0
		for _, num := range numbersIHave {
			for _, winningNum := range winningNumbers {
				if num == winningNum {
					count++
				}
			}
		}
		if count > 0 {
			finalAnswer += power(2, count-1)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading file:", err)
	}
	log.Println("Final answer:", finalAnswer)
}

func power(base, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}

type scratchCard struct {
	gameNum            int
	matchingCardsCount int
}

func Q2() {

	file, err := os.Open("day4/main.txt")
	if err != nil {
		log.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	finalAnswer := 0

	queue := []scratchCard{}

	winningNumbersStore := make(map[int]int)

	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, ": ")[1]
		gameInfo := strings.Split(strings.Split(line, ": ")[0], " ")
		log.Println(gameInfo)
		// remove all empty strings from gameInfo
		for i := 0; i < len(gameInfo); i++ {
			if gameInfo[i] == "" {
				gameInfo = append(gameInfo[:i], gameInfo[i+1:]...)
				i--
			}
		}
		gameNum, err := strconv.Atoi(gameInfo[1])
		if err != nil {
			log.Println("error converting string to int", err)
			return
		} else {
			log.Println("gameNum:", gameNum)
		}
		winningNumbersStr := strings.Split(strings.Split(numbers, " | ")[0], " ")
		numbersIHaveStr := strings.Split(strings.Split(numbers, " | ")[1], " ")
		winningNumbers := []int{}
		numbersIHave := []int{}
		log.Println(winningNumbersStr, numbersIHaveStr)
		for _, num := range winningNumbersStr {
			if num == "" {
				continue
			}
			val, err := strconv.Atoi(num)
			if err != nil {
				log.Println("error converting string to int", err)
				return
			}
			winningNumbers = append(winningNumbers, val)
		}
		log.Println(winningNumbers)
		for _, num := range numbersIHaveStr {
			if num == "" {
				continue
			}
			val, err := strconv.Atoi(num)
			if err != nil {
				log.Println("error converting string to int", err)
				return
			}
			numbersIHave = append(numbersIHave, val)
		}
		log.Println(winningNumbers, numbersIHave)

		// count how many numbersIHave are in winningNumbers
		count := 0
		for _, num := range numbersIHave {
			for _, winningNum := range winningNumbers {
				if num == winningNum {
					count++
				}
			}
		}
		queue = append(queue, scratchCard{gameNum, count})
		winningNumbersStore[gameNum] = count
	}

	// while queue is not empty
	total := 0
	for len(queue) > 0 {
		// pop from queue
		card := queue[0]
		total++
		queue = queue[1:]
		for i := card.gameNum + 1; i <= winningNumbersStore[card.gameNum]+card.gameNum; i++ {
			queue = append(queue, scratchCard{i, winningNumbersStore[i]})
		}
	}
	finalAnswer = total

	if err := scanner.Err(); err != nil {
		log.Println("Error reading file:", err)
	}
	log.Println("Final answer:", finalAnswer)
}
