package day4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	out, err := part1(test)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 13, out)
}

func TestPart2(t *testing.T) {
	out, err := part2(test)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 30, out)
}
