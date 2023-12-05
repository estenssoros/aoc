package day2

import (
	"bufio"
	"strings"

	"github.com/pkg/errors"
)

func part2(input string) (int, error) {
	var out int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ":")
		g, err := parseGame(fields[0], strings.TrimSpace(fields[1]))
		if err != nil {
			return 0, errors.Wrap(err, "parseGame")
		}
		power := g.Power()
		out += power
	}
	return out, nil
}
