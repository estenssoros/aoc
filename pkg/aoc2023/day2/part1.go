package day2

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

var part1Constraints = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func part1(input string) (int, error) {
	var validGames int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ":")
		g, err := parseGame(fields[0], strings.TrimSpace(fields[1]))
		if err != nil {
			return 0, errors.Wrap(err, "parseGame")
		}
		if g.IsValid(part1Constraints) {
			validGames += g.ID
		}

	}
	return validGames, nil
}

func parseGame(game, drawsString string) (*Game, error) {
	var id int
	_, err := fmt.Sscanf(game, "Game %d", &id)
	if err != nil {
		return nil, errors.Wrap(err, "fmt.Sscanf")
	}
	g := &Game{ID: id}
	rounds := strings.Split(drawsString, "; ")
	for _, roundString := range rounds {
		round, err := parseRound(roundString)
		if err != nil {
			return nil, errors.Wrap(err, "parseRound")
		}
		g.Rounds = append(g.Rounds, round)
	}
	return g, nil
}

func parseRound(line string) (map[string]int, error) {
	fields := strings.Split(line, ", ")
	out := map[string]int{}
	for _, field := range fields {
		var count int
		var color string
		_, err := fmt.Sscanf(field, "%d %s", &count, &color)
		if err != nil {
			return nil, errors.Wrap(err, "fmt.Sscanf")
		}
		out[color] = count
	}
	return out, nil
}
