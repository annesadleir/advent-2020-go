package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("C:\\Workarea\\advent2020\\inputs\\day01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var inverses [2021]bool
	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)
		if inverses[num] == true {
			result := num * (2020-num)
			fmt.Println(result)
			return
		} else {
			inverses[2020-num] = true
		}
	}

}
