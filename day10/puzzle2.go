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

	currentRow   string
	renderedRows []string
}

func (c *cpu) noop() {
	c.incrementCycleNumberAndRender()
}

func (c *cpu) addx(x int) {
	c.incrementCycleNumberAndRender()
	c.incrementCycleNumberAndRender()

	c.x += x
}

func (c *cpu) incrementCycleNumberAndRender() {
	c.cycleNumber++

	currentPixel := c.currentPixel()
	if c.x-1 <= currentPixel && currentPixel <= c.x+1 {
		c.currentRow += "#"
	} else {
		c.currentRow += "."
	}

	if c.currentPixel() == 39 {
		c.renderedRows = append(c.renderedRows, c.currentRow)
		c.currentRow = ""
	}
}

func (c *cpu) currentPixel() int {
	return c.cycleNumber % 40
}

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	cpu := cpu{
		x:            1,
		cycleNumber:  -1,
		currentRow:   "",
		renderedRows: make([]string, 0),
	}

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

	for _, row := range cpu.renderedRows {
		log.Println(row)
	}
}
