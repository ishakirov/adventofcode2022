package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	n, m := 0, 0
	var grid [][]byte
	var visibilityGrid [][]bool

	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		tmp := make([]byte, len(scanner.Bytes()))
		copy(tmp, scanner.Bytes())
		grid = append(grid, tmp)
		visibilityGrid = append(visibilityGrid, make([]bool, len(scanner.Bytes())))
	}

	n, m = len(grid), len(grid[0])

	for i := 0; i < n; i++ {
		maxHeight := byte(0)
		for j := 0; j < m; j++ {
			if grid[i][j] > maxHeight {
				visibilityGrid[i][j] = true
				maxHeight = grid[i][j]
			}
		}

		maxHeight = byte(0)
		for j := m - 1; j >= 0; j-- {
			if grid[i][j] > maxHeight {
				visibilityGrid[i][j] = true
				maxHeight = grid[i][j]
			}
		}
	}

	for j := 0; j < m; j++ {
		maxHeight := byte(0)
		for i := 0; i < n; i++ {
			if grid[i][j] > maxHeight {
				visibilityGrid[i][j] = true
				maxHeight = grid[i][j]
			}
		}

		maxHeight = byte(0)
		for i := n - 1; i >= 0; i-- {
			if grid[i][j] > maxHeight {
				visibilityGrid[i][j] = true
				maxHeight = grid[i][j]
			}
		}
	}

	answer := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if visibilityGrid[i][j] {
				answer++
			}
		}
	}

	log.Println(answer)
}
