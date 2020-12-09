package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	bigInts := readBigInts()
	part1 := lookForFirstNonSum(bigInts)
	fmt.Print("Part one answer is: ")
	fmt.Println(part1)
	part2 := findContiguousNumsSum(bigInts, part1)
	fmt.Print("Part two answer is: ")
	fmt.Println(part2)
}

func readBigInts() []int64 {
	file, err := os.Open("C:\\Workarea\\advent-2020-go\\inputs\\day09.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	numbers := make([]int64, 0)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, num)
	}
	return numbers
}

func lookForFirstNonSum(bigInts []int64) int64 {
	counter := 0
	for _, bigInt := range bigInts {
		if counter >= 25 &&
			!sliceContainsSum(bigInts[counter-25:counter], bigInt) {
			return bigInt
		}
		counter++
	}
	return -1
}

func sliceContainsSum(previous []int64, sum int64) bool {
	mapped := make(map[int64]bool, 0)
	for _, n := range previous {
		req := sum - n
		if mapped[req] {
			return true
		}
		mapped[n] = true
	}
	return false
}

func findContiguousNumsSum(all []int64, required int64) int64 {
	startIndex := 0
	endIndex := 1

	for startIndex < len(all) {
		slice := all[startIndex:endIndex]
		sum := sliceSum(slice)
		if sum == required {
			return smallestPlusLargestInSlice(slice)
		} else if sum < required {
			endIndex++
		} else {
			startIndex++
			endIndex = startIndex + 1
		}
	}
	return -1
}
func smallestPlusLargestInSlice(slice []int64) int64 {
	min := slice[0]
	max := slice[0]
	for _, num := range slice {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	return min + max
}

func sliceSum(slice []int64) int64 {
	var total int64 = 0
	for _, n := range slice {
		total += n
	}
	return total
}
