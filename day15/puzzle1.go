package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type Point struct {
	X, Y int
}

type Interval struct {
	A, B int
}

func (i Interval) Contains(x int) bool {
	return i.A <= x && x <= i.B
}

func (i Interval) Length() int {
	return abs(i.B-i.A) + 1
}

type Sensor struct {
	Position      Point
	ClosestBeacon Point
}

func manhattanDistance(a, b Point) int {
	return abs(a.X-b.X) + abs(a.Y-b.Y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

var (
	MessageRegex = regexp.MustCompile(`Sensor\sat\sx=(-?\d+),\sy=(-?\d+):\sclosest\sbeacon\sis\sat\sx=(-?\d+),\sy=(-?\d+)`)
)

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	sensors := make([]Sensor, 0)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		matches := MessageRegex.FindAllStringSubmatch(scanner.Text(), -1)
		if matches == nil {
			panic("parsing error")
		}

		sensorX, _ := strconv.Atoi(matches[0][1])
		sensorY, _ := strconv.Atoi(matches[0][2])
		beaconX, _ := strconv.Atoi(matches[0][3])
		beaconY, _ := strconv.Atoi(matches[0][4])

		sensors = append(sensors, Sensor{
			Position:      Point{sensorX, sensorY},
			ClosestBeacon: Point{beaconX, beaconY},
		})
	}

	rowY := 2000000
	intervals := make([]Interval, 0)

	for _, s := range sensors {
		d := manhattanDistance(s.Position, s.ClosestBeacon)
		if abs(s.Position.Y-rowY) <= d {
			m := abs(s.Position.Y - rowY)
			leftX := s.Position.X - (d - m)
			rightX := s.Position.X + (d - m)

			intervals = append(intervals, Interval{leftX, rightX})
		}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].A < intervals[j].A
	})

	newIntervals := make([]Interval, 0)
	currentInterval := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if currentInterval.B >= intervals[i].A {
			currentInterval.B = max(intervals[i].B, currentInterval.B)
		} else {
			newIntervals = append(newIntervals, currentInterval)
			currentInterval = intervals[i]
		}
	}
	newIntervals = append(newIntervals, currentInterval)

	answer := 0
	for _, v := range newIntervals {
		answer += v.Length()
	}

	removed := make(map[int]struct{})
	for _, s := range sensors {
		if s.ClosestBeacon.Y == rowY {
			for _, interval := range newIntervals {
				if interval.Contains(s.ClosestBeacon.X) {
					if _, ok := removed[s.ClosestBeacon.X]; !ok {
						answer--
						removed[s.ClosestBeacon.X] = struct{}{}
					}
				}
			}
		}
	}

	log.Println(answer)
}
