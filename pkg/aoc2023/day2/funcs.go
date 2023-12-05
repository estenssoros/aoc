package day2

type Game struct {
	ID     int
	Rounds []map[string]int
}

func (g *Game) IsValid(constraints map[string]int) bool {
	for _, round := range g.Rounds {
		if !isRoundValid(round, constraints) {
			return false
		}
	}
	return true
}

func (g *Game) Power() int {
	maxes := map[string]int{}
	for _, round := range g.Rounds {
		for color, count := range round {
			maxes[color] = max(maxes[color], count)
		}
	}
	var prod = 1
	for _, count := range maxes {
		prod *= count
	}
	return prod
}

func isRoundValid(round, constraints map[string]int) bool {
	for color, count := range round {
		constraint := constraints[color]
		if count > constraint {
			return false
		}
	}
	return true
}
