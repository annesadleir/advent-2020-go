package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type BusRequirement struct {
	Period int64
	Offset int64
}

func main() {
	//originalBusRequirements := readBusSituationPartTwo()
	originalBusRequirements := simpleBusSituation()

	busRequirements := originalBusRequirements

	sort.Slice(busRequirements, func(i, j int) bool {
		return busRequirements[i].Period > busRequirements[j].Period
	})

	for len(busRequirements) > 1 {
		newBusRequirement := combineBusRequirements(busRequirements[:2])
		smallerBusRequirements := make([]BusRequirement, 0)
		smallerBusRequirements = append(smallerBusRequirements, newBusRequirement)
		for _, busReq := range busRequirements[2:] {
			smallerBusRequirements = append(smallerBusRequirements, busReq)
		}
		busRequirements = smallerBusRequirements
	}

	lastReq := busRequirements[0]

	zeroPoint := -lastReq.Offset

	for zeroPoint < 0 {
		zeroPoint += lastReq.Period
	}

	for zeroPoint > lastReq.Period {
		zeroPoint -= lastReq.Period
	}

	for _, req := range originalBusRequirements {
		multiplicand := (zeroPoint + req.Offset) / req.Period
		if multiplicand*req.Period-req.Offset != zeroPoint {
			fmt.Print("error")
		}

	}
	fmt.Println(zeroPoint)
}

func combineBusRequirements(requirements []BusRequirement) BusRequirement {
	first := requirements[0]
	second := requirements[1]
	// find Bezout numbers
	firstBez, secondBez := Bezout(first.Period, second.Period)

	startPoint := firstBez*first.Period*second.Offset +
		secondBez*second.Period*first.Offset

	gcm := first.Period * second.Period

	for startPoint < 0 {
		startPoint += gcm
	}

	for startPoint > gcm {
		startPoint -= gcm
	}
	return BusRequirement{gcm, startPoint}
}

// 17,x,13,19
func simpleBusSituation() []BusRequirement {
	reqs := make([]BusRequirement, 0)
	reqs = append(reqs, BusRequirement{17, 0})
	reqs = append(reqs, BusRequirement{13, 2})
	reqs = append(reqs, BusRequirement{19, 3})
	return reqs
}

func nextBusSituation() []BusRequirement {
	reqs := make([]BusRequirement, 0)
	reqs = append(reqs, BusRequirement{67, 0})
	reqs = append(reqs, BusRequirement{7, 1})
	reqs = append(reqs, BusRequirement{59, 2})
	reqs = append(reqs, BusRequirement{61, 3})
	return reqs
}

func thirdBusSituation() []BusRequirement {
	reqs := make([]BusRequirement, 0)
	reqs = append(reqs, BusRequirement{67, 0})
	reqs = append(reqs, BusRequirement{7, 2})
	reqs = append(reqs, BusRequirement{59, 3})
	reqs = append(reqs, BusRequirement{61, 4})
	return reqs
}

func fourthBusSituation() []BusRequirement {
	reqs := make([]BusRequirement, 0)
	reqs = append(reqs, BusRequirement{67, 0})
	reqs = append(reqs, BusRequirement{7, 1})
	reqs = append(reqs, BusRequirement{59, 3})
	reqs = append(reqs, BusRequirement{61, 4})
	return reqs
}

func fifthBusSituation() []BusRequirement {
	reqs := make([]BusRequirement, 0)
	reqs = append(reqs, BusRequirement{1789, 0})
	reqs = append(reqs, BusRequirement{37, 1})
	reqs = append(reqs, BusRequirement{47, 2})
	reqs = append(reqs, BusRequirement{1889, 3})
	return reqs
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
			busReq := BusRequirement{Period: num, Offset: int64(index)}
			buses = append(buses, busReq)
		}
	}
	return buses
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
