package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func readAdapters() []int {
	file, err := os.Open("C:\\Workarea\\advent-2020-go\\inputs\\day10.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	numbers := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, num)
	}
	return numbers
}

func main() {
	adapters := readAdapters()

	partOne := partOneDayTen(adapters)
	fmt.Print("Part one answer is: ")
	fmt.Println(partOne)

	partTwo := partTwoDayTen(adapters)
	fmt.Println(partTwo)
}

func partOneDayTen(adapters []int) int {
	sort.Ints(adapters)
	countOneDiff := 0
	countThreeDiff := 1  // the diff to the device at the end is always 3
	previousJoltage := 0 // start at the outlet, at 0
	for _, adapter := range adapters {
		diff := adapter - previousJoltage
		if diff == 1 {
			countOneDiff++
		} else if diff == 3 {
			countThreeDiff++
		}
		previousJoltage = adapter
	}
	return countThreeDiff * countOneDiff
}

func partTwoDayTen(adapters []int) int {
	runsOfOne := runsOfOne(adapters)
	combinationsForRun := map[int]int{
		1: 1, 2: 2, 3: 4, 4: 7,
	}
	combinations := 1
	for _, run := range runsOfOne {
		combinations *= combinationsForRun[run]
	}
	return combinations
}

func runsOfOne(adapters []int) []int {
	sort.Ints(adapters)
	previousJoltage := 0 // start at the outlet, at 0
	currentRunOfOne := 0
	runsOfOne := make([]int, 0)

	for _, adapter := range adapters {
		diff := adapter - previousJoltage
		if diff == 1 {
			currentRunOfOne++
		} else if diff == 3 {
			if currentRunOfOne > 0 {
				runsOfOne = append(runsOfOne, currentRunOfOne)
			}
			currentRunOfOne = 0
		}
		previousJoltage = adapter
	}
	if currentRunOfOne > 0 {
		runsOfOne = append(runsOfOne, currentRunOfOne)
	}

	return runsOfOne
}
