package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type cpu struct {
	x           int
	cycleNumber int

	signalStrengthSum int
}

func (c *cpu) noop() {
	c.incrementCycleNumber()
}

func (c *cpu) addx(x int) {
	c.incrementCycleNumber()
	c.incrementCycleNumber()

	c.x += x
}

func (c *cpu) incrementCycleNumber() {
	c.cycleNumber++
	if (c.cycleNumber-20)%40 == 0 {
		c.signalStrengthSum += c.x * c.cycleNumber
	}
}

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	cpu := cpu{1, 0, 0}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		if scanner.Text() == "noop" {
			cpu.noop()
		} else {
			tokens := strings.Split(scanner.Text(), " ")
			toAdd, _ := strconv.Atoi(tokens[1])

			cpu.addx(toAdd)
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println(cpu.signalStrengthSum)
}
