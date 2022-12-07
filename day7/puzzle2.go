package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type node struct {
	name          string
	fileSize      int
	isDirectory   bool
	directorySize int

	parent *node
	childs []*node
}

func startsWith(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func calculateDirectorySize(node *node) {
	size := 0
	for _, child := range node.childs {
		if child.isDirectory {
			if child.directorySize == -1 {
				calculateDirectorySize(child)
			}
			size += child.directorySize
		} else {
			size += child.fileSize
		}
	}

	node.directorySize = size
}

func findDirectoryToDelete(n *node, toDelete int) int {
	queue := make([]*node, 0)
	queue = append(queue, n)

	minDiff := int(2e9)
	answer := -1

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		if v.directorySize >= toDelete {
			if v.directorySize-toDelete < minDiff {
				minDiff = v.directorySize - toDelete
				answer = v.directorySize
			}
		}

		for _, c := range v.childs {
			if c.isDirectory {
				queue = append(queue, c)
			}
		}
	}

	return answer
}

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	root := &node{
		name:          "/",
		fileSize:      -1,
		isDirectory:   true,
		directorySize: -1,

		parent: nil,
		childs: make([]*node, 0),
	}

	currentNode := root

	for scanner.Scan() {
		if startsWith(scanner.Text(), "$") {
			command := scanner.Text()[2:]
			if startsWith(command, "cd") {
				tokens := strings.Split(command, " ")
				if tokens[1] == "/" {
					currentNode = root
				} else if tokens[1] == ".." {
					currentNode = currentNode.parent
				} else {
					dirExists := false
					for _, v := range currentNode.childs {
						if v.name == tokens[1] {
							dirExists = true
							currentNode = v
							break
						}
					}
					if !dirExists {
						newNode := &node{
							name:          tokens[1],
							fileSize:      -1,
							isDirectory:   true,
							directorySize: -1,

							parent: currentNode,
							childs: make([]*node, 0),
						}
						currentNode.childs = append(currentNode.childs, newNode)
						currentNode = newNode
					}
				}
			} else if startsWith(command, "ls") {
				// do nothing
			} else {
				log.Fatal("unexpected command", command)
			}
		} else {
			if startsWith(scanner.Text(), "dir") {
				tokens := strings.Split(scanner.Text(), " ")
				newNode := &node{
					name:          tokens[1],
					fileSize:      -1,
					isDirectory:   true,
					directorySize: -1,

					parent: currentNode,
					childs: make([]*node, 0),
				}
				currentNode.childs = append(currentNode.childs, newNode)
			} else {
				tokens := strings.Split(scanner.Text(), " ")
				fileSize, _ := strconv.Atoi(tokens[0])
				newNode := &node{
					name:          tokens[1],
					fileSize:      fileSize,
					isDirectory:   false,
					directorySize: -1,

					parent: currentNode,
					childs: make([]*node, 0),
				}
				currentNode.childs = append(currentNode.childs, newNode)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	calculateDirectorySize(root)

	totalFsSize := 70000000
	requiredSpace := 30000000

	toFree := requiredSpace - (totalFsSize - root.directorySize)
	if toFree <= 0 {
		log.Fatal("no additional space required")
	}

	log.Println(findDirectoryToDelete(root, toFree))

}
