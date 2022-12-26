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

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	cubes := make(map[Point]struct{})

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		coords := strings.Split(scanner.Text(), ",")
		X, _ := strconv.Atoi(coords[0])
		Y, _ := strconv.Atoi(coords[1])
		Z, _ := strconv.Atoi(coords[2])

		cubes[Point{X, Y, Z}] = struct{}{}
	}

	answer := 0
	for point := range cubes {
		if _, ok := cubes[Point{point.X + 1, point.Y, point.Z}]; !ok {
			answer++
		}
		if _, ok := cubes[Point{point.X - 1, point.Y, point.Z}]; !ok {
			answer++
		}
		if _, ok := cubes[Point{point.X, point.Y + 1, point.Z}]; !ok {
			answer++
		}
		if _, ok := cubes[Point{point.X, point.Y - 1, point.Z}]; !ok {
			answer++
		}
		if _, ok := cubes[Point{point.X, point.Y, point.Z + 1}]; !ok {
			answer++
		}
		if _, ok := cubes[Point{point.X, point.Y, point.Z - 1}]; !ok {
			answer++
		}
	}

	log.Println(answer)
}
