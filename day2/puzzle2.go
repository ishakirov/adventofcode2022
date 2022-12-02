package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var (
	figuresCodes = map[string]string{
		"A": "R",
		"B": "P",
		"C": "S",
		"X": "R",
		"Y": "P",
		"Z": "S",
	}
	figureScore = map[string]int{
		"R": 1,
		"P": 2,
		"S": 3,
	}
	rules = map[string][2]string{
		"R": {"S", "P"},
		"P": {"R", "S"},
		"S": {"P", "R"},
	}
)

func calculateScore(a, b string) int {
	p1 := figuresCodes[a]
	roundScore := 0

	switch b {
	case "X":
		roundScore += 0 + figureScore[rules[p1][0]]
	case "Y":
		roundScore += 3 + figureScore[p1]
	case "Z":
		roundScore += 6 + figureScore[rules[p1][1]]
	}

	return roundScore
}

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	totalScore := 0

	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		totalScore += calculateScore(tokens[0], tokens[1])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println(totalScore)
}
