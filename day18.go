package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
)

const (
	OPEN  = "("
	CLOSE = ")"
	TIMES = "*"
	PLUS  = "+"
)

func main() {
	day18Sums := readDay18()
	fmt.Println(partOneDay18(day18Sums))
	// 6640667297513
	fmt.Println(partTwoDay18(day18Sums))
	// 451589894841552
}

func partTwoDay18(sums []string) string {
	sumParts := splitLines(sums)
	sumParts = reduceLines(sumParts, reduce18pt2)
	result := sumBigly(sumParts)
	return result.String()
}

func partOneDay18(sums []string) string {
	sumParts := splitLines(sums)
	sumParts = reduceLines(sumParts, reduce18pt1)
	result := sumBigly(sumParts)
	return result.String()
}

func sumBigly(parts [][]string) *big.Int {
	total := big.NewInt(0)
	for _, line := range parts {
		if len(line) > 1 {
			panic("Summing on incompletely reduced lines")
		}
		asInt64, err := strconv.ParseInt(line[0], 10, 64)
		if err != nil {
			panic(err)
		}
		total = total.Add(total, big.NewInt(asInt64))
	}
	return total
}

func reduceLines(parts [][]string, reducer func([]string) []string) [][]string {
	reducedLines := make([][]string, 0)
	for _, line := range parts {
		for len(line) > 1 {
			line = reducer(line)
		}
		reducedLines = append(reducedLines, line)
	}
	return reducedLines
}

func reduce18pt2(line []string) []string {
	if len(line)%2 != 1 {
		panic("Should be odd number of symbols")
	}
	_, contains := symbolsContains(line, OPEN)
	if contains {
		start, end := findIndexesOfInnermostBracketPair(line)
		bracketContentEvaluated := reduce18pt2(line[start+1 : end])
		concat := concatenateThreeArrays(line[:start], bracketContentEvaluated, line[end+1:])
		return concat
	} else {
		index, contains := symbolsContains(line, PLUS)
		if contains {
			adduand1, _ := strconv.ParseInt(line[index-1], 10, 64)
			adduand2, _ := strconv.ParseInt(line[index+1], 10, 64)
			additionEvaluated := adduand1 + adduand2
			additionStringArray := []string{strconv.FormatInt(additionEvaluated, 10)}
			var beforeAddition []string
			if index == 1 {
				beforeAddition = []string{}
			} else {
				beforeAddition = line[:index-1]
			}
			var afterAddition []string
			if index == len(line)-2 {
				afterAddition = []string{}
			} else {
				afterAddition = line[index+2:]
			}
			concat := concatenateThreeArrays(beforeAddition, additionStringArray, afterAddition)
			return reduce18pt2(concat)
		} else {
			firstNumber, _ := strconv.ParseInt(line[0], 10, 64)
			for counter := 2; counter < len(line); counter += 2 {
				secondNumber, _ := strconv.ParseInt(line[counter], 10, 64)
				if line[counter-1] == PLUS {
					firstNumber = firstNumber + secondNumber
				} else if line[counter-1] == TIMES {
					firstNumber = firstNumber * secondNumber
				} else {
					panic("Second symbol should have been * or +")
				}
				if firstNumber < 0 {
					panic("Looks like an int overflow")
				}
			}
			asStr := strconv.FormatInt(firstNumber, 10)
			return []string{asStr}
		}
	}
}

func reduce18pt1(line []string) []string {
	_, contains := symbolsContains(line, OPEN)
	if contains {
		start, end := findIndexesOfInnermostBracketPair(line)
		bracketContentEvaluated := reduce18pt1(line[start+1 : end])
		concat := concatenateThreeArrays(line[:start], bracketContentEvaluated, line[end+1:])
		return concat
	} else {
		if len(line)%2 != 1 {
			panic("Should be odd number of symbols")
		}
		firstNumber, _ := strconv.ParseInt(line[0], 10, 64)
		for counter := 2; counter < len(line); counter += 2 {
			secondNumber, _ := strconv.ParseInt(line[counter], 10, 64)
			if line[counter-1] == PLUS {
				firstNumber = firstNumber + secondNumber
			} else if line[counter-1] == TIMES {
				firstNumber = firstNumber * secondNumber
			} else {
				panic("Second symbol should have been * or +")
			}
			if firstNumber < 0 {
				panic("Looks like an int overflow")
			}
		}
		asStr := strconv.FormatInt(firstNumber, 10)
		return []string{asStr}
	}
}

func findIndexesOfInnermostBracketPair(line []string) (int, int) {
	deepestOpenIndex := indexOfDeepestOpenBracket(line)
	deepestCloseIndex := indexOfItsCloseBracket(line, deepestOpenIndex)
	return deepestOpenIndex, deepestCloseIndex
}

func indexOfDeepestOpenBracket(line []string) int {
	openBrackets := 0
	maxOpenBrackets := 0
	deepestOpenIndex := -1
	for i, symbol := range line {
		if symbol == OPEN {
			openBrackets++
			if openBrackets > maxOpenBrackets {
				deepestOpenIndex = i
				maxOpenBrackets = openBrackets
			}
		} else if symbol == CLOSE {
			openBrackets--
		}
	}
	if openBrackets > 0 {
		panic("Bracket matching error")
	}
	return deepestOpenIndex
}

func indexOfItsCloseBracket(line []string, openIndex int) int {
	openBrackets := 1
	for index := openIndex + 1; index < len(line); index++ {
		if line[index] == OPEN {
			openBrackets++
		} else if line[index] == CLOSE {
			openBrackets--
		}
		if openBrackets == 0 {
			return index
		}
	}
	return -1
}

func concatenateThreeArrays(a []string, b []string, c []string) []string {
	result := make([]string, 0)
	for _, ai := range a {
		result = append(result, ai)
	}
	for _, bi := range b {
		result = append(result, bi)
	}
	for _, ci := range c {
		result = append(result, ci)
	}
	return result
}

func symbolsContains(symbols []string, str string) (int, bool) {
	for i, s := range symbols {
		if s == str {
			return i, true
		}
	}
	return -1, false
}

func splitLines(sums []string) [][]string {
	output := make([][]string, 0)
	for _, line := range sums {
		line = strings.ReplaceAll(line, " ", "")
		parts := strings.Split(line, "")
		output = append(output, parts)
	}
	return output
}

func readDay18() []string {
	file, err := os.Open("C:\\Workarea\\advent-2020-go\\inputs\\day18.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	expressions := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		expressions = append(expressions, line)
	}
	return expressions
}
