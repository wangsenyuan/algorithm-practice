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
1 3 1 4
4 3 2 2`

	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6
2 1 5 3 6 4
3 2 4 5 1 6`

	expect := 3
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2
1 2
2 1`

	expect := 0
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `6
2 5 1 3 6 4
3 5 2 3 4 6`

	expect := 4
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `4
1 3 2 2
2 1 3 4`

	expect := 3
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `8
3 1 4 6 2 2 5 7
4 2 3 7 1 1 6 5`

	expect := 5
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `10
5 1 2 7 3 9 4 10 6 8
6 2 3 6 4 10 5 1 7 9`

	expect := 6
	runSample(t, s, expect)
}

func TestSample8(t *testing.T) {
	s := `5
3 2 4 1 5
2 4 5 1 3`

	expect := 4
	runSample(t, s, expect)
}

func TestSample9(t *testing.T) {
	s := `7
2 2 6 4 1 3 5
3 1 6 5 1 4 2`

	expect := 5
	runSample(t, s, expect)
}

func TestSample10(t *testing.T) {
	s := `5
4 1 3 2 5
3 2 1 5 4`

	expect := 2
	runSample(t, s, expect)
}
