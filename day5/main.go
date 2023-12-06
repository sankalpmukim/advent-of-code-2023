package day5

import (
	"log"
	"os"
	"strconv"
	str "strings"
)

func Q1() {
	bytesContent, err := os.ReadFile("day5/main.txt")
	if err != nil {
		log.Println("Error opening file:", err)
		return
	}
	content := str.Split(string(bytesContent), "\n\n")
	stringSeeds := str.Split(str.Split(content[0], ": ")[1], " ")
	// remove empty strings from seeds
	for i := 0; i < len(stringSeeds); i++ {
		if stringSeeds[i] == "" {
			stringSeeds = append(stringSeeds[:i], stringSeeds[i+1:]...)
			i--
		}
	}
	// get integer seeds
	seeds := make([]int, len(stringSeeds))
	for i := 0; i < len(stringSeeds); i++ {
		val, err := strconv.Atoi(stringSeeds[i])
		if err != nil {
			log.Println("Error converting string to int:", err)
			return
		}
		seeds[i] = val
	}
	log.Println("Seeds:", seeds)

	finalAnswer := 0

	allCurrVals := make([]int, 0, len(seeds))

	for i := 0; i < len(seeds); i++ {
		// iterate content from 1 to len(content) - 1
		currVal := seeds[i]
		log.Println("num content: ", len(content))
		for j := 1; j < len(content); j++ {
			// lines split by \n
			lines := str.Split(str.TrimSpace(content[j]), "\n")[1:]
			// sourceToDest := strings.Split(strings.Split(lines[0], " map:")[0], "-to-")
			// source := sourceToDest[0]
			// dest := sourceToDest[1]
			// log.Println("Source:", source, "\t\tDest:", dest)
			for k := 0; k < len(lines); k++ {
				dest, src, rnge, err := parseMapRow(lines[k])
				if err != nil {
					log.Println("Error parsing map row:", err)
					return
				}
				// log.Println("Dest:", dest, "\t\tSrc:", src, "\t\tRnge:", rnge)
				if currVal >= src && currVal < src+rnge {
					currVal = dest + (currVal - src)
					log.Println("currVal:", currVal)
					break
				}
			}
		}
		allCurrVals = append(allCurrVals, currVal)
	}

	finalAnswer = allCurrVals[0]
	log.Println("allCurrVals:", allCurrVals, "\t\tlen(allCurrVals):",
		len(allCurrVals), "\t\tfinalANswer:", finalAnswer)
	for i := 1; i < len(allCurrVals); i++ {
		if allCurrVals[i] < finalAnswer {
			finalAnswer = allCurrVals[i]
		}
	}

	log.Println("Final answer:", finalAnswer)
}

func parseMapRow(line string) (dest int, src int, rnge int, err error) {
	splits := str.Split(line, " ")
	// remove empty strings from splits
	for i := 0; i < len(splits); i++ {
		if splits[i] == "" {
			splits = append(splits[:i], splits[i+1:]...)
			i--
		}
	}
	// get dest
	dest, err = strconv.Atoi(splits[0])
	if err != nil {
		return
	}
	// get src
	src, err = strconv.Atoi(splits[1])
	if err != nil {
		return
	}
	// get rnge
	rnge, err = strconv.Atoi(splits[2])
	if err != nil {
		return
	}
	return
}

func Q2() {
	bytesContent, err := os.ReadFile("day5/dummy.txt")
	if err != nil {
		log.Println("Error opening file:", err)
		return
	}
	content := str.Split(string(bytesContent), "\n\n")
	stringSeedRanges := str.Split(str.Split(content[0], ": ")[1], " ")
	// remove empty strings from seeds
	for i := 0; i < len(stringSeedRanges); i++ {
		if stringSeedRanges[i] == "" {
			stringSeedRanges = append(stringSeedRanges[:i], stringSeedRanges[i+1:]...)
			i--
		}
	}
	// get integer seedRanges
	seedRanges := make([]int, len(stringSeedRanges))
	for i := 0; i < len(stringSeedRanges); i++ {
		val, err := strconv.Atoi(stringSeedRanges[i])
		if err != nil {
			log.Println("Error converting string to int:", err)
			return
		}
		seedRanges[i] = val
	}
	log.Println("Seeds Ranges:", seedRanges)

	seeds := make([]int, 0)
	for i := 0; i < len(seedRanges); i += 2 {
		for j := seedRanges[i]; j <= seedRanges[i]+seedRanges[i+1]; j++ {
			seeds = append(seeds, j)
		}
	}

	finalAnswer := 0

	allCurrVals := make([]int, 0, len(seedRanges))
	log.Println("Seeds:", seeds)

	for i := 0; i < len(seeds); i++ {
		// iterate content from 1 to len(content) - 1
		currVal := seeds[i]
		log.Println("num content: ", len(content))
		for j := 1; j < len(content); j++ {
			// lines split by \n
			lines := str.Split(str.TrimSpace(content[j]), "\n")[1:]
			// sourceToDest := strings.Split(strings.Split(lines[0], " map:")[0], "-to-")
			// source := sourceToDest[0]
			// dest := sourceToDest[1]
			// log.Println("Source:", source, "\t\tDest:", dest)
			for k := 0; k < len(lines); k++ {
				dest, src, rnge, err := parseMapRow(lines[k])
				if err != nil {
					log.Println("Error parsing map row:", err)
					return
				}
				// log.Println("Dest:", dest, "\t\tSrc:", src, "\t\tRnge:", rnge)
				if currVal >= src && currVal < src+rnge {
					currVal = dest + (currVal - src)
					log.Println("currVal:", currVal)
					break
				}
			}
		}
		allCurrVals = append(allCurrVals, currVal)
	}

	finalAnswer = allCurrVals[0]
	log.Println("allCurrVals:", allCurrVals, "\t\tlen(allCurrVals):",
		len(allCurrVals), "\t\tfinalANswer:", finalAnswer)
	for i := 1; i < len(allCurrVals); i++ {
		if allCurrVals[i] < finalAnswer {
			finalAnswer = allCurrVals[i]
		}
	}

	log.Println("Final answer:", finalAnswer)
}
