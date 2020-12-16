package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type TicketSituation struct {
	FieldDefinitions []FieldDefinition
	MyTicket         Ticket
	NearbyTickets    []Ticket
}

type FieldDefinition struct {
	name     string
	validity []Condition
}

type Condition struct {
	min int
	max int
}

func (condition Condition) valid(num int) bool {
	return num >= condition.min && num <= condition.max
}

func (field FieldDefinition) valid(num int) bool {
	for _, condition := range field.validity {
		if condition.valid(num) {
			return true
		}
	}
	return false
}

func (situation TicketSituation) possiblyValid(num int) bool {
	for _, field := range situation.FieldDefinitions {
		if field.valid(num) {
			return true
		}
	}
	return false
}

func (situation TicketSituation) possiblyValidTicket(ticket Ticket) bool {
	for _, value := range ticket {
		if !situation.possiblyValid(value) {
			return false
		}
	}
	return true
}

type Ticket = []int

func main() {
	situation := readTicketSituation()
	partOneAnswer := partOneDay16(situation) // 22000
	fmt.Println(partOneAnswer)

	partTwoAnswer := partTwoDay16(situation)
	fmt.Println(partTwoAnswer) // 410460648673
}

func partTwoDay16(situation TicketSituation) int {
	validTickets := removeInvalidTickets(situation)
	fields := situation.FieldDefinitions
	allTickets := append(validTickets, situation.MyTicket)

	possibleIndexForFields := make(map[string][]int, 0)
	possibleFieldForIndex := make(map[int][]FieldDefinition)
	doFirstRunThrough(fields, allTickets, possibleIndexForFields, possibleFieldForIndex)

	definiteIndexForFields := make(map[string]int, 0)

	for len(definiteIndexForFields) < 20 {
		deduce(possibleFieldForIndex, possibleIndexForFields, definiteIndexForFields)
	}

	runningMultiple := 1;
	for fieldName, index := range definiteIndexForFields {
		if strings.Contains(fieldName, "departure") {
			runningMultiple *= situation.MyTicket[index]
		}
	}

	return runningMultiple
}

func deduce(indexMap map[int][]FieldDefinition, fieldsMap map[string][]int, identities map[string]int) {
	for ind, fields := range indexMap {
		if len(fields) == 1 {
			handleAFoundIdentity(ind, fields[0].name, indexMap, fieldsMap, identities)
		}
	}

	for fieldName, possible := range fieldsMap {
		if len(possible) == 1 {
			handleAFoundIdentity(possible[0], fieldName, indexMap, fieldsMap, identities)
		}
	}
}

func handleAFoundIdentity(index int, fieldName string, indexMap map[int][]FieldDefinition, fieldsMap map[string][]int, identities map[string]int) {
	identities[fieldName] = index
	delete(indexMap, index)
	delete(fieldsMap, fieldName)
	for fieldName, possIndexes := range fieldsMap {
		fieldsMap[fieldName] = withoutIndex(possIndexes, index)
	}
	for indx, fields := range indexMap {
		indexMap[indx] = withoutField(fields, fieldName)
	}
}

func withoutField(fields []FieldDefinition, name string) []FieldDefinition {
	remaining := make([]FieldDefinition, 0)
	for _, f := range fields {
		if f.name != name {
			remaining = append(remaining, f)
		}
	}
	return remaining
}

func withoutIndex(possIndexes []int, notPossible int) []int {
	stillPossible := make([]int, 0)
	for _, i := range possIndexes {
		if i != notPossible {
			stillPossible = append(stillPossible, i)
		}
	}
	return stillPossible
}

func doFirstRunThrough(fields []FieldDefinition, tickets []Ticket,
	fieldsMap map[string][]int, indexMap map[int][]FieldDefinition) {

	for index := 0; index < 20; index++ {
		possibleFields := make([]FieldDefinition, 0)
		for _, field := range fields {
			fieldCouldBeAtIndex := true
			for _, ticket := range tickets {
				if !field.valid(ticket[index]) {
					fieldCouldBeAtIndex = false
				}
			}
			if fieldCouldBeAtIndex {
				possibleFields = append(possibleFields, field)
			}
		}
		indexMap[index] = possibleFields
	}

	for _, field := range fields {
		possibleIndices := make([]int, 0)
		for ind, possFs := range indexMap {
			if contains(possFs, field) {
				possibleIndices = append(possibleIndices, (ind))
			}
		}
		fieldsMap[field.name] = possibleIndices
	}

}

func contains(fs []FieldDefinition, field FieldDefinition) bool {
	for _, f := range fs {
		if f.name == field.name {
			return true
		}
	}
	return false
}

func removeInvalidTickets(situation TicketSituation) []Ticket {
	validTickets := make([]Ticket, 0)
	for _, ticket := range situation.NearbyTickets {
		if situation.possiblyValidTicket(ticket) {
			validTickets = append(validTickets, ticket)
		}
	}
	return validTickets
}

func partOneDay16(situation TicketSituation) int {
	invalidNums := findInvalidFields(situation)

	runningTotal := 0
	for _, val := range invalidNums {
		runningTotal += val
	}
	return runningTotal
}

func findInvalidFields(situation TicketSituation) []int {
	invalidNums := make([]int, 0)
	for _, ticket := range situation.NearbyTickets {
		for _, val := range ticket {
			if !situation.possiblyValid(val) {
				invalidNums = append(invalidNums, val)
			}
		}
	}
	return invalidNums
}

func readTicketSituation() TicketSituation {
	file, err := os.Open("C:\\Workarea\\advent-2020-go\\inputs\\day16.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	section := 0

	fieldDefinitions := make([]FieldDefinition, 0)
	var myTicket Ticket
	nearbyTickets := make([]Ticket, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			section++
		} else if !strings.Contains(line, "ticket") {
			switch section {
			case 0:
				fieldDefinitions = append(fieldDefinitions, readFieldDefinition(line))
			case 1:
				myTicket = readTicket(line)
			case 2:
				nearbyTickets = append(nearbyTickets, readTicket(line))
			}
		}
	}

	return TicketSituation{fieldDefinitions, myTicket, nearbyTickets}
}

func readTicket(line string) Ticket {
	nums := strings.Split(line, ",")
	values := make([]int, 0)
	for _, num := range nums {
		n, _ := strconv.Atoi(num)
		values = append(values, n)
	}
	return values
}

func readFieldDefinition(line string) FieldDefinition {
	halves := strings.Split(line, ": ")
	name := halves[0]
	conditionDefs := strings.Split(halves[1], " or ")
	conditions := make([]Condition, 0)
	for _, condition := range conditionDefs {
		bounds := strings.Split(condition, "-")
		min, _ := strconv.Atoi(bounds[0])
		max, _ := strconv.Atoi(bounds[1])
		conditions = append(conditions, Condition{min, max})
	}
	return FieldDefinition{name, conditions}
}
