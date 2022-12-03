package main

import (
	"bufio"
	"log"
	"os"
)

func findCommonItem(rucksack string) byte {
	compartment1, compartment2 := rucksack[:len(rucksack)/2], rucksack[len(rucksack)/2:]
	m := make(map[byte]struct{})
	for i := 0; i < len(compartment1); i++ {
		m[compartment1[i]] = struct{}{}
	}

	for i := 0; i < len(compartment2); i++ {
		if _, ok := m[compartment2[i]]; ok {
			return compartment2[i]
		}
	}

	return 0 // should be unreachable
}

func getItemPriority(item byte) int {
	if 'a' <= item && item <= 'z' {
		return int(item - 'a' + 1)
	}
	if 'A' <= item && item <= 'Z' {
		return int(item - 'A' + 27)
	}

	return 0 // should be unreachable
}

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	prioritiesSum := 0
	for scanner.Scan() {
		commonItem := findCommonItem(scanner.Text())
		prioritiesSum += getItemPriority(commonItem)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println(prioritiesSum)
}
