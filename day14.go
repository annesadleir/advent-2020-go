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

func main() {
	partOneAnswer := partOneDay14()
	fmt.Println(partOneAnswer)
}

func partOneDay14() int64 {
	inputs := readLinesDay14()

	mask := ""
	memory := make(map[int]int64)

	for _, input := range inputs {
		if strings.Contains(input, MASK) {
			mask = extractMask(input)
		} else if strings.Contains(input, MEM) {
			memory = updateMemory(memory, mask, input)
		}
	}
	return sumMemory(memory)
}

func updateMemory(memory map[int]int64, mask string, input string) map[int]int64 {
	address, value := readAddressAndValue(input)
	maskedValue := applyMask(mask, value)
	memory[address] = maskedValue
	return memory
}

func applyMask(mask string, value int64) int64 {
	booleanVal := strconv.FormatInt(value, 2)
	maskedValue := mask

	for index := len(mask) - 1; index >= 0; index -- {
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

func replaceAtIndex(original string, char string, index int) string {
	out := []rune(original)
	chars := []rune(char)
	out[index] = chars[0]
	return string(out)
}

func readAddressAndValue(input string) (int, int64) {
	parts := strings.Split(input, EQUALS)
	addressStr := strings.Replace(parts[0], MEM, "", -1)
	addressStr = strings.Replace(addressStr, "[", "", -1)
	addressStr = strings.Replace(addressStr, "]", "", -1)
	address, err := strconv.Atoi(addressStr)
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

func sumMemory(memory map[int]int64) int64 {
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
