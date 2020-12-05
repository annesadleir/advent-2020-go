package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Forest struct {
	rowsOfTrees []string
	repeatWidth int
}

func createForest(fileName string) Forest {
	file, err := os.Open("C:\\Workarea\\advent-2020-go\\inputs\\day03.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	rows := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		rows = append(rows, line)
	}
	return Forest{rows, len(rows[0])}
}

func (forest Forest) countTreesAtAngle(right, down int) int {
	row := 0
	column := 0

	treeCount := 0
	for row < len(forest.rowsOfTrees) {
		effectiveTreeIndex := column % forest.repeatWidth

		if forest.rowsOfTrees[row][effectiveTreeIndex:effectiveTreeIndex+1] == "#" {
			treeCount++
		}

		row += down
		column += right
	}

	return treeCount
}

func main() {
	forest := createForest("C:\\Workarea\\advent-2020-go\\inputs\\day03.txt")
	r3d1 := forest.countTreesAtAngle(3, 1)
	fmt.Println(r3d1)

	// part 2
	r1d1 := forest.countTreesAtAngle(1, 1)
	r5d1 := forest.countTreesAtAngle(5, 1)
	r7d1 := forest.countTreesAtAngle(7, 1)
	r1d2 := forest.countTreesAtAngle(1, 2)
	product := r1d1 * r3d1 * r5d1 * r7d1 * r1d2
	fmt.Println(product)
}
