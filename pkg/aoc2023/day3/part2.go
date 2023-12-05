package day3

func part2(input string) (int, error) {
	lines := readLines(input)
	gears := findGears(lines)
	numbers := findNumbers(lines)
	var sum int
	for _, gear := range gears {
		adjacentNumbers := gear.adjacentNumbers(numbers)
		if len(adjacentNumbers) == 2 {
			sum += (adjacentNumbers[0].Value * adjacentNumbers[1].Value)
		}
	}
	return sum, nil
}

func findGears(lines [][]byte) []*Coordinate {
	gears := []*Coordinate{}
	for y, row := range lines {
		for x, b := range row {
			if b == '*' {
				gears = append(gears, &Coordinate{x, y})
			}
		}
	}
	return gears
}
