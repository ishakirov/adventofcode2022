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

func insertDescending(a []int, n int) {
	pos := -1
	for i, v := range a {
		if n > v {
			pos = i
			break
		}
	}

	if pos != -1 {
		tmp := a[pos]
		a[pos] = n
		for i := pos + 1; i < len(a); i++ {
			tmp, a[i] = a[i], tmp
		}
	}
}

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	currentElfCalories := make([]int, 0)
	maxCalories := []int{0, 0, 0}

	for scanner.Scan() {
		if len(scanner.Text()) > 0 {
			calories, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}

			currentElfCalories = append(currentElfCalories, calories)
		} else {
			insertDescending(maxCalories, sum(currentElfCalories))
			currentElfCalories = currentElfCalories[:0]
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if len(currentElfCalories) > 0 {
		insertDescending(maxCalories, sum(currentElfCalories))
	}

	log.Println(sum(maxCalories))
}
