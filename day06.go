package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type CustomsGroup struct {
	positives map[string]bool
}

func createCustomsGroup() CustomsGroup {
	return CustomsGroup{make(map[string]bool, 0)}
}

func (group CustomsGroup) groupPositiveCount() int {
	return len(group.positives)
}

func readCustomsGroups() []CustomsGroup {
	file, err := os.Open("C:\\Workarea\\advent-2020-go\\inputs\\day06.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	groups := make([]CustomsGroup, 0)
	current := createCustomsGroup()
	groups = append(groups, current)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			current = createCustomsGroup()
			groups = append(groups, current)
		} else {
			for _, ch := range line {
				current.positives[string(ch)] = true
			}
		}
	}
	return groups
}

func readCustomsGroupsForPart2() []CustomsGroup {
	file, err := os.Open("C:\\Workarea\\advent-2020-go\\inputs\\day06.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	groups := make([]CustomsGroup, 0)
	current := createCustomsGroup()
	groups = append(groups, current)
	firstLine := true

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			current = createCustomsGroup()
			groups = append(groups, current)
			firstLine = true
		} else if !firstLine {
			for currentPositive, _ := range current.positives {
				if !strings.Contains(line, currentPositive) {
					delete(current.positives, currentPositive)
				}
			}
		} else {
			for _, ch := range line {
				current.positives[string(ch)] = true
			}
			firstLine = false
		}
	}
	return groups
}

func main() {

	groups := readCustomsGroups()

	runningTotal := 0
	for _, group := range groups {
		runningTotal += group.groupPositiveCount()
	}

	fmt.Println("Part one answer: " + strconv.Itoa(runningTotal))

	groups = readCustomsGroupsForPart2()

	runningTotal = 0
	for _, group := range groups {
		runningTotal += group.groupPositiveCount()
	}

	fmt.Println("Part two answer: " + strconv.Itoa(runningTotal))
}
