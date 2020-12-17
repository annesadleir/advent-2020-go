package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type XxxYyy struct {
	xxx int
	yyy int
}

type ZzzXxxYyy struct {
	zzz    int
	xxxYYY XxxYyy
}

type MaxMin struct {
	Min int
	Max int
}

func (maxMin MaxMin) advanceMM() MaxMin {
	return MaxMin{maxMin.Min - 1, maxMin.Max + 1}
}

type Field struct {
	zzzSpan    MaxMin
	xxxYYYSpan MaxMin
}

func (field Field) advance() Field {
	return Field{field.zzzSpan.advanceMM(), field.xxxYYYSpan.advanceMM()}
}

type Layer = map[XxxYyy]bool

type ThreeD = map[int]Layer

type SuperDimension struct {
	superCubes map[int]ThreeD
	field      Field
}
type Dimension struct {
	cubes ThreeD
	field Field
}

func (superDimension SuperDimension) activeCount() int {
	count := 0
	for _, threeD := range superDimension.superCubes {
		count += countActive(threeD)
	}
	return count
}

func countActive(threeD ThreeD) int {
	count := 0
	for _, layer := range threeD {
		count += len(layer)
	}
	return count
}

func main() {
	fmt.Println(partOneDay17()) // 353
	fmt.Println(partTwoDay17()) // 2472
}

func partTwoDay17() int {
	dimensionNow := readDay17Input()
	superCubes := make(map[int]ThreeD, 0)
	superCubes[0] = dimensionNow.cubes
	superDimension := SuperDimension{superCubes, dimensionNow.field}
	for cycle := 1; cycle <= 6; cycle++ {
		superDimension = bootCycle4D(superDimension)
	}
	return superDimension.activeCount()
}

func partOneDay17() int {
	dimensionNow := readDay17Input()
	for cycle := 1; cycle <= 6; cycle++ {
		dimensionNow = bootCycle(dimensionNow)
	}
	return countActive(dimensionNow.cubes)
}

func bootCycle4D(current SuperDimension) SuperDimension {
	newField := current.field.advance()
	newSuperCubes := make(map[int]ThreeD)

	for www := newField.zzzSpan.Min; www <= newField.zzzSpan.Max; www++ {
		thisThreeD := make(map[int]Layer, 0)
		newSuperCubes[www] = thisThreeD
		for layer := newField.zzzSpan.Min; layer <= newField.zzzSpan.Max; layer++ {
			thisLayer := make(map[XxxYyy]bool, 0)
			thisThreeD[layer] = thisLayer
			for xxx := newField.xxxYYYSpan.Min; xxx <= newField.xxxYYYSpan.Max; xxx++ {
				for yyy := newField.xxxYYYSpan.Min; yyy <= newField.xxxYYYSpan.Max; yyy++ {
					coords := XxxYyy{xxx, yyy}
					active := activeInNext4DCycle(current, www, layer, coords)
					if active {
						thisLayer[coords] = true
					}
				}
			}
		}
	}
	return SuperDimension{newSuperCubes, newField}
}

func bootCycle(current Dimension) Dimension {
	newField := current.field.advance()
	newCubes := make(map[int]Layer, 0)
	for layer := newField.zzzSpan.Min; layer <= newField.zzzSpan.Max; layer++ {
		thisLayer := make(map[XxxYyy]bool, 0)
		newCubes[layer] = thisLayer
		for xxx := newField.xxxYYYSpan.Min; xxx <= newField.xxxYYYSpan.Max; xxx++ {
			for yyy := newField.xxxYYYSpan.Min; yyy <= newField.xxxYYYSpan.Max; yyy++ {
				coords := XxxYyy{xxx, yyy}
				active := activeInNextCycle(current, layer, coords)
				if active {
					thisLayer[coords] = true
				}
			}
		}
	}
	return Dimension{newCubes, newField}
}

func activeInNext4DCycle(current SuperDimension, www int, layer int, coords XxxYyy) bool {
	countActiveSurrounding := countActiveSurroundingIn4D(current, www, layer, coords)
	thisActive := activeCubeInSuperDimension(current, www, layer, coords)
	return activeCubeInNextCycle(thisActive, countActiveSurrounding)
}

