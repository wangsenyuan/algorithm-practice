package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	best, _ := drive(reader)
	if best != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, best)
	}
}

func TestSample1(t *testing.T) {
	s := `5 5
2 1
4 5
2 3
1 3
3 5`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7 2
3 7
4 2`
	expect := 3
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `9 17
3 6
2 6
6 9
4 1
2 8
1 9
7 9
8 5
1 7
4 9
6 7
3 4
9 3
8 4
2 1
3 8
2 7`
	expect := 7
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `5 6
2 5
3 4
1 3
4 5
5 3
2 3`
	expect := 3
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `12 8
10 2
9 2
6 9
10 6
8 2
4 10
11 2
4 11`
	expect := 10
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `19 10
6 2
3 12
17 7
2 19
17 4
1 13
7 1
13 7
6 8
11 7`
	expect := 13
	runSample(t, s, expect)
}
