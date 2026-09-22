package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := drive(reader)

	assert.Equal(t, expect, res)
}

func TestSample1(t *testing.T) {
	s := `5 5
25957 6405 15770 26287 26465 
2 2 1
3 4 1
4 5 1
1 2 2
4 4 1
`
	expect := []int{6405, 15770, 26287, 25957, 26287}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 1
25957 6405 15770 26287 26465 
3 4 1
`
	expect := []int{15770}
	runSample(t, s, expect)
}

func TestDuplicatesNeedPositionBIT(t *testing.T) {
	runSample(t, `4 1
1 1 1 2
1 4 2
`, []int{1})
}
