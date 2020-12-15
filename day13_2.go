package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type BusRequirement struct {
	Period int64
	Offset int64
}

func (busReq BusRequirement) String() string {
	return fmt.Sprintf("Period: %d, offset: %d", busReq.Period, busReq.Offset)
}

func reduce(busRequirements []BusRequirement) BusRequirement {
	//sort.Slice(busRequirements, func(i, j int) bool {
	//	return busRequirements[i].Period < busRequirements[j].Period
	//})

	for len(busRequirements) > 1 {
		newBusRequirement := combineBusRequirements(busRequirements[:2])
		smallerBusRequirements := make([]BusRequirement, 0)
		smallerBusRequirements = append(smallerBusRequirements, newBusRequirement)
		for _, busReq := range busRequirements[2:] {
			smallerBusRequirements = append(smallerBusRequirements, busReq)
		}
		busRequirements = smallerBusRequirements
	}

	return busRequirements[0]
}

func partTwo(originalBusRequirements []BusRequirement) {

	busRequirements := originalBusRequirements

	lastReq := reduce(busRequirements)

	zeroPoint := lastReq.Offset

	for zeroPoint < 0 {
		zeroPoint += lastReq.Period
	}

	for zeroPoint > lastReq.Period {
		zeroPoint -= lastReq.Period
	}

	for _, req := range originalBusRequirements {
		checkFitsReq("After reduction", zeroPoint, req)
	}
	fmt.Println(zeroPoint)
}

func main() {
	partTwo(simpleBusSituation())
	partTwo(nextBusSituation())
	partTwo(thirdBusSituation())
	partTwo(fourthBusSituation())
	partTwo(fifthBusSituation())
	partTwo(demoBusSituation())
	partTwo(readBusSituationPartTwo())
	
	// correct answer = 404517869995362
	// but I had to give up and calculate it in Java

	// 1525636236640101
	// 1525636236640101
	// 1525636236640101
	// 506410616693994

	// fmt.Println(result.Offset)
	// 313793439223925 wrong
	// 1882780318552089
	// 6722635721222522470
	// 126337847429708
	// 1502707549287803
}

func combineBusRequirements(requirements []BusRequirement) BusRequirement {
	var first, second BusRequirement
	if requirements[0].Period > requirements[1].Period {
		first = requirements[0]
		second = requirements[1]
	} else {
		first = requirements[1]
		second = requirements[0]
	}

	// find Bezout numbers
	firstBez, secondBez := Bezout(first.Period, second.Period)
	if firstBez*first.Period+secondBez*second.Period != 1 {
		panic("Bezout numbers not working???")
	}

	startPoint := firstBez*first.Period*second.Offset +
		secondBez*second.Period*first.Offset

	gcm := first.Period * second.Period

	//for startPoint < 0 {
	//	startPoint = -startPoint
	//}

	for startPoint < 0 {
		startPoint += gcm
	}

	for startPoint > gcm {
		startPoint -= gcm
	}

	checkFitsReq("first in reduction", startPoint, first)
	checkFitsReq("second in reduction", startPoint, second)

	return BusRequirement{gcm, startPoint}
}

func checkFitsReq(info string, time int64, requirement BusRequirement) {
	multiplicand := (time - requirement.Offset) / requirement.Period
	if multiplicand*requirement.Period+requirement.Offset != time {
		fmt.Println(info + ": " + requirement.String())
	}
}

func readBusSituationPartTwo() []BusRequirement {
	file, err := os.Open("C:\\Workarea\\advent-2020-go\\inputs\\day13.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	scanner.Text()

	scanner.Scan()
	buses := make([]BusRequirement, 0)
	parts := strings.Split(scanner.Text(), ",")
	for index, id := range parts {
		if id != "x" {
			num, _ := strconv.ParseInt(id, 10, 64)
			//busReq := BusRequirement{Period: num, Offset: int64(index)}
			busReq := createBusRequirement(num, int64(index))
			buses = append(buses, busReq)
		}
	}
	return buses
}

func createBusRequirement(period int64, offset int64) BusRequirement {
	adjOffset := offset
	for adjOffset > period {
		adjOffset -= period
	}
	return BusRequirement{period, -adjOffset}
}

func Bezout(a, b int64) (x, y int64) {
	oldR := a
	r := b
	var oldS, s, oldT, t int64
	oldS = 1
	s = 0
	oldT = 0
	t = 1

	for r != 0 {
		quotient := oldR / r
		oldR, r = quotientSwapThing(oldR, r, quotient)
		oldS, s = quotientSwapThing(oldS, s, quotient)
		oldT, t = quotientSwapThing(oldT, t, quotient)
	}
	return oldS, oldT
}

func quotientSwapThing(oldX, x, quotient int64) (a, b int64) {
	prov := x
	x = oldX - quotient*x
	oldX = prov
	return oldX, x
}

// 17,x,13,19
func simpleBusSituation() []BusRequirement {
	reqs := make([]BusRequirement, 0)
	reqs = append(reqs, BusRequirement{17, 0})
	reqs = append(reqs, BusRequirement{13, -2})
	reqs = append(reqs, BusRequirement{19, -3})
	return reqs
}

func nextBusSituation() []BusRequirement {
	reqs := make([]BusRequirement, 0)
	reqs = append(reqs, BusRequirement{67, 0})
	reqs = append(reqs, BusRequirement{7, -1})
	reqs = append(reqs, BusRequirement{59, -2})
	reqs = append(reqs, BusRequirement{61, -3})
	return reqs
}

func thirdBusSituation() []BusRequirement {
	reqs := make([]BusRequirement, 0)
	reqs = append(reqs, BusRequirement{67, 0})
	reqs = append(reqs, BusRequirement{7, -2})
	reqs = append(reqs, BusRequirement{59, -3})
	reqs = append(reqs, BusRequirement{61, -4})
	return reqs
}

func fourthBusSituation() []BusRequirement {
	reqs := make([]BusRequirement, 0)
	reqs = append(reqs, BusRequirement{67, 0})
	reqs = append(reqs, BusRequirement{7, -1})
	reqs = append(reqs, BusRequirement{59, -3})
	reqs = append(reqs, BusRequirement{61, -4})
	return reqs
}

func fifthBusSituation() []BusRequirement {
	reqs := make([]BusRequirement, 0)
	reqs = append(reqs, BusRequirement{1789, 0})
	reqs = append(reqs, BusRequirement{37, -1})
	reqs = append(reqs, BusRequirement{47, -2})
	reqs = append(reqs, BusRequirement{1889, -3})
	return reqs
}

func demoBusSituation() []BusRequirement {
	reqs := make([]BusRequirement, 0)
	reqs = append(reqs, BusRequirement{7, 0})
	reqs = append(reqs, BusRequirement{13, -1})
	reqs = append(reqs, BusRequirement{59, -4})
	reqs = append(reqs, BusRequirement{31, -6})
	reqs = append(reqs, BusRequirement{19, -7})
	return reqs
}
