package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	i, j int
}

type Grid struct {
	g [][]byte
}

func NewGrid(n, m int) *Grid {
	g := make([][]byte, n)
	for i := 0; i < n; i++ {
		g[i] = make([]byte, m)
		for j := 0; j < m; j++ {
			g[i][j] = '.'
		}
	}

	return &Grid{g}
}

func (g *Grid) ApplyPath(path []Point) {
	for i := 1; i < len(path); i++ {
		if path[i-1].i == path[i].i && path[i-1].j != path[i].j {
			start, end := path[i-1].j, path[i].j
			if start > end {
				start, end = end, start
			}

			for j := start; j <= end; j++ {
				g.g[path[i].i][j] = '#'
			}
		} else if path[i-1].i != path[i].i && path[i-1].j == path[i].j {
			start, end := path[i-1].i, path[i].i
			if start > end {
				start, end = end, start
			}

			for j := start; j <= end; j++ {
				g.g[j][path[i].j] = '#'
			}
		} else {
			panic("wat")
		}
	}
}

func (g *Grid) SimulateSand(i, j int) bool {
	currentI, currentJ, ok := g.makeMove(i, j)
	if !ok {
		return false
	}

	for ok {
		currentI, currentJ, ok = g.makeMove(currentI, currentJ)
	}

	g.g[currentI][currentJ] = 'o'

	return true
}

func (g *Grid) makeMove(i, j int) (int, int, bool) {
	if g.checkBoundaries(i+1, j) && g.g[i+1][j] == '.' {
		return i + 1, j, true
	}
	if g.checkBoundaries(i+1, j-1) && g.g[i+1][j-1] == '.' {
		return i + 1, j - 1, true
	}
	if g.checkBoundaries(i+1, j+1) && g.g[i+1][j+1] == '.' {
		return i + 1, j + 1, true
	}

	return i, j, false
}

func (g *Grid) checkBoundaries(i, j int) bool {
	return 0 <= i && i < len(g.g) && 0 <= j && j < len(g.g[i])
}

func (g *Grid) FindAbyssLevel() int {
	for i := len(g.g) - 1; i >= 0; i-- {
		rocksFound := false
		for j := 0; j < len(g.g[i]); j++ {
			if g.g[i][j] == '#' {
				rocksFound = true
				break
			}
		}

		if rocksFound {
			return i
		}
	}

	return 0
}

func (g *Grid) AddFloor() {
	floorLevel := g.FindAbyssLevel() + 2

	for j := 0; j < len(g.g[floorLevel]); j++ {
		g.g[floorLevel][j] = '#'
	}
}

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	grid := NewGrid(200, 1200)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		path := make([]Point, 0)

		coordinates := strings.Split(scanner.Text(), " -> ")
		for _, c := range coordinates {
			coord := strings.Split(c, ",")
			i, _ := strconv.Atoi(coord[0])
			j, _ := strconv.Atoi(coord[1])

			path = append(path, Point{j, i})
		}

		grid.ApplyPath(path)
	}

	grid.AddFloor()

	sandCount := 0
	for grid.SimulateSand(0, 500) {
		sandCount++
	}

	log.Println(sandCount + 1)
}
