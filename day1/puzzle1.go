package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func sum(a []int) int {
	re := 0
	for _, v := range a {
		re += v
	}

	return re
}

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

	currentElfCalories := make([]int, 0)
	maxCalories := 0

	for scanner.Scan() {
		if len(scanner.Text()) > 0 {
			calories, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}

			currentElfCalories = append(currentElfCalories, calories)
		} else {
			maxCalories = max(maxCalories, sum(currentElfCalories))
			currentElfCalories = currentElfCalories[:0]
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if len(currentElfCalories) > 0 {
		maxCalories = max(maxCalories, sum(currentElfCalories))
	}

	log.Println(maxCalories)
}
