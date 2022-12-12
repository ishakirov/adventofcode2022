package main

import (
	"bufio"
	"log"
	"os"
)

type vertex struct {
	i, j int
}

func bfs(grid [][]byte, start, exit vertex) int {
	lengths := make([][]int, len(grid))
	for i := 0; i < len(lengths); i++ {
		lengths[i] = make([]int, len(grid[i]))
		for j := 0; j < len(lengths[i]); j++ {
			lengths[i][j] = -1
		}
	}

	queue := []vertex{start}
	lengths[start.i][start.j] = 0

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		// up
		if checkBoundaries(grid, v.i-1, v.j) && lengths[v.i-1][v.j] == -1 && grid[v.i-1][v.j] <= grid[v.i][v.j]+1 {
			queue = append(queue, vertex{v.i - 1, v.j})
			lengths[v.i-1][v.j] = lengths[v.i][v.j] + 1
		}

		// down
		if checkBoundaries(grid, v.i+1, v.j) && lengths[v.i+1][v.j] == -1 && grid[v.i+1][v.j] <= grid[v.i][v.j]+1 {
			queue = append(queue, vertex{v.i + 1, v.j})
			lengths[v.i+1][v.j] = lengths[v.i][v.j] + 1
		}

		// left
		if checkBoundaries(grid, v.i, v.j-1) && lengths[v.i][v.j-1] == -1 && grid[v.i][v.j-1] <= grid[v.i][v.j]+1 {
			queue = append(queue, vertex{v.i, v.j - 1})
			lengths[v.i][v.j-1] = lengths[v.i][v.j] + 1
		}

		// right
		if checkBoundaries(grid, v.i, v.j+1) && lengths[v.i][v.j+1] == -1 && grid[v.i][v.j+1] <= grid[v.i][v.j]+1 {
			queue = append(queue, vertex{v.i, v.j + 1})
			lengths[v.i][v.j+1] = lengths[v.i][v.j] + 1
		}
	}

	return lengths[exit.i][exit.j]
}

func checkBoundaries(grid [][]byte, i, j int) bool {
	return 0 <= i && i < len(grid) && 0 <= j && j < len(grid[i])
}

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	var grid [][]byte
	var start, exit vertex

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		grid = append(grid, make([]byte, len(scanner.Bytes())))
		copy(grid[len(grid)-1], scanner.Bytes())

		for j := 0; j < len(grid[len(grid)-1]); j++ {
			if grid[len(grid)-1][j] == 'S' {
				start = vertex{len(grid) - 1, j}
				grid[len(grid)-1][j] = 'a'
			}
			if grid[len(grid)-1][j] == 'E' {
				exit = vertex{len(grid) - 1, j}
				grid[len(grid)-1][j] = 'z'
			}
		}
	}

	log.Println(bfs(grid, start, exit))
}
