package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func scanInputState(input []string) [][]byte {
	re := make([][]byte, 0)

	columnPositions := make([]int, 0)
	legend := input[len(input)-1]

	for i := 0; i < len(legend); i++ {
		if legend[i] != ' ' {
			columnPositions = append(columnPositions, i)
			re = append(re, make([]byte, 0))
		}
	}

	for i := len(input) - 2; i >= 0; i-- {
		for j, pos := range columnPositions {
			if pos < len(input[i]) && input[i][pos] != ' ' {
				re[j] = append(re[j], input[i][pos])
			}
		}
	}

	return re
}

func applyCommand(state [][]byte, from, to, n int) {
	toMove := state[from][len(state[from])-n:]

	state[from] = state[from][:len(state[from])-n]
	state[to] = append(state[to], toMove...)
}

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	// reading crates state (empty line is a delimeter)
	inputState := make([]string, 0)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		inputState = append(inputState, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	state := scanInputState(inputState)
	commandRegex := regexp.MustCompile(`move\s(\d+)\sfrom\s(\d+)\sto\s(\d+)`)

	for scanner.Scan() {
		matches := commandRegex.FindAllStringSubmatch(scanner.Text(), -1)
		if matches == nil {
			log.Fatalf("command %s cannot be parsed", scanner.Text())
		}

		n, _ := strconv.Atoi(matches[0][1])
		from, _ := strconv.Atoi(matches[0][2])
		to, _ := strconv.Atoi(matches[0][3])

		applyCommand(state, from-1, to-1, n)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	answer := ""
	for i := 0; i < len(state); i++ {
		answer += string(state[i][len(state[i])-1])
	}
	log.Println(answer)
}
