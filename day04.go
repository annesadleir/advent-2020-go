package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var Part1Reqs = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

var ValidEyeColours = map[string]bool{"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true}

const HairColourPattern = "#[0-9a-f]{6}"
const PassportIdPattern = "\\b[0-9]{9}\\b"

type Passport struct {
	Values map[string]string
}

func (passport Passport) validPart1() bool {
	for _, req := range Part1Reqs {
		_, present := passport.Values[req]
		if !present {
			return false
		}
	}
	return true
}

func (passport Passport) validPart2() bool {
	validByr := passport.validYear("byr", 1920, 2002)
	ivalidIyr := passport.validYear("iyr", 2010, 2020)
	validEyr := passport.validYear("eyr", 2020, 2030)
	validHgt := passport.validHeight()
	validHcl := passport.validRegex("hcl", HairColourPattern)
	validEyc := passport.validEyeColour()
	validPid := passport.validRegex("pid", PassportIdPattern)
	allValid := validByr && ivalidIyr && validEyr && validHgt && validHcl && validEyc && validPid
	return allValid
}

func (passport Passport) validRegex(key, pattern string) bool {
	value, ok := passport.Values[key]
	if !ok {
		return false
	} else {
		match, _ := regexp.MatchString(pattern, value)
		return match
	}
}

func (passport Passport) validEyeColour() bool {
	value, ok := passport.Values["ecl"]
	if !ok {
		return false
	} else {
		return ValidEyeColours[value]
	}
}

func (passport Passport) validHeight() bool {
	value, ok := passport.Values["hgt"]
	if !ok {
		return false
	} else {
		length := len(value)
		if length < 4 {
			return false
		}
		lastTwoChars := value[length-2:]
		if lastTwoChars != "cm" && lastTwoChars != "in" {
			return false
		}
		numbers := value[:length-2]
		asInt, err := strconv.Atoi(numbers)
		if err != nil {
			return false
		}
		if lastTwoChars == "cm" && (asInt < 150 || asInt > 193) {
			return false
		}
		if lastTwoChars == "in" && (asInt < 59 || asInt > 76) {
			return false
		}
		return true
	}
}

func (passport Passport) validYear(key string, min, max int) bool {
	value, ok := passport.Values[key]
	if !ok {
		return false
	} else {
		if len(value) != 4 {
			return false
		}
		asInt, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
		return asInt >= min && asInt <= max
	}
}

func readPassports() []Passport {
	file, err := os.Open("C:\\Workarea\\advent-2020-go\\inputs\\day04.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	passports := make([]Passport, 0)
	current := Passport{make(map[string]string)}
	passports = append(passports, current)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			current = Passport{make(map[string]string)}
			passports = append(passports, current)
		} else {
			pairs := strings.Split(line, " ")
			for _, pair := range pairs {
				kv := strings.Split(pair, ":")
				current.Values[kv[0]] = kv[1]
			}
		}
	}
	return passports
}

func main() {
	passports := readPassports()
	fmt.Println("Passports read: " + strconv.Itoa(len(passports)))

	runningTotalPart1 := 0
	runningTotalPart2 := 0
	for _, passport := range passports {
		if passport.validPart1() {
			runningTotalPart1++
		}
		if passport.validPart2() {
			runningTotalPart2++
		}
	}
	fmt.Println("Part 1 answer: " + strconv.Itoa(runningTotalPart1))
	fmt.Println("Part 2 answer: " + strconv.Itoa(runningTotalPart2))
}
