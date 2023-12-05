package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDigit(t *testing.T) {
	assert.Equal(t, true, isDigit('1'))
}
