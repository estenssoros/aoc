package day4

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Card struct {
	ID          int
	HaveNumbers map[int]struct{}
	WantNumbers map[int]struct{}
}

func (c Card) String() string {
	return fmt.Sprint(c.ID)
}

func (c *Card) winningNumbers() int {
	var out int
	for have := range c.HaveNumbers {
		if _, ok := c.WantNumbers[have]; ok {
			out++
		}
	}
	return out
}

func parseCards(input string) []*Card {
	cards := []*Card{}
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		cards = append(cards, parseCard(line))
	}
	return cards
}

func parseCard(line string) *Card {
	fields := strings.Split(line, ":")
	var id int
	_, err := fmt.Sscanf(fields[0], "Card %d", &id)
	if err != nil {
		panic(err)
	}
	fields = strings.Split(fields[1], "|")
	for i, field := range fields {
		fields[i] = strings.TrimSpace(field)
	}
	return &Card{
		ID:          id,
		HaveNumbers: parseNumbers(fields[0]),
		WantNumbers: parseNumbers(fields[1]),
	}
}

func parseNumbers(line string) map[int]struct{} {
	fields := strings.Split(line, " ")
	out := map[int]struct{}{}
	for _, field := range fields {
		field = strings.TrimSpace(field)
		if field == "" {
			continue
		}
		i, err := strconv.Atoi(field)
		if err != nil {
			panic(err)
		}
		_, ok := out[i]
		if ok {
			panic("number repeated twice")
		}
		out[i] = struct{}{}
	}
	return out
}
