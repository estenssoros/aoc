package day4

func part1(input string) (int, error) {
	var out int
	cards := parseCards(input)
	for _, card := range cards {
		winningNumbers := card.winningNumbers()
		if winningNumbers > 0 {
			var winnings = 1
			for i := 1; i < winningNumbers; i++ {
				winnings *= 2
			}
			out += winnings
		}
	}
	return out, nil
}
