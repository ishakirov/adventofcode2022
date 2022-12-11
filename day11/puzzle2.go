package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var (
	ExpressionRegex = regexp.MustCompile(`new\s=\s(old|\d+)\s([*+])\s(old|\d+)`)
)

type Monkey struct {
	Number            int
	Items             []int
	Operation         *Operation
	TestDivisor       int
	MonkeyIfTestTrue  int
	MonkeyIfTestFalse int

	InspectedItems int
}

type Operation struct {
	const1, const2       int
	isOldVar1, isOldVar2 bool
	op                   byte
}

func (o *Operation) Execute(old int) int {
	var arg1, arg2 int
	if o.isOldVar1 {
		arg1 = old
	} else {
		arg1 = o.const1
	}
	if o.isOldVar2 {
		arg2 = old
	} else {
		arg2 = o.const2
	}

	switch o.op {
	case '+':
		return arg1 + arg2
	case '*':
		return arg1 * arg2
	default:
		panic(fmt.Sprintf("Unexpected operation %s", string(o.op)))
	}
}

func ParseOperationFromString(expression string) (*Operation, error) {
	matches := ExpressionRegex.FindAllStringSubmatch(expression, -1)
	if matches == nil {
		return nil, errors.New(fmt.Sprintf("Expression [%s] cannot be parsed", expression))
	}

	re := &Operation{}
	var err error

	arg1, op, arg2 := matches[0][1], matches[0][2], matches[0][3]
	if arg1 == "old" {
		re.isOldVar1 = true
	} else {
		re.const1, err = strconv.Atoi(arg1)
		if err != nil {
			panic(err)
		}
	}
	if arg2 == "old" {
		re.isOldVar2 = true
	} else {
		re.const2, err = strconv.Atoi(arg2)
		if err != nil {
			panic(err)
		}
	}

	re.op = op[0]

	return re, nil
}

func startsWith(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	monkeys := make([]*Monkey, 0)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		if startsWith(scanner.Text(), "Monkey") {
			monkey := Monkey{}

			tokens := strings.Split(scanner.Text(), " ")
			monkey.Number, _ = strconv.Atoi(tokens[1][:len(tokens[1])-1])

			scanner.Scan() // Starting items
			tokens = strings.Split(scanner.Text(), ":")
			startingItemsString := strings.TrimSpace(tokens[1])
			for _, item := range strings.Split(startingItemsString, ", ") {
				v, _ := strconv.Atoi(item)
				monkey.Items = append(monkey.Items, v)
			}

			scanner.Scan() // Operation
			tokens = strings.Split(scanner.Text(), ":")
			expressionString := strings.TrimSpace(tokens[1])
			op, err := ParseOperationFromString(expressionString)
			if err != nil {
				panic(fmt.Sprintf("Couldn't parse monkey operation: %v", err))
			}
			monkey.Operation = op

			scanner.Scan() // Test
			tokens = strings.Split(scanner.Text(), ":")
			testString := strings.TrimSpace(tokens[1])
			tokens = strings.Split(testString, " ")
			monkey.TestDivisor, _ = strconv.Atoi(tokens[2])

			scanner.Scan() // If test true
			tokens = strings.Split(scanner.Text(), ":")
			testTrueString := strings.TrimSpace(tokens[1])
			tokens = strings.Split(testTrueString, " ")
			monkey.MonkeyIfTestTrue, _ = strconv.Atoi(tokens[3])

			scanner.Scan() // If test false
			tokens = strings.Split(scanner.Text(), ":")
			testFalseString := strings.TrimSpace(tokens[1])
			tokens = strings.Split(testFalseString, " ")
			monkey.MonkeyIfTestFalse, _ = strconv.Atoi(tokens[3])

			monkeys = append(monkeys, &monkey)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	commonDivisor := 1
	for _, m := range monkeys {
		commonDivisor *= m.TestDivisor
	}

	rounds := 10_000
	for round := 0; round < rounds; round++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.Items {
				monkey.InspectedItems++
				newWorryLevel := monkey.Operation.Execute(item)
				newWorryLevel %= commonDivisor

				if newWorryLevel%monkey.TestDivisor == 0 {
					monkeys[monkey.MonkeyIfTestTrue].Items = append(monkeys[monkey.MonkeyIfTestTrue].Items, newWorryLevel)
				} else {
					monkeys[monkey.MonkeyIfTestFalse].Items = append(monkeys[monkey.MonkeyIfTestFalse].Items, newWorryLevel)
				}
			}

			monkey.Items = monkey.Items[:0]
		}

		roundNumber := round + 1
		if roundNumber == 1 || roundNumber == 20 || roundNumber%1000 == 0 {
			log.Printf("Round #%d", round+1)
			for _, m := range monkeys {
				log.Printf("Monkey %d inspected items %d times", m.Number, m.InspectedItems)
			}
		}
	}

	inspectedItemsCounts := make([]int, len(monkeys))
	for i, monkey := range monkeys {
		inspectedItemsCounts[i] = monkey.InspectedItems
	}

	sort.Ints(inspectedItemsCounts)

	log.Println(inspectedItemsCounts[len(inspectedItemsCounts)-1] * inspectedItemsCounts[len(inspectedItemsCounts)-2])
}
