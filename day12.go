package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Ship struct {
	Position  Position
	Direction CompassPoint
	Waypoint  Position
}

type Position struct {
	Northing int
	Easting  int
}

func (position Position) manhattanDistance() int {
	return abs(position.Northing) + abs(position.Easting)
}

func (startingPosition Position) move(direction Position, times int) Position {
	newPosition := Position{startingPosition.Northing, startingPosition.Easting}
	newPosition.Northing += times * direction.Northing
	newPosition.Easting += times * direction.Easting
	return newPosition
}

type CompassPoint = int

const (
	N CompassPoint = 0
	E CompassPoint = 1
	S CompassPoint = 2
	W CompassPoint = 3
)

var compassPoints = map[string]CompassPoint{"N": N, "E": E, "S": S, "W": W}
var movements = map[CompassPoint]Position{
	N: {1, 0},
	E: {0, 1},
	S: {-1, 0},
	W: {0, -1}}

type Instruction struct {
	Action string
	Value  int
}

func readNavigationInstructions() []Instruction {

	file, err := os.Open("C:\\Workarea\\advent-2020-go\\inputs\\day12.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	instructions := make([]Instruction, 0)
	for scanner.Scan() {
		line := scanner.Text()
		action := line[:1]
		value, _ := strconv.Atoi(line[1:])
		instructions = append(instructions, Instruction{action, value})

		if (action == "L" || action == "R") && value%90 != 0 {
			fmt.Println("Not a right-angle turn!")
		}
	}

	return instructions
}

func main() {
	instructions := readNavigationInstructions()
	ship := Ship{Position{0, 0}, E, Position{0, 0}}
	endPosition := followInstructionsPartOne(ship, instructions)
	partOneResult := endPosition.manhattanDistance()
	fmt.Println(partOneResult)

	instructions = readNavigationInstructions()
	ship = Ship{Position{0, 0}, E, Position{1, 10}}
	partTwoEndPosition := followInstructionsPartTwo(ship, instructions)
	partTwoResult := partTwoEndPosition.manhattanDistance()
	fmt.Println(partTwoResult)
}

func followInstructionsPartTwo(ship Ship, instructions []Instruction) Position {
	for _, instruction := range instructions {
		if instruction.Action == "L" || instruction.Action == "R" {
			ship = turnWaypoint(ship, instruction)
		} else if instruction.Action == "F" {
			ship = moveShip(ship, instruction)
		} else {
			ship = moveWaypoint(ship, instruction)
		}
	}
	return ship.Position
}

func turnWaypoint(ship Ship, instruction Instruction) Ship {

	numTurns := instruction.Value / 90
	if instruction.Action == "L" {
		numTurns = 4 - numTurns
	}

	var northing int
	var easting int

	switch numTurns {
	case 1:
		northing = -ship.Waypoint.Easting
		easting = ship.Waypoint.Northing
	case 2:
		northing = -ship.Waypoint.Northing
		easting = -ship.Waypoint.Easting
	case 3:
		northing = ship.Waypoint.Easting
		easting = -ship.Waypoint.Northing
	case 0:
		northing = ship.Waypoint.Northing
		easting = ship.Waypoint.Easting
	}

	ship.Waypoint = Position{northing, easting}
	return ship
}

func moveShip(ship Ship, instruction Instruction) Ship {
	ship.Position = ship.Position.move(ship.Waypoint, instruction.Value)
	return ship
}

func moveWaypoint(ship Ship, instruction Instruction) Ship {
	direction := movements[compassPoints[instruction.Action]]
	times := instruction.Value
	ship.Waypoint = ship.Waypoint.move(direction, times)
	return ship
}

func followInstructionsPartOne(ship Ship, instructions []Instruction) Position {
	for _, instruction := range instructions {
		if instruction.Action == "L" || instruction.Action == "R" {
			ship = turn(ship, instruction)
		} else {
			ship = move(ship, instruction)
		}
	}
	return ship.Position
}

func turn(ship Ship, instruction Instruction) Ship {
	numTurns := instruction.Value / 90
	if instruction.Action == "L" {
		numTurns = 4 - numTurns
	}
	ship.Direction = (ship.Direction + numTurns) % 4
	return ship
}

func move(ship Ship, instruction Instruction) Ship {
	var direction Position
	if instruction.Action == "F" {
		direction = movements[ship.Direction]
	} else {
		direction = movements[compassPoints[instruction.Action]]
	}
	times := instruction.Value
	ship.Position = ship.Position.move(direction, times)
	return ship
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
