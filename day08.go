package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type InstructionD8 struct {
	op  string
	num int
}

func readInstructions() []InstructionD8 {
	file, err := os.Open("C:\\Workarea\\advent-2020-go\\inputs\\day08.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	instructions := make([]InstructionD8, 0)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		num, _ := strconv.Atoi(parts[1])
		instructions = append(instructions, InstructionD8{parts[0], num})
	}

	return instructions
}

func main() {
	instructions := readInstructions()
	fmt.Print("Part one answer: ")
	part1(instructions)
	fmt.Print("Part two answer: ")
	part2(instructions)
}

func part1(instructions []InstructionD8) {
	accumulator := 0
	counter := 0
	visited := make([]bool, len(instructions)+1)

	for !visited[counter] {
		visited[counter] = true
		instruction := instructions[counter]

		switch instruction.op {
		case "acc":
			accumulator += instruction.num
			counter++
		case "jmp":
			counter += instruction.num
		case "nop":
			counter++
		}
	}
	fmt.Println(accumulator)
}

func part2(instructions []InstructionD8) {
	changeIndex := 0
	foundAnswer := false

	for !foundAnswer {
		if instructions[changeIndex].op != "acc" {
			newInstructions := copyWithChangeAt(instructions, changeIndex)
			if test(newInstructions) {
				foundAnswer = true
			}
		}
		changeIndex++
	}
}

func copyWithChangeAt(instructions []InstructionD8, index int) []InstructionD8 {
	copy := make([]InstructionD8, len(instructions))
	for i, instr := range instructions {
		copy[i] = InstructionD8{instr.op, instr.num}
	}
	instrToChangeOp := instructions[index].op
	if instrToChangeOp == "jmp" {
		copy[index].op = "nop"
	} else if instrToChangeOp == "nop" {
		copy[index].op = "jmp"
	}
	return copy
}

func test(instructions []InstructionD8) bool {
	accumulator := 0
	counter := 0
	visited := make([]bool, len(instructions)+1)

	for counter < len(instructions) {
		if visited[counter] {
			return false
		}
		visited[counter] = true
		instruction := instructions[counter]

		switch instruction.op {
		case "acc":
			accumulator += instruction.num
			counter++
		case "jmp":
			counter += instruction.num
		case "nop":
			counter++
		}
	}
	fmt.Println(accumulator)
	return true
}
