package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type BoardingPass struct {
	row    int
	column int
}

func (pass BoardingPass) seatNumber() int {
	return pass.row*8 + pass.column
}

func createBoardingPass(line string) BoardingPass {
	rowStr := line[0:7]
	rowBin := strings.Replace(rowStr, "F", "0", -1)
	rowBin = strings.Replace(rowBin, "B", "1", -1)
	row, _ := strconv.ParseInt(rowBin, 2, 64)

	colStr := line[7:]
	colBin := strings.Replace(colStr, "L", "0", -1)
	colBin = strings.Replace(colBin, "R", "1", -1)
	column, _ := strconv.ParseInt(colBin, 2, 64)
	pass := BoardingPass{int(row), int(column)}
	return pass
}

func readBoardingPasses() []BoardingPass {
	file, err := os.Open("C:\\Workarea\\advent-2020-go\\inputs\\day05.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	passes := make([]BoardingPass, 0)
	for scanner.Scan() {
		line := scanner.Text()
		pass := createBoardingPass(line)
		passes = append(passes, pass)
	}
	return passes
}

func main() {
	passes := readBoardingPasses()

	maxSeat := 0
	for _, pass := range passes {
		if pass.seatNumber() > maxSeat {
			maxSeat = pass.seatNumber()
		}
	}
	fmt.Println("Max seat number is " + strconv.Itoa(maxSeat))

	assignedSeats := make([]bool, maxSeat+1)
	for _, pass := range passes {
		assignedSeats[pass.seatNumber()] = true
	}

	for i, assigned := range assignedSeats {
		if !assigned {
			if i > 1 && assignedSeats[i-1] && assignedSeats[i+1] {
				fmt.Println("Unassigned seat with assigned seats either side is " +
					strconv.Itoa(i))
			}
		}
	}

}
