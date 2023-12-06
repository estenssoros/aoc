package day5

import (
	"fmt"
	"math"
)

func part2(input string) (int, error) {
	var out = math.MaxInt
	almanac := parseAlmanac(input)
	seedRanges := [][]int{}
	for i := 0; i < len(almanac.Seeds); i += 2 {
		seedRanges = append(seedRanges, []int{almanac.Seeds[i], almanac.Seeds[i+1]})
	}
	trueSeeds := []int{}
	for _, seedRange := range seedRanges {
		for i := 0; i < seedRange[1]; i++ {
			trueSeeds = append(trueSeeds, seedRange[0]+i)
		}
	}
	almanac.Seeds = trueSeeds
	fmt.Printf("%d seeds\n", len(almanac.Seeds))
	seedDestinations := almanac.SeedDestinations()
	for _, seedDestination := range seedDestinations {
		if seedDestination < out {
			out = seedDestination
		}
	}
	return out, nil
}
