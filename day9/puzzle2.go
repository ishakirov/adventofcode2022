package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

func (p point) isAdjacent(other point) bool {
	return abs(p.x-other.x) <= 1 && abs(p.y-other.y) <= 1
}

func (p point) add(other point) point {
	return point{p.x + other.x, p.y + other.y}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

var (
	up    = point{0, 1}
	down  = point{0, -1}
	left  = point{-1, 0}
	right = point{1, 0}
)

func getDirection(label string) point {
	var direction point
	switch label {
	case "U":
		direction = up
	case "D":
		direction = down
	case "L":
		direction = left
	case "R":
		direction = right
	}

	return direction
}

func moveTailToHead(head, tail point) point {
	if !head.isAdjacent(tail) {
		x, y := head.x-tail.x, head.y-tail.y
		if x > 0 {
			x = 1
		} else if x < 0 {
			x = -1
		}
		if y > 0 {
			y = 1
		} else if y < 0 {
			y = -1
		}

		return point{tail.x + x, tail.y + y}
	}

	return tail
}

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	head := point{0, 0}
	tails := make([]point, 9)

	visited := make(map[point]struct{})
	visited[tails[0]] = struct{}{}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		steps, _ := strconv.Atoi(tokens[1])
		direction := getDirection(tokens[0])

		for i := 0; i < steps; i++ {
			head = head.add(direction)
			previousTail := head
			for j := 0; j < len(tails); j++ {
				tails[j] = moveTailToHead(previousTail, tails[j])
				previousTail = tails[j]
				if j == len(tails)-1 {
					visited[tails[j]] = struct{}{}
				}
			}

		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println(len(visited))
}
