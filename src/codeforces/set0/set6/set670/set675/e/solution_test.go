package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4
4 4 4
`
	expect := 6
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
2 3 5 5
`
	expect := 17
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10
2 10 8 7 8 8 10 9 10
`
	// 1 2 3 4 5 6 7 8 9 10
	//         7 6 3 3 1 0
	// 单独考虑位置6
	// 6 - 7, 6 - 8 只需要1 * 2
	// 6 - 9 = 6 - 7 - 9 = 2
	// 6 - 10 = 6 - 7 - 10 = 2
	expect := 63
	runSample(t, s, expect)
}
