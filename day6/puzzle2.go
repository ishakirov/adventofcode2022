package main

import (
	"io/ioutil"
	"log"
)

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	answer := -1
	windowLength := 14
	window := make(map[byte]int)

	for i := 0; i < len(input); i++ {
		window[input[i]]++
		if i >= windowLength {
			window[input[i-windowLength]]--
			if window[input[i-windowLength]] == 0 {
				delete(window, input[i-windowLength])
			}

			if len(window) == windowLength {
				answer = i + 1
				break
			}
		}

	}

	log.Println(answer)
}
