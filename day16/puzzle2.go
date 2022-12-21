package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type vertex struct {
	flowRate int
	edges    []*edge
	opened   bool
}

type edge struct {
	to     string
	weight int
}

type graph map[string]*vertex

func (g graph) AddVertex(name string, flowRate int) {
	if _, ok := g[name]; !ok {
		g[name] = &vertex{flowRate: flowRate, edges: make([]*edge, 0), opened: false}
	}
}

func (v *vertex) AddEdge(to string, weight int) {
	v.edges = append(v.edges, &edge{to, weight})
}

var (
	messageRegex = regexp.MustCompile(`^Valve\s(\w+)\shas\sflow\srate=(\d+);\stunnels?\sleads?\sto\svalves?\s(.+)$`)
)

func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func bfsLengths(g graph, v string) map[string]int {
	queue := []string{v}
	visited := make(map[string]struct{})
	lengths := make(map[string]int)
	lengths[v] = 0
	visited[v] = struct{}{}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, edge := range g[curr].edges {
			if _, ok := visited[edge.to]; !ok {
				lengths[edge.to] = lengths[curr] + 1
				visited[edge.to] = struct{}{}
				queue = append(queue, edge.to)
			}
		}
	}

	return lengths
}

type bitmask struct {
	labels  map[string]int
	bitmask int64
}

func (s *bitmask) add(v string) {
	if _, ok := s.labels[v]; !ok {
		s.labels[v] = len(s.labels)
	}

	s.bitmask = s.bitmask | (1 << s.labels[v])
}

func (s *bitmask) delete(v string) {
	if _, ok := s.labels[v]; ok {
		s.bitmask = s.bitmask & ^(1 << s.labels[v])
	}
}

func (s *bitmask) has(v string) bool {
	_, ok := s.labels[v]

	return ok && s.bitmask&(1<<s.labels[v]) != 0
}

func (s *bitmask) bits() int64 {
	return s.bitmask
}

func (s *bitmask) String() string {
	re := ""
	labels := make([]string, 0, len(s.labels))
	for k := range s.labels {
		labels = append(labels, k)
	}
	sort.Slice(labels, func(i, j int) bool { return labels[i] < labels[j] })

	for _, k := range labels {
		re += fmt.Sprintf("%s %v;", k, s.has(k))
	}

	return re
}

type state struct {
	v           string
	timeLeft    int
	valveOpened int64
}

var memo map[state]int = make(map[state]int)

func dfsMaxFlow2(g graph, v string, timeLeft int, valveOpened *bitmask) int {
	state := state{v, timeLeft, valveOpened.bits()}
	if res, ok := memo[state]; ok {
		return res
	}

	valveOpened.add(v)

	flow := 0
	if g[v].flowRate > 0 {
		timeLeft--
		flow = g[v].flowRate * timeLeft
	}

	max := 0
	for _, e := range g[v].edges {
		if !valveOpened.has(e.to) && e.weight <= timeLeft+2 {
			max = maxInt(dfsMaxFlow2(g, e.to, timeLeft-e.weight, valveOpened), max)
		}
	}

	valveOpened.delete(v)

	memo[state] = flow + max
	return memo[state]
}

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	g := graph{}
	valvesWithFlow := make([]string, 0)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		matches := messageRegex.FindAllStringSubmatch(scanner.Text(), -1)
		if matches == nil {
			panic(fmt.Sprintf("invalid message [%s]", scanner.Text()))
		}

		name := matches[0][1]
		flowRate, _ := strconv.Atoi(matches[0][2])
		adjs := strings.Split(matches[0][3], ", ")

		g.AddVertex(name, flowRate)
		for _, v := range adjs {
			g[name].AddEdge(v, 1)
		}

		if flowRate > 0 {
			valvesWithFlow = append(valvesWithFlow, name)
		}
	}

	lengths := make(map[string]map[string]int)
	for _, v := range valvesWithFlow {
		lengths[v] = bfsLengths(g, v)
	}

	compactedGraph := graph{}
	for k := range lengths {
		compactedGraph.AddVertex(k, g[k].flowRate)
	}

	for k, v := range lengths {
		for label, vertex := range compactedGraph {
			if label != k {
				vertex.AddEdge(k, v[label])
			}
		}
	}

	AALenghths := bfsLengths(g, "AA")
	compactedGraph.AddVertex("AA", g["AA"].flowRate)
	for k := range compactedGraph {
		if k != "AA" {
			compactedGraph["AA"].AddEdge(k, AALenghths[k])
		}
	}

	for label, vertex := range compactedGraph {
		neighbours := make([]string, 0)
		for _, v := range vertex.edges {
			neighbours = append(neighbours, fmt.Sprintf("%s(len:%d)", v.to, v.weight))
		}

		log.Printf("%s(flowRate:%d)\t[%s]", label, vertex.flowRate, strings.Join(neighbours, ","))
	}

	log.Println()

	bitmask := &bitmask{make(map[string]int), 0}
	bitmask.add("AA")
	for k := range compactedGraph {
		bitmask.add(k)
	}
	answer := 0
	for i := 0; i < 1<<15; i++ {
		bitmask.bitmask = int64((i << 1)) | int64(1)
		me := dfsMaxFlow2(compactedGraph, "AA", 26, bitmask)

		bitmask.bitmask = ^(int64((i << 1)) | int64(1))
		elephant := dfsMaxFlow2(compactedGraph, "AA", 26, bitmask)

		answer = maxInt(answer, me+elephant)
	}

	log.Println(answer)
}
