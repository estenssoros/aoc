package day3

func part1(input string) (int, error) {
	lines := readLines(input)
	numbers := findNumbers(lines)
	var sum int
	for _, number := range numbers {
		if number.SymbolAdjacent(lines) {
			sum += number.Value
		}
	}
	return sum, nil
}
