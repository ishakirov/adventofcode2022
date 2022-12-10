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
	answer := 0

	for i := 1; i < n-1; i++ {
		for j := 1; j < m-1; j++ {

			left, right, up, down := -1, -1, -1, -1

			// left
			for k := j - 1; k >= 0; k-- {
				if grid[i][k] >= grid[i][j] {
					left = j - k
					break
				}
			}
			if left == -1 {
				left = j
			}

			// right
			for k := j + 1; k < m; k++ {
				if grid[i][k] >= grid[i][j] {
					right = k - j
					break
				}
			}
			if right == -1 {
				right = m - j - 1
			}

			// down
			for k := i + 1; k < n; k++ {
				if grid[k][j] >= grid[i][j] {
					down = k - i
					break
				}
			}
			if down == -1 {
				down = n - i - 1
			}

			// up
			for k := i - 1; k >= 0; k-- {
				if grid[k][j] >= grid[i][j] {
					up = i - k
					break
				}
			}
			if up == -1 {
				up = i
			}

			score := left * right * up * down
			if score > answer {
				answer = score
			}

		}
	}

	log.Println(answer)
}
