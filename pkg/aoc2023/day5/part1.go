package day5

import (
	"math"
)

func part1(input string) (int, error) {
	var out = math.MaxInt
	almanac := parseAlmanac(input)
	seedDestinations := almanac.SeedDestinations()
	for _, seedDestination := range seedDestinations {
		if seedDestination < out {
			out = seedDestination
		}
	}
	return out, nil
}
