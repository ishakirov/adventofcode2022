package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X, Y, Z int
}

func (p Point) Add(other Point) Point {
	return Point{p.X + other.X, p.Y + other.Y, p.Z + other.Z}
}

var (
	visited    = make(map[Point]struct{})
	cubes      = make(map[Point]struct{})
	directions = []Point{
		{1, 0, 0},
		{-1, 0, 0},
		{0, 1, 0},
		{0, -1, 0},
		{0, 0, 1},
		{0, 0, -1},
	}
)

// 0..22
func dfs(p Point) {
	visited[p] = struct{}{}

	for _, d := range directions {
		next := p.Add(d)
		if checkBoundaries(next) {
			_, pointVisited := visited[next]
			_, cube := cubes[next]

			if !pointVisited && !cube {
				dfs(next)
			}
		}
	}
}

func checkBoundaries(p Point) bool {
	return -1 <= p.X && p.X <= 22 &&
		-1 <= p.Y && p.Y <= 22 &&
		-1 <= p.Z && p.Z <= 22
}

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		coords := strings.Split(scanner.Text(), ",")
		X, _ := strconv.Atoi(coords[0])
		Y, _ := strconv.Atoi(coords[1])
		Z, _ := strconv.Atoi(coords[2])

		cubes[Point{X, Y, Z}] = struct{}{}
	}

	dfs(Point{-1, -1, -1})

	answer := 0
	for p := range cubes {
		for _, d := range directions {
			if _, ok := visited[p.Add(d)]; ok {
				answer++
			}
		}
	}

	log.Println(answer)
	log.Printf("total space: %d; points: %d; visited: %d", 23*23*23, len(cubes), len(visited))
}
