package day1

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var part1Cmd = &cobra.Command{
	Use:     "part1",
	Short:   "",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE:    func(cmd *cobra.Command, args []string) error { return part1(getInput()) },
}

func part1(input string) error {
	var ttl int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		digits := getDigits(line)
		ttl += firstAndLastDigits(digits)
	}
	fmt.Println(ttl)
	return nil
}

func firstAndLastDigits(digits []int) int {
	return digits[0]*10 + digits[len(digits)-1]
}

func getDigits(input string) []int {
	var digits []int
	for _, char := range input {
		if isDigit(char) {
			digits = append(digits, int(char-'0'))
		}
	}
	return digits
}

func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}
