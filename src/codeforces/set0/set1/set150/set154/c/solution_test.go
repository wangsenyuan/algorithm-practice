package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int64) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 3
1 2
2 3
1 3
`
	expect := int64(3)
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 0
`
	expect := int64(3)
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4 1
1 3
`
	expect := int64(2)
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `10 9
8 3
3 1
3 4
4 10
4 7
7 3
8 4
10 3
4 1
`
	expect := int64(13)
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `20 50
2 15
20 1
3 16
3 4
15 4
19 3
11 6
1 15
18 19
12 9
12 16
15 12
2 12
4 12
15 14
14 18
7 11
15 3
6 7
15 20
19 15
16 2
12 3
18 3
19 2
20 16
7 13
2 3
18 16
18 12
14 3
6 13
16 14
19 12
14 9
9 15
12 14
1 16
11 13
19 14
9 19
3 9
14 4
19 16
4 16
16 9
2 14
16 15
4 19
15 18
`
	expect := int64(26)
	runSample(t, s, expect)
}
