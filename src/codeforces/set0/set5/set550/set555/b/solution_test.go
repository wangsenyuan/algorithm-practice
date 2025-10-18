package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	islands, a, res := drive(reader)
	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}

	m := len(a)
	marked := make([]bool, m+1)

	for i := 0; i < len(islands)-1; i++ {
		j := res[i] - 1
		if marked[j] {
			t.Fatalf("Sample result %v, use %d multiples times", res, j)
		}
		marked[j] = true
		if islands[i+1][0]-islands[i][1] > a[j] || islands[i+1][1]-islands[i][0] < a[j] {
			t.Fatalf("Sample result %v, use %d, but %d is not a valid bridge", res, j, a[j])
		}
	}
}

func TestSample1(t *testing.T) {
	s := `4 4
1 4
7 8
9 10
12 14
4 5 3 8
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 2
11 14
17 18
2 9
`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 1
1 1
1000000000000000000 1000000000000000000
999999999999999999
`
	expect := true
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `6 9
1 4
10 18
23 29
33 43
46 57
59 77
11 32 32 19 20 17 32 24 32
`
	expect := true
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `5 9
1 2
3 3
5 7
11 13
14 20
2 3 4 10 6 2 6 9 5
`
	expect := true
	runSample(t, s, expect)
}
