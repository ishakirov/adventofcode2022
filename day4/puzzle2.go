package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func overlaps(a1, a2, b1, b2 int) bool {
	return !(a2 < b1 || a1 > b2)
}

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	answer := 0

	for scanner.Scan() {
		pairs := strings.Split(scanner.Text(), ",")
		pair1 := strings.Split(pairs[0], "-")
		pair2 := strings.Split(pairs[1], "-")

		a1, _ := strconv.Atoi(pair1[0])
		a2, _ := strconv.Atoi(pair1[1])
		b1, _ := strconv.Atoi(pair2[0])
		b2, _ := strconv.Atoi(pair2[1])

		if overlaps(a1, a2, b1, b2) {
			answer += 1
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println(answer)
}
