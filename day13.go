package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type BusSituationPartOne struct {
	earliest int
	buses    []Bus
}

type Bus = int

func main() {
	fmt.Print("Part one answer is: ")
	fmt.Println(partOneDay13())
}

func partOneDay13() int {
	partOneSituation := readPartOneDay13()

	arrival := partOneSituation.earliest
	leastPossibleTimeToWait := 0
	earliestPossibleBusNumber := 0

	for _, bus := range partOneSituation.buses {
		timeToWait := bus - (arrival % bus)
		if earliestPossibleBusNumber == 0 || timeToWait < leastPossibleTimeToWait {
			leastPossibleTimeToWait = timeToWait
			earliestPossibleBusNumber = bus
		}
	}

	return leastPossibleTimeToWait * earliestPossibleBusNumber
}

func readPartOneDay13() BusSituationPartOne {
	file, err := os.Open("C:\\Workarea\\advent-2020-go\\inputs\\day13.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	earliest, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	buses := make([]Bus, 0)
	parts := strings.Split(scanner.Text(), ",")
	for _, id := range parts {
		if id != "x" {
			num, _ := strconv.Atoi(id)
			buses = append(buses, num)
		}
	}

	return BusSituationPartOne{earliest, buses}
}
