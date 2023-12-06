package day5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRange(t *testing.T) {
	rng := &Range{
		Destination: 52,
		Source:      50,
		RangeLength: 48,
	}
	assert.True(t, true, rng.Has(79))
	assert.Equal(t, 81, rng.Translate(79))
}

func TestPart1(t *testing.T) {
	out, err := part1(test)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 35, out)
}

func TestPart2(t *testing.T) {
	out, err := part2(test)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 46, out)
}
