package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type PasswordInfo struct {
	lowNum    int
	highNum   int
	character string
	password  string
}

func (passwordInfo PasswordInfo) ValidPart1() bool {
	count := strings.Count(passwordInfo.password, passwordInfo.character)
	return count <= passwordInfo.highNum && count >= passwordInfo.lowNum
}

func (passwordInfo PasswordInfo) ValidPart2() bool {
	firstMatches := passwordInfo.character == passwordInfo.password[passwordInfo.lowNum-1:passwordInfo.lowNum]
	secondMatches := passwordInfo.character == passwordInfo.password[passwordInfo.highNum-1:passwordInfo.highNum]
	return (firstMatches && !secondMatches) || (secondMatches && !firstMatches)
}

func createParsewordInfo(input string) PasswordInfo {
	halves := strings.Split(input, ": ")
	password := halves[1]

	rules := strings.Split(halves[0], " ")
	character := rules[1]

	span := strings.Split(rules[0], "-")

	min, _ := strconv.Atoi(span[0])
	max, _ := strconv.Atoi(span[1])

	return PasswordInfo{
		lowNum:    min,
		highNum:   max,
		character: character,
		password:  password,
	}
}

func main() {

	file, err := os.Open("C:\\Workarea\\advent-2020-go\\inputs\\day02.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	runningTotalPart1 := 0
	runningTotalPart2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		passwordInfo := createParsewordInfo(line)
		if passwordInfo.ValidPart1() {
			runningTotalPart1++
		}
		if passwordInfo.ValidPart2() {
			runningTotalPart2++
		}
	}
	fmt.Println(runningTotalPart1)
	fmt.Println(runningTotalPart2)
}
