package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Ferry struct {
	SeatingStates []SeatingState
	MaxRow        int
	MaxCol        int
}

type SeatingState = map[Location]Status

type Location struct {
	Row    int
	Column int
}

type Status = string

const (
	Occupied Status = "#"
	Empty    Status = "L"
	Floor    Status = "."
)

type Direction = Location

var (
	NorthWest  Direction = Location{-1, -1}
	North      Direction = Location{-1, 0}
	NorthEast  Direction = Location{-1, 1}
	West       Direction = Location{0, -1}
	East       Direction = Location{0, 1}
	SouthWest  Direction = Location{1, -1}
	South      Direction = Location{1, 0}
	SouthEast  Direction = Location{1, 1}
	directions           = []Direction{NorthWest, North, NorthEast, West, East, SouthWest, South, SouthEast}
)

func main() {
	//ferry := readFerryPlan()
	//occupiedOnceStable := partOneDay11(ferry)
	//fmt.Println(occupiedOnceStable)
	ferry := readFerryPlan()
	occupiedOnceStable := partTwoDay11(ferry)
	fmt.Println(occupiedOnceStable)
}

func partOneDay11(ferry Ferry) int {
	for unstableSeatingStates(ferry) {
		ferry.SeatingStates = append(ferry.SeatingStates, nextSeatingState(ferry, partOneNextStatus))
	}
	return countOccupiedTotal(ferry)
}

func partTwoDay11(ferry Ferry) int {
	loops := 0
	for unstableSeatingStates(ferry) {
		loops++
		if loops % 10 == 0 {
			fmt.Println(loops)
		}
		ferry.SeatingStates = append(ferry.SeatingStates, nextSeatingState(ferry, partTwoNextStatus))
	}
	return countOccupiedTotal(ferry)
}

func nextSeatingState(ferry Ferry, nextStatusOfLocation func(ferry Ferry, state SeatingState, location Location) Status) SeatingState {
	current := ferry.currentSeatingState()
	next := make(map[Location]Status, 0)

	for location, _ := range current {
		next[location] = nextStatusOfLocation(ferry, current, location)
	}
	return next
}

func partOneNextStatus(ferry Ferry, state SeatingState, location Location) Status {
	nextStatus := state[location]
	if nextStatus == Empty && countOccupiedSurrounding(ferry, state, location) == 0 {
		nextStatus = Occupied
	} else if nextStatus == Occupied && countOccupiedSurrounding(ferry, state, location) >= 4 {
		nextStatus = Empty
	}
	return nextStatus
}

func partTwoNextStatus(ferry Ferry, state SeatingState, location Location) Status {
	nextStatus := state[location]
	if nextStatus == Empty && countVisibleSurrounding(ferry, state, location) == 0 {
		nextStatus = Occupied
	} else if nextStatus == Occupied && countVisibleSurrounding(ferry, state, location) >= 5 {
		nextStatus = Empty
	}
	return nextStatus
}

func unstableSeatingStates(ferry Ferry) bool {
	seatingStates := ferry.SeatingStates
	if len(seatingStates) == 1 {
		return true
	}
	return !equalSeatingStates(seatingStates[len(seatingStates)-2:])
}

func equalSeatingStates(seatingStates []SeatingState) bool {
	first := seatingStates[0]
	second := seatingStates[1]
	for location, status := range first {
		if second[location] != status {
			return false
		}
	}
	return true
}

func readFerryPlan() Ferry {
	file, err := os.Open("C:\\Workarea\\advent-2020-go\\inputs\\day11.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	seatingStates := make([]SeatingState, 0)
	initialSeatingState := make(map[Location]Status, 0)
	seatingStates = append(seatingStates, initialSeatingState)

	row := 0
	maxCol := 0
	for scanner.Scan() {
		line := scanner.Text()
		if maxCol == 0 {
			maxCol = len(line) - 1
		}
		for col := 0; col <= maxCol; col++ {
			location := Location{row, col}
			initialSeatingState[location] = line[col : col+1]
		}
		row++
	}
	return Ferry{seatingStates, row - 1, maxCol}
}

func (ferry Ferry) currentSeatingState() SeatingState {
	seatingStates := ferry.SeatingStates
	return seatingStates[len(seatingStates)-1]
}

func countOccupiedTotal(ferry Ferry) int {
	count := 0
	for _, status := range ferry.currentSeatingState() {
		if status == Occupied {
			count++
		}
	}
	return count
}

func countVisibleSurrounding(ferry Ferry, seatingState SeatingState, seat Location) int {
	count := 0
	for _, direction := range directions {
		if occupiedSeatVisibleInDirection(ferry, seatingState, seat, direction) {
			count++
		}
	}
	return count
}

func occupiedSeatVisibleInDirection(ferry Ferry, state SeatingState, seat Location, direction Direction) bool {
	for neighbour := seat.neighbour(direction); ferry.contains(neighbour); neighbour = neighbour.neighbour(direction) {
		if state[neighbour] != Floor {
			return state[neighbour] == Occupied
		}
	}
	return false
}

func countOccupiedSurrounding(ferry Ferry, seatingState SeatingState, seat Location) int {
	count := 0
	for _, direction := range directions {
		neighbour := seat.neighbour(direction)
		if ferry.contains(neighbour) && seatingState[neighbour] == Occupied {
			count++
		}
	}
	return count
}

func (ferry Ferry) contains(location Location) bool {
	return location.Row >= 0 && location.Row <= ferry.MaxRow &&
		location.Column >= 0 && location.Column <= ferry.MaxCol
}

func (location Location) neighbour(direction Direction) Location {
	return Location{location.Row + direction.Row, location.Column + direction.Column}
}
