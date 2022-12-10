package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
)

func unique(s string) string {
	m := make(map[byte]struct{})
	re := bytes.Buffer{}

	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; !ok { // if byte not in map yet
			re.WriteByte(s[i])
			m[s[i]] = struct{}{}
		}
	}

	return re.String()
}

func findCommonItem(racksacks []string) byte {
	uniqueRacksacks := make([]string, 0, len(racksacks))
	for _, racksack := range racksacks {
		uniqueRacksacks = append(uniqueRacksacks, unique(racksack))
	}

	m := make(map[byte]int)
	for _, racksack := range uniqueRacksacks {
		for i := 0; i < len(racksack); i++ {
			m[racksack[i]]++
			if m[racksack[i]] == len(racksacks) {
				return racksack[i]
			}
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
	groupRacksacks := make([]string, 0, 3)
	for scanner.Scan() {
		groupRacksacks = append(groupRacksacks, scanner.Text())
		if len(groupRacksacks) == cap(groupRacksacks) {
			badge := findCommonItem(groupRacksacks)
			prioritiesSum += getItemPriority(badge)

			groupRacksacks = groupRacksacks[:0]
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println(prioritiesSum)
}