func countActiveSurroundingIn4D(current SuperDimension, www int, layer int, coords XxxYyy) int {
	surrounders := surroundingFourD(www, layer, coords)
	countActiveSurrounding := 0
	for aroundW, zzzXxxYyys := range surrounders {
		for _, zzzXxxYyy := range zzzXxxYyys {
			if activeCubeInSuperDimension(current, aroundW, zzzXxxYyy.zzz, zzzXxxYyy.xxxYYY) {
				countActiveSurrounding++
			}
		}
	}
	return countActiveSurrounding
}

func activeCubeInNextCycle(activeNow bool, activeSurrounding int) bool {
	if activeNow {
		if activeSurrounding == 2 || activeSurrounding == 3 {
			return true
		} else {
			return false
		}
	} else {
		if activeSurrounding == 3 {
			return true
		} else {
			return false
		}
	}
}

func activeInNextCycle(current Dimension, layer int, coords XxxYyy) bool {
	surrounders := surroundingCubes(layer, coords)
	countActiveSurrounding := countActiveSurrounders(current, surrounders)
	thisActive := activeCube(current, layer, coords)
	return activeCubeInNextCycle(thisActive, countActiveSurrounding)
}

func countActiveSurrounders(current Dimension, cubes []ZzzXxxYyy) int {
	count := 0
	for _, cube := range cubes {
		if activeCube(current, cube.zzz, cube.xxxYYY) {
			count++
		}
	}
	return count
}

func surroundingFourD(fourthD int, layer int, coords XxxYyy) map[int][]ZzzXxxYyy {
	results := make(map[int][]ZzzXxxYyy, 0)
	for www := fourthD - 1; www <= fourthD+1; www++ {
		zzzXxxYyys := make([]ZzzXxxYyy, 0)
		for zzz := layer - 1; zzz <= layer+1; zzz++ {
			for xxx := coords.xxx - 1; xxx <= coords.xxx+1; xxx++ {
				for yyy := coords.yyy - 1; yyy <= coords.yyy+1; yyy++ {
					if !(zzz == layer && xxx == coords.xxx && yyy == coords.yyy && www == fourthD) {
						zzzXxxYyys = append(zzzXxxYyys, ZzzXxxYyy{zzz, XxxYyy{xxx, yyy}})
					}
				}
			}
		}
		results[www] = zzzXxxYyys
	}
	return results
}

func surroundingCubes(layer int, coords XxxYyy) []ZzzXxxYyy {
	results := make([]ZzzXxxYyy, 0)
	for zzz := layer - 1; zzz <= layer+1; zzz++ {
		for xxx := coords.xxx - 1; xxx <= coords.xxx+1; xxx++ {
			for yyy := coords.yyy - 1; yyy <= coords.yyy+1; yyy++ {
				if !(zzz == layer && xxx == coords.xxx && yyy == coords.yyy) {
					results = append(results, ZzzXxxYyy{zzz, XxxYyy{xxx, yyy}})
				}
			}
		}
	}
	return results
}

func activeCubeInSuperDimension(dimension SuperDimension, www int, zzz int, coords XxxYyy) bool {
	threeD, exists := dimension.superCubes[www]
	if !exists {
		return false
	} else {
		layer, isThere := threeD[zzz]
		if !isThere {
			return false
		}
		return layer[coords]
	}
}

func activeCube(dimension Dimension, zzz int, coords XxxYyy) bool {
	layer, exists := dimension.cubes[zzz]
	if !exists {
		return false
	} else {
		return layer[coords]
	}
}

func readDay17Input() Dimension {
	file, err := os.Open("C:\\Workarea\\advent-2020-go\\inputs\\day17.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	layer := make(map[XxxYyy]bool, 0)
	xxx := 0
	yyy := 0
	for scanner.Scan() {
		line := scanner.Text()
		for xxx = 0; xxx < len(line); xxx++ {
			if line[xxx:xxx+1] == "#" {
				coords := XxxYyy{xxx, yyy}
				layer[coords] = true
			}
		}
		yyy++
	}

	dimension := make(map[int]Layer, 0)
	dimension[0] = layer
	return Dimension{dimension, Field{MaxMin{0, 0}, MaxMin{0, yyy - 1}}}
}
