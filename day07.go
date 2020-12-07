package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Requirements struct {
	count     int
	bagColour string
}

func totalBagsInside(requirementsByColour map[string][]Requirements,
	bagColour string) int {
	reqs, ok := requirementsByColour[bagColour]
	if !ok {
		return 0
	} else {
		// how many inside this bag
		totalInside := 0
		for _, req := range reqs {
			totalInside += req.count
			totalInside += req.count * (totalBagsInside(requirementsByColour, req.bagColour))
		}
		return totalInside
	}
}

func bagColoursInside(requirementsByColour map[string][]Requirements,
	bagColour string) map[string]bool {
	reqs, ok := requirementsByColour[bagColour]
	if !ok {
		return make(map[string]bool)
	} else {
		directReqs := make(map[string]bool)
		for _, req := range reqs {
			directReqs[req.bagColour] = true
		}
		totalReqs := make(map[string]bool)
		for directColour, _ := range directReqs {
			totalReqs[directColour] = true
			indirectReqsForThatReq := bagColoursInside(requirementsByColour, directColour)
			for indirectCol, _ := range indirectReqsForThatReq {
				totalReqs[indirectCol] = true
			}
		}
		return totalReqs
	}
}

func readBagsInfo() map[string][]Requirements {
	file, err := os.Open("C:\\Workarea\\advent-2020-go\\inputs\\day07.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	requirementsByBagColour := make(map[string][]Requirements, 0)

	for scanner.Scan() {
		line := scanner.Text()
		line = line[0 : len(line)-1] // remove the .
		objSubj := strings.Split(line, " bags contain ")
		bagColour := objSubj[0]

		requirements := make([]Requirements, 0)

		if objSubj[1] != "no other bags" {

			bagContents := strings.Split(objSubj[1], ", ")
			for _, bagContent := range bagContents {

				// find first space
				firstSpaceIsAt := strings.Index(bagContent, " ")
				countRequired, _ := strconv.Atoi(bagContent[:firstSpaceIsAt])
				colourRequired := bagContent[firstSpaceIsAt+1:]
				colourRequired = strings.Replace(colourRequired, " bags", "", -1)
				colourRequired = strings.Replace(colourRequired, " bag", "", -1)
				requirements = append(requirements, Requirements{countRequired, colourRequired})
			}
		}
		requirementsByBagColour[bagColour] = requirements
	}
	return requirementsByBagColour
}

func main() {

	bagsInfo := readBagsInfo()
	fmt.Println("Read total of " + strconv.Itoa(len(bagsInfo)))

	// It would have felt more satisfying to start with the shiny gold bag and work out
	// but never mind
	runningCount := 0
	for bag := range bagsInfo {
		colours := bagColoursInside(bagsInfo, bag)
		if colours["shiny gold"] {
			runningCount++
		}
	}
	fmt.Println("Number of differently coloured bags that might contain a shiny gold bag: " +
		strconv.Itoa(runningCount))

	inner := totalBagsInside(bagsInfo, "shiny gold")
	fmt.Println("Number of bags there must be inside a shiny gold bag: " +
		strconv.Itoa(inner))
}
