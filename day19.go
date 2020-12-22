package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type RulesAndInputs struct {
	Rules  map[string]string
	Inputs []string
}

func main() {
	//fmt.Println(partOneDay19())
	partTwoDay19()
}

func partTwoDay19() {
	situation := readDay19Part2()
	fmt.Println(len(situation.Rules))
	newRules := reduceDay19PartTwo(situation.Rules)
	fmt.Println(len(newRules))
	printOutRules(newRules)
}

func printOutRules(lines map[string]map[string]bool) {

	keys := make([]int, 0, len(lines))
	for k := range lines {
		i, _ := strconv.Atoi(k)
		keys = append(keys, i)
	}
	sort.Ints(keys)

	for _, key := range keys {
		k := strconv.Itoa(key)
		line := k + ": "
		for poss := range lines[k] {
			line += poss + " | "
		}
		fmt.Println(line)
	}
}

func reduceDay19PartTwo(lines map[string]string) map[string]map[string]bool {

	rules := make(map[string]map[string]bool, 0)

	for key, value := range lines {
		parts := strings.Split(value, " | ")
		partsMap := make(map[string]bool, 0)
		for _, part := range parts {
			partsMap[part] = true
		}
		rules[key] = partsMap
	}

	hasChanged := true
	for hasChanged {
		var change1, change2, change3 bool
		rules, change1 = removeSingles(rules)
		rules, change2 = concatenateLetters(rules)
		rules, change3 = substituteResolved(rules)
		hasChanged = change1 || change2 || change3
	}
	return rules
}

func jj() {
	// do not replace with itself in loop
}

func substituteResolved(rules map[string]map[string]bool) (map[string]map[string]bool, bool) {
	changedAnything := false
	for resolvable, possibilities := range rules {
		if noDigits(possibilities) {
			for k1, v1 := range rules {
				amalgam := make(map[string]bool, 0)
				for str := range v1 {
					if str == resolvable {
						for poss := range possibilities {
							amalgam[poss] = true
							changedAnything = true
						}
					} else if strings.HasSuffix(str, " " + resolvable) {
						whereSpace := strings.LastIndex(str, " ")
						firstPart := str[:whereSpace + 1]
						for poss := range possibilities {
							newStr := firstPart + poss
							newStr = strings.Replace(newStr, "\" \"", "", -1)
							amalgam[newStr] = true
							changedAnything = true
						}
					} else if strings.HasPrefix(str, resolvable+" ") {
						whereSpace := strings.Index(str, " ")
						secondPart := str[whereSpace:]
						for poss := range possibilities {
							newStr := poss + secondPart
							newStr = strings.Replace(newStr, "\" \"", "", -1)
							amalgam[newStr] = true
							changedAnything = true
						}
					} else {
						amalgam[str] = true
					}
				}
				rules[k1] = amalgam
			}
			delete(rules, resolvable)
		}
	}
	return rules, changedAnything
}

func noDigits(possibles map[string]bool) bool {
	for k := range possibles {
		if digitInString(k) {
			return false
		}
	}
	return true
}

func concatenateLetters(rules map[string]map[string]bool) (map[string]map[string]bool, bool) {
	changedAnything := false
	for _, v := range rules {
		for str := range v {
			if strings.Contains(str, "\" \"") {
				shorter := strings.ReplaceAll(str, "\" \"", "")
				v[shorter] = true
				delete(v, str)
				changedAnything = true
			}
		}
	}
	return rules, changedAnything
}

func removeSingles(rules map[string]map[string]bool) (map[string]map[string]bool, bool) {
	changedAnything := false
	for k, v := range rules {
		if len(v) == 1 {
			onlyPossibility := singleValue(v)
			if !strings.Contains(onlyPossibility, " ") {
				for _, v1 := range rules {
					for possString := range v1 {
						if possString == k || strings.HasPrefix(possString, k+" ") || strings.HasSuffix(possString, " "+k) {
							changedString := strings.Replace(possString, k, onlyPossibility, -1)
							v1[changedString] = true
							delete(v1, possString)
							changedAnything = true
						}
					}
				}
				delete(rules, k)
			}
		}
	}
	return rules, changedAnything
}

