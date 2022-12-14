package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Packet interface {
	Compare(other Packet) int
}

type IntegerPacket struct {
	Value int
}

func NewIntegerPacket(value int) *IntegerPacket {
	return &IntegerPacket{value}
}

func (p *IntegerPacket) Compare(other Packet) int {
	switch otherPacket := other.(type) {
	case *IntegerPacket:
		return compare(p.Value, otherPacket.Value)
	case *ListPacket:
		return p.ConvertToListPacket().Compare(otherPacket)
	default:
		panic("Unexpected packet type")
	}
}

func (p *IntegerPacket) ConvertToListPacket() *ListPacket {
	return &ListPacket{[]Packet{p}}
}

type ListPacket struct {
	Values []Packet
}

func NewListPacket() *ListPacket {
	return &ListPacket{}
}

func (p *ListPacket) Compare(other Packet) int {
	switch otherPacket := other.(type) {
	case *IntegerPacket:
		return p.Compare(otherPacket.ConvertToListPacket())
	case *ListPacket:
		for i := 0; i < min(len(p.Values), len(otherPacket.Values)); i++ {
			compare := p.Values[i].Compare(otherPacket.Values[i])
			if compare != 0 {
				return compare
			}
		}

		return compare(len(p.Values), len(otherPacket.Values))
	default:
		panic("Unexpected packet type")
	}
}

func ParsePacketFromString(s string) (Packet, error) {
	var p Packet
	stack := make([]*ListPacket, 0)

	ptr := 0
	for ptr < len(s) {
		switch true {
		case s[ptr] == '[':
			newListPacket := NewListPacket()
			stack = append(stack, newListPacket)
			if p == nil {
				p = stack[len(stack)-1]
			} else {
				stack[len(stack)-2].Values = append(stack[len(stack)-2].Values, newListPacket)
			}
			ptr++
		case s[ptr] == ']':
			stack = stack[:len(stack)-1]
			ptr++
		case isDigit(s[ptr]):
			digits := make([]byte, 0)
			for isDigit(s[ptr]) && ptr < len(s) {
				digits = append(digits, s[ptr])
				ptr++
			}

			number, err := strconv.Atoi(string(digits))
			if err != nil {
				return p, err
			}

			if p == nil {
				p = NewIntegerPacket(number)
				break
			}

			stack[len(stack)-1].Values = append(stack[len(stack)-1].Values, NewIntegerPacket(number))
		case s[ptr] == ',':
			ptr++
		}
	}

	return p, nil
}

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}

func compare(a, b int) int {
	if a < b {
		return -1
	} else if a == b {
		return 0
	} else {
		return 1
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func PrintPacket(p Packet) string {
	switch packet := p.(type) {
	case *IntegerPacket:
		return strconv.Itoa(packet.Value)
	case *ListPacket:
		packetStrings := make([]string, 0, len(packet.Values))
		for _, v := range packet.Values {
			packetStrings = append(packetStrings, PrintPacket(v))
		}

		return "[" + strings.Join(packetStrings, ",") + "]"
	default:
		panic("Unexpected packet type")
	}
}

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	var p1, p2 Packet
	pairIndex := 0
	answer := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		pairIndex++

		p1, _ = ParsePacketFromString(scanner.Text())
		scanner.Scan()
		p2, _ = ParsePacketFromString(scanner.Text())
		scanner.Scan()

		// log.Println(pairIndex, PrintPacket(p1), PrintPacket(p2), p1.Compare(p2))
		if p1.Compare(p2) == -1 {
			answer += pairIndex
		}

	}

	log.Println(answer)
}
