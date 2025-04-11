package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `6 6 3
1 3 6
1 2
2 3
4 2
5 6
4 5
3 4
1 6
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6 5 3
1 5 6
1 2
2 3
3 4
4 5
6 3
1 5
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6 14 1
1
1 5
3 2
3 4
3 5
4 2
6 2
4 1
4 6
2 5
6 5
3 1
3 6
5 4
1 6
1 2
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `7 10 3
1 3 4
1 2
4 2
7 5
4 5
7 1
2 5
7 2
3 7
3 2
5 1
4 6
`
	expect := -1
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `6 10 1
1
3 4
4 1
1 5
4 2
6 1
4 5
6 5
6 4
6 3
2 1
1 2
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `10 33 2
3 6
1 6
8 1
4 6
10 6
6 3
7 8
9 8
3 2
5 8
1 10
5 10
8 3
1 5
9 10
2 9
6 9
10 2
5 6
3 4
9 1
1 4
9 3
10 7
6 7
8 4
6 2
1 7
6 8
1 2
4 9
3 7
9 7
4 5
3 1
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `6 5 4
1 3 4 6
1 2
2 3
2 5
5 6
4 5
1 4
`
	expect := 3
	runSample(t, s, expect)
}
