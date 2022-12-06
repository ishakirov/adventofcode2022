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

	for i := 0; i < len(input)-4; i++ {
		if input[i] != input[i+1] && input[i] != input[i+2] && input[i] != input[i+3] &&
			input[i+1] != input[i+2] && input[i+1] != input[i+3] &&
			input[i+2] != input[i+3] {

			answer = i + 4
			break

		}
	}

	log.Println(answer)
}
