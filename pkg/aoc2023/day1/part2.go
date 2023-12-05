package day1

import (
	"bufio"
	"fmt"
	"strings"

	_ "embed"

	"github.com/spf13/cobra"
)

var numberStrings = []string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

//go:embed part2sample.txt
var part2Sample string

var part2Cmd = &cobra.Command{
	Use:     "part2",
	Short:   "",
	PreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	RunE:    func(cmd *cobra.Command, args []string) error { return part2(input) },
}

func part2(input string) error {
	var ttl int
	trie := NewTrie(numberStrings)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		digits := getDigitsStrings(trie, line)
		ttl += firstAndLastDigits(digits)
	}
	fmt.Println(ttl)
	return nil
}

func getDigitsStrings(trie *Trie, line string) []int {
	digits := []int{}
	for i := 0; i < len(line); i++ {
		if isDigit(rune(line[i])) {
			digits = append(digits, int(line[i]-'0'))
			continue
		}
		if number, found := trie.Lookup(line[i:]); found {
			digits = append(digits, number)
			// i += len(numberStrings[number]) - 1
		}
	}
	return digits
}

type Trie struct {
	Children map[rune]*Trie
	Value    *int
}

func (t *Trie) Insert(number string, value int) {
	if len(number) == 0 {
		t.Value = &value
		return
	}
	char := rune(number[0])
	if _, ok := t.Children[char]; !ok {
		t.Children[char] = &Trie{
			Children: map[rune]*Trie{},
		}
	}
	t.Children[char].Insert(number[1:], value)
}

func (t *Trie) Lookup(line string) (int, bool) {
	if t.Value != nil {
		return *t.Value, true
	}
	if len(line) == 0 {
		return 0, false
	}
	char := rune(line[0])
	if _, ok := t.Children[char]; !ok {
		return 0, false
	}
	return t.Children[char].Lookup(line[1:])
}

func NewTrie(numbers []string) *Trie {
	trie := &Trie{
		Children: map[rune]*Trie{},
	}
	for i, number := range numbers {
		trie.Insert(number, i)
	}
	return trie
}
