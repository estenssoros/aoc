package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDigit(t *testing.T) {
	assert.Equal(t, true, IsDigit('1'))
	assert.Equal(t, false, IsDigit('='))
}

func TestPart1(t *testing.T) {
	out, err := part1(test)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 4361, out)
}

func TestPart2(t *testing.T) {
	out, err := part2(test)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 467835, out)
}
