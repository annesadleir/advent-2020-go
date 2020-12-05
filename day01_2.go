package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Pair struct {
	A int
	B int
}

func (pair Pair) Total() int {
	return pair.A + pair.B
}

func (pair Pair) Product() int {
	return pair.A * pair.B
}

func main() {
	file, err := os.Open("C:\\Workarea\\advent2020\\inputs\\day01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	pairs := make([]Pair, 0)
	numbers := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)

		for _, pair := range pairs {
			if pair.Total()+num == 2020 {
				fmt.Println(pair.Product() * num)
				return
			}
		}
		for _, number := range numbers {
			pairs = append(pairs, Pair{number, num})
		}
		numbers = append(numbers, num)
	}

}