func singleValue(in map[string]bool) string {
	if len(in) != 1 {
		panic("Single value not single")
	}
	for k := range in {
		return k
	}
	return "chaos"
}

func digitInString(input string) bool {
	for _, r := range []rune(input) {
		if unicode.IsDigit(r) {
			return true
		}
	}
	return false
}

func copyMap(m map[string]string) map[string]string {
	copy := make(map[string]string, 0)
	for k, v := range m {
		copy[k] = v
	}
	return copy
}

func partOneDay19() int {
	situation := readDay19()
	allValid := possibleMatchesMap(situation.Rules, situation.Rules["0"]) // len = 2097152
	partOne := valid(allValid, situation.Inputs)
	return partOne
}

func valid(allValid map[string]bool, inputs []string) int {
	count := 0
	for _, input := range inputs {
		exists := allValid[input]
		if exists {
			count++
		}
	}
	return count
}

func possibleMatchesMap(rules map[string]string, def string) map[string]bool {
	if strings.Contains(def, "\"") {
		letter := string(def[1])
		return map[string]bool{letter: true}
	} else if strings.Contains(def, "|") {
		parts := strings.Split(def, " | ")
		firstPart := possibleMatchesMap(rules, parts[0])
		secondPart := possibleMatchesMap(rules, parts[1])
		return concat(firstPart, secondPart)
	} else {
		parts := strings.Split(def, " ")
		numParts := len(parts)
		if numParts > 3 || numParts < 1 {
			panic("Bad assumption about number of rules " + def)
		}
		if numParts == 3 {
			matchesFirst := possibleMatchesMap(rules, parts[0])
			matchesSecond := possibleMatchesMap(rules, parts[1])
			matchesThird := possibleMatchesMap(rules, parts[2])
			matches := make(map[string]bool, 0)
			for mf := range matchesFirst {
				for ms := range matchesSecond {
					for mt := range matchesThird {
						matches[mf+ms+mt] = true
					}
				}
			}
			return matches
		} else if numParts == 2 {
			matchesFirst := possibleMatchesMap(rules, parts[0])
			matchesSecond := possibleMatchesMap(rules, parts[1])
			matches := make(map[string]bool, 0)
			for mf, _ := range matchesFirst {
				for ms, _ := range matchesSecond {
					matches[mf+ms] = true
				}
			}
			return matches
		} else {
			return possibleMatchesMap(rules, rules[parts[0]])
		}
	}
}

func concat(first map[string]bool, second map[string]bool) map[string]bool {
	both := make(map[string]bool)
	for k, v := range first {
		both[k] = v
	}
	for k, v := range second {
		both[k] = v
	}
	return both
}

func readDay19Part2() RulesAndInputs {
	file, err := os.Open("C:\\Workarea\\advent-2020-go\\inputs\\day19_2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	rules := make(map[string]string, 0)
	inputs := make([]string, 0)
	readingRules := true
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			readingRules = false
		} else if readingRules {
			parts := strings.Split(line, ": ")
			rules[parts[0]] = parts[1]

		} else {
			inputs = append(inputs, line)
		}
	}
	return RulesAndInputs{rules, inputs}
}

func readDay19() RulesAndInputs {
	file, err := os.Open("C:\\Workarea\\advent-2020-go\\inputs\\day19.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	rules := make(map[string]string, 0)
	inputs := make([]string, 0)
	readingRules := true
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			readingRules = false
		} else if readingRules {
			parts := strings.Split(line, ": ")
			rules[parts[0]] = parts[1]

		} else {
			inputs = append(inputs, line)
		}
	}
	return RulesAndInputs{rules, inputs}
}

func printRules(rules map[string]string) {
	keys := make([]int, 0, len(rules))
	for k := range rules {
		i, _ := strconv.Atoi(k)
		keys = append(keys, i)
	}
	sort.Ints(keys)

	for _, key := range keys {
		a := strconv.Itoa(key)
		fmt.Println(key, rules[a])
	}
}
