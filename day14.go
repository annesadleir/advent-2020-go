package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const MASK = "mask"
const MEM = "mem"
const EQUALS = " = "
const X = "X"
const ONE = "1"
const ZERO = "0"

func main() {
	partOneAnswer := partOneDay14()
	fmt.Println(partOneAnswer)
	partTwoAnswer := partTwoDay14()
	fmt.Println(partTwoAnswer)
}

func partTwoDay14() int64 {
	inputs := readLinesDay14()

	mask := ""
	memory := make(map[int64]int64)

	for _, input := range inputs {
		if strings.Contains(input, MASK) {
			mask = extractMask(input)
		} else if strings.Contains(input, MEM) {
			memory = updateMemoryPartTwo(memory, mask, input)
		}
	}
	return sumMemory(memory)
}

func updateMemoryPartTwo(memory map[int64]int64, mask string, input string) map[int64]int64 {
	address, value := readAddressAndValue(input)
	maskedAddresses := allMaskedAddresses(mask, address)
	for _, address := range maskedAddresses {
		memory[address] = value
	}
	return memory
}

func allMaskedAddresses(mask string, address int64) []int64 {
	booleanVal := boolean36Char(address)

	maskedAddressStrings := make([]string, 0)
	maskedAddressStrings = append(maskedAddressStrings, booleanVal)

	for index := len(mask) - 1; index >= 0; index-- {
		maskChar := mask[index : index+1]
		if maskChar == ONE {
			maskedAddressStrings = replaceAllAtIndex(maskedAddressStrings, ONE, index)
		} else if maskChar == X {
			expanded := make([]string, 0)
			for _, maskedAddress := range maskedAddressStrings {
				expanded = append(expanded, replaceAtIndex(maskedAddress, ONE, index))
				expanded = append(expanded, replaceAtIndex(maskedAddress, ZERO, index))
			}
			if len(expanded) != len(maskedAddressStrings) *2 {
				fmt.Println("Error here")
			}
			maskedAddressStrings = expanded
		}
	}

	maskedValues := make([]int64, 0)
	for _, valueStr := range maskedAddressStrings {
		result, err := strconv.ParseInt(valueStr, 2, 64)
		if err != nil {
			panic(err)
		}
		maskedValues = append(maskedValues, result)
	}

	return maskedValues
}

func replaceAllAtIndex(originals []string, char string, index int) []string {
	result := make([]string, 0)
	for _, original := range originals {
		result = append(result, replaceAtIndex(original, char, index))
	}
	return result
}

func replaceAtIndex(original string, char string, index int) string {
	out := []rune(original)
	chars := []rune(char)
	out[index] = chars[0]
	return string(out)
}

func boolean36Char(address int64) string {
	result := strconv.FormatInt(address, 2)
	for len(result) < 36 {
		result = "0" + result
	}
	return result
}

func partOneDay14() int64 {
	inputs := readLinesDay14()

	mask := ""
	memory := make(map[int64]int64)

	for _, input := range inputs {
		if strings.Contains(input, MASK) {
			mask = extractMask(input)
		} else if strings.Contains(input, MEM) {
			memory = updateMemoryPartOne(memory, mask, input)
		}
	}
	return sumMemory(memory)
}

func updateMemoryPartOne(memory map[int64]int64, mask string, input string) map[int64]int64 {
	address, value := readAddressAndValue(input)
	maskedValue := applyMask(mask, value)
	memory[address] = maskedValue
	return memory
}

func applyMask(mask string, value int64) int64 {
	booleanVal := strconv.FormatInt(value, 2)
	maskedValue := mask

	for index := len(mask) - 1; index >= 0; index-- {
		if mask[index:index+1] == X {
			reqString := ""
			valIndex := len(booleanVal) - (len(mask) - index)
			if valIndex < 0 {
				reqString = "0"
			} else {
				reqString = string(booleanVal[valIndex])
			}
			maskedValue = replaceAtIndex(maskedValue, reqString, index)
		}
	}
	result, _ := strconv.ParseInt(maskedValue, 2, 64)
	return result
}

func readAddressAndValue(input string) (int64, int64) {
	parts := strings.Split(input, EQUALS)
	addressStr := strings.Replace(parts[0], MEM, "", -1)
	addressStr = strings.Replace(addressStr, "[", "", -1)
	addressStr = strings.Replace(addressStr, "]", "", -1)
	address, err := strconv.ParseInt(addressStr, 10, 64)
	if err != nil {
		panic(err)
	}

	value, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		panic(err)
	}
	return address, value
}

func extractMask(input string) string {
	parts := strings.Split(input, EQUALS)
	return parts[1]
}

func sumMemory(memory map[int64]int64) int64 {
	var runningTotal int64
	for _, value := range memory {
		runningTotal += value
	}
	return runningTotal
}

func readLinesDay14() []string {
	file, err := os.Open("C:\\Workarea\\advent-2020-go\\inputs\\day14.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	instructions := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		instructions = append(instructions, line)
	}

	return instructions
}
