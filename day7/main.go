package day7

import (
	"bufio"
	"log"
	"os"
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
		log.Println(line)
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error scanning file:", err)
		return
	}
	log.Println("Final Answer:", finalAnswer)

}
