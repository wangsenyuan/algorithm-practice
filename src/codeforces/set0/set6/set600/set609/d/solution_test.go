package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, in string, expect int) {
	reader := bufio.NewReader(strings.NewReader(in))
	s, a, b, gadgets, d, res := drive(reader)
	if d != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, d)
	}
	if d < 0 {
		return
	}
	var sum int
	for _, cur := range res {
		i, j := cur[0]-1, cur[1]-1
		if gadgets[i][0] == 1 {
			sum += a[j] * gadgets[i][1]
		} else {
			sum += b[j] * gadgets[i][1]
		}
	}
	if sum > s {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	s := `5 4 2 2
1 2 3 2 1
3 2 1 2 3
1 1
2 1
1 2
2 2
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 3 2 200
69 70 71 72
104 105 106 107
1 1
2 2
1 2
`
	expect := -1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4 3 1 1000000000
900000 910000 940000 990000
990000 999000 999900 999990
1 87654
2 76543
1 65432
`
	expect := -1
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `5 5 3 1000000
921 853 547 187 164
711 462 437 307 246
2 94
2 230
1 373
1 476
2 880
`
	expect := 1
	runSample(t, s, expect)
}
