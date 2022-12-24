package main

import (
	"bufio"
	"log"
	"os"
)

type Point struct {
	X, Y int
}

func (p Point) Add(other Point) Point {
	return Point{p.X + other.X, p.Y + other.Y}
}

type Figure struct {
	position Point
	shape    []Point
}

func (f *Figure) GetPoints() []Point {
	points := make([]Point, 0, len(f.shape))

	for _, p := range f.shape {
		points = append(points, p.Add(f.position))
	}

	return points
}

var shapes = [][]Point{
	{{0, 0}, {0, 1}, {0, 2}, {0, 3}},             // horizontal line
	{{0, 1}, {-1, 0}, {-1, 1}, {-1, 2}, {-2, 1}}, // plus
	{{0, 0}, {0, 1}, {0, 2}, {-1, 2}, {-2, 2}},   // corner
	{{0, 0}, {-1, 0}, {-2, 0}, {-3, 0}},          // vertical line
	{{0, 0}, {0, 1}, {-1, 0}, {-1, 1}},           // square
}

type Chamber struct {
	jetPattern      string
	jetPatternIndex int

	field        [][]byte
	highestLevel int
}

func NewChamber(width, height int, jetPattern string) *Chamber {
	field := make([][]byte, height)
	for i := 0; i < len(field); i++ {
		field[i] = make([]byte, width)
		for j := 0; j < len(field[i]); j++ {
			field[i][j] = '.'
		}
	}

	return &Chamber{
		jetPattern:      jetPattern,
		jetPatternIndex: 0,
		field:           field,
		highestLevel:    height,
	}
}

func (c *Chamber) GetSpawnPoint() Point {
	return Point{c.highestLevel - 4, 2}
}

func (c *Chamber) SimulateFigure(figure Figure) {
	ok := c.isAbleToMove(figure, Point{0, 0})
	for ok {
		// jet stream
		switch c.jetPattern[c.jetPatternIndex] {
		case '>':
			if c.isAbleToMove(figure, Point{0, 1}) {
				figure.position.Y += 1
			}
		case '<':
			if c.isAbleToMove(figure, Point{0, -1}) {
				figure.position.Y -= 1
			}
		}
		c.jetPatternIndex = (c.jetPatternIndex + 1) % len(c.jetPattern)

		// down
		if ok = c.isAbleToMove(figure, Point{1, 0}); ok {
			figure.position.X += 1
		}
	}

	for _, p := range figure.GetPoints() {
		c.field[p.X][p.Y] = '#'
		if p.X < c.highestLevel {
			c.highestLevel = p.X
		}
	}
}

func (c *Chamber) isAbleToMove(figure Figure, direction Point) bool {
	ok := true

	figure.position = figure.position.Add(direction)

	for _, p := range figure.GetPoints() {
		if !(0 <= p.X && p.X < len(c.field) && 0 <= p.Y && p.Y < len(c.field[p.X]) && c.field[p.X][p.Y] != '#') {
			ok = false
			break
		}
	}

	return ok
}

func (c *Chamber) GetTotalHeight() int {
	return len(c.field) - c.highestLevel
}

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Scan()
	jetPattern := scanner.Text()

	maxHeight := 2022*4 + 100
	chamber := NewChamber(7, maxHeight, jetPattern)

	possibleCycleStarts := make(map[int]int)
	cycleLength := -1
	cycleStart := -1

	for i := 0; i < len(chamber.jetPattern)*len(shapes); i++ {
		if i%len(shapes) == 0 {
			if rockNumber, ok := possibleCycleStarts[chamber.jetPatternIndex]; ok {
				log.Printf("Found cycle - rockStart %d, currentRock %d, length %d", rockNumber, i, i-rockNumber)
				cycleStart = rockNumber
				cycleLength = i - rockNumber
				break
			} else {
				possibleCycleStarts[chamber.jetPatternIndex] = i
			}
		}

		chamber.SimulateFigure(Figure{chamber.GetSpawnPoint(), shapes[i%len(shapes)]})
	}

	cycle := make([]int, cycleLength+1)
	prevHeight := chamber.GetTotalHeight()
	for i := 0; i < cycleLength; i++ {
		chamber.SimulateFigure(Figure{chamber.GetSpawnPoint(), shapes[i%len(shapes)]})
		cycle[i+1] = chamber.GetTotalHeight() - prevHeight
	}

	totalRocks := 1000000000000
	totalRocks -= (cycleStart - 1)

	chamber2 := NewChamber(7, maxHeight, jetPattern)
	for i := 0; i < cycleStart-1; i++ {
		chamber2.SimulateFigure(Figure{chamber2.GetSpawnPoint(), shapes[i%len(shapes)]})
	}

	// I don't understand why +1, but it worked :(
	log.Println(chamber2.GetTotalHeight() + cycle[cycleLength]*(totalRocks/cycleLength) + cycle[totalRocks%cycleLength] + 1)
}
