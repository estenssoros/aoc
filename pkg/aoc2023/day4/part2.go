package day4

func part2(input string) (int, error) {
	cards := parseCards(input)
	out := []*Card{}
	queue := make([]*Card, len(cards))
	// for i, card := range cards {
	// 	out = append(out, card)
	// 	numWinning := card.winningNumbers()
	// 	for j := i + 1; j < i+numWinning; j++ {
	// 		queue = append(queue, cards[j])
	// 	}
	// }
	// for len(queue) > 0 {

	// }
	copy(queue, cards)
	// for i, card := range cards {
	// 	queue[i] = card
	// }

	for len(queue) > 0 {
		card := queue[0]
		queue = queue[1:]
		out = append(out, card)
		numWinning := card.winningNumbers()
		for i := 0; i < numWinning; i++ {

			queue = append(queue, cards[card.ID+i])
		}

	}

	return len(out), nil
}
