package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	out, err := part1(test)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 8, out)
}

func TestPart2(t *testing.T) {
	out, err := part2(test)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 2286, out)
}
