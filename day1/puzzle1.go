package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	currentElfCalories := 0
	maxCalories := 0

	for scanner.Scan() {
		if len(scanner.Text()) > 0 {
			calories, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}

			currentElfCalories += calories
		} else {
			maxCalories = max(maxCalories, currentElfCalories)
			currentElfCalories = 0
		}
	}
	maxCalories = max(maxCalories, currentElfCalories)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println(maxCalories)
}
