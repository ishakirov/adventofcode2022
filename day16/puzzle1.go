package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		visited[curr] = struct{}{}

		for _, edge := range g[curr].edges {
			if _, ok := visited[edge.to]; !ok {
				lengths[edge.to] = lengths[curr] + 1
				queue = append(queue, edge.to)
			}
		}
	}

	return lengths
}

func dfsMaxflow(g graph, v string, valveOpened map[string]struct{}, timeLeft int) int {
	if timeLeft <= 0 {
		return 0
	}

	valveOpened[v] = struct{}{}

	currentFlow := g[v].flowRate * timeLeft

	maxFlow := 0
	for _, e := range g[v].edges {
		if _, ok := valveOpened[e.to]; !ok && e.weight <= timeLeft+1 {
			maxFlow = maxInt(maxFlow, dfsMaxflow(g, e.to, valveOpened, timeLeft-e.weight))
		}
	}

	delete(valveOpened, v)

	return currentFlow + maxFlow
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
	if _, ok := lengths["AA"]; !ok {
		lengths["AA"] = bfsLengths(g, "AA")
	}

	compactedGraph := graph{}
	for k := range lengths {
		compactedGraph.AddVertex(k, g[k].flowRate)
	}

	for k, v := range lengths {
		for label, vertex := range compactedGraph {
			if label != k {
				weight := v[label]
				if g[k].flowRate > 0 {
					weight++
				}
				vertex.AddEdge(k, weight)
			}
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

	start := "AA"
	totalTime := 30
	answer := 0

	for _, e := range compactedGraph[start].edges {
		answer = maxInt(answer, dfsMaxflow(compactedGraph, e.to, make(map[string]struct{}), totalTime-e.weight))
	}

	log.Println(answer)
}
