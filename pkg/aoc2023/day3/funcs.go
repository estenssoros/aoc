package day3

import (
	"bufio"
	"strconv"
	"strings"
)

var directions = []Coordinate{
	{-1, 0}, //left
	{1, 0},  //right
	{0, -1}, //down
	{0, 1},  //up
	{1, 1},
	{1, -1},
	{-1, -1},
	{-1, 1},
}

func readLines(input string) [][]byte {
	scanner := bufio.NewScanner(strings.NewReader(input))
	lines := [][]byte{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, []byte(line))
	}
	return lines
}

func findNumbers(lines [][]byte) []*Number {
	numbers := []*Number{}
	for i, line := range lines {
		numbers = append(numbers, parseNumbers(i, line)...)
	}
	return numbers
}

func IsDigit(r byte) bool {
	return r >= '0' && r <= '9'
}

func IsSymbol(r byte) bool {
	if IsDigit(r) {
		return false
	}
	if r == '.' {
		return false
	}
	return true
}

type Number struct {
	Value       int
	Coordinates []*Coordinate
}

func (n *Number) SymbolAdjacent(lines [][]byte) bool {
	for _, coord := range n.Coordinates {
		if coord.SymbolAdjacent(lines) {
			return true
		}
	}
	return false
}

type Coordinate struct {
	X int
	Y int
}

func (c *Coordinate) adjacentNumbers(numbers []*Number) []*Number {
	out := []*Number{}
	for _, number := range numbers {
		if c.isNumberAdjacent(number) {
			out = append(out, number)
		}
	}
	return out
}

func (c *Coordinate) isNumberAdjacent(number *Number) bool {
	for _, nc := range number.Coordinates {
		if c.isAdjacent(nc) {
			return true
		}
	}
	return false
}

func (c *Coordinate) isAdjacent(other *Coordinate) bool {
	for _, direction := range directions {
		newX, newY := c.X+direction.X, c.Y+direction.Y
		if other.X == newX && other.Y == newY {
			return true
		}
	}
	return false
}

func (c *Coordinate) SymbolAdjacent(lines [][]byte) bool {
	maxY, maxX := len(lines)-1, len(lines[0])-1
	for _, direction := range directions {
		newX, newY := c.X+direction.X, c.Y+direction.Y
		if newX < 0 || newY < 0 {
			continue
		}
		if newX > maxX || newY > maxY {
			continue
		}
		if IsSymbol(lines[newY][newX]) {
			return true
		}
	}
	return false
}

func parseNumbers(row int, line []byte) []*Number {
	numbers := []*Number{}
	var current []byte
	coords := []*Coordinate{}
	for i := 0; i < len(line); i++ {
		if IsDigit(line[i]) {
			current = append(current, line[i])
			coords = append(coords, &Coordinate{i, row})
		} else {
			if len(current) > 0 {
				val, err := strconv.Atoi(string(current))
				if err != nil {
					panic(err)
				}
				numbers = append(numbers, &Number{val, coords})
				current = nil
				coords = nil
			}
		}
	}
	if len(current) > 0 {
		val, err := strconv.Atoi(string(current))
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, &Number{val, coords})
	}
	return numbers
}
