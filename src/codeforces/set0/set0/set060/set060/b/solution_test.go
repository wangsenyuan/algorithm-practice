package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := drive(reader)

	assert.Equal(t, expect, res)
}

func TestSampel1(t *testing.T) {
	s := `1 1 1

.

1 1
`
	expect := 1
	runSample(t, s, expect)
}

func TestSampel2(t *testing.T) {
	s := `3 2 2

#.
##

#.
.#

..
..

1 2
`
	expect := 7
	runSample(t, s, expect)
}
