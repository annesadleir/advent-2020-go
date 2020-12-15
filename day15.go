package main

import "fmt"

type MapOfLastSpoken = map[int]int

func main() {
	fmt.Println(fmt.Sprintf("Part one answer is: %d", memoryGameDay15(2020)))
	fmt.Println(fmt.Sprintf("Part two answer is: %d", memoryGameDay15(30000000)))
}


func memoryGameDay15(endTurn int) int {
	mapOfLastSpoken, numberSpoken := readDay15Input()

	for lastTurn := len(mapOfLastSpoken) + 1; lastTurn < endTurn; lastTurn++ {
		turnWhenLastSpoken, exists := mapOfLastSpoken[numberSpoken]
		mapOfLastSpoken[numberSpoken] = lastTurn
		if !exists {
			numberSpoken = 0
		} else {
			numberSpoken = lastTurn - turnWhenLastSpoken
		}

	}
	return numberSpoken;
}

func exampleInput() (MapOfLastSpoken, int) {
	lastSpoken := make(map[int]int, 0)
	// 0,3,6
	lastSpoken[0] = 1
	lastSpoken[3] = 2
	return lastSpoken, 6
}

func readDay15Input() (MapOfLastSpoken, int) {
	lastSpoken := make(map[int]int, 0)
	// 7,14,0,17,11,1,2
	lastSpoken[7] = 1
	lastSpoken[14] = 2
	lastSpoken[0] = 3
	lastSpoken[17] = 4
	lastSpoken[11] = 5
	lastSpoken[1] = 6
	return lastSpoken, 2
}
