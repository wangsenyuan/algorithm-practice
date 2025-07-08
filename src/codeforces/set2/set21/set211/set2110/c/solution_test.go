package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	l, r, res := process(reader)

	expect := readString(reader)

	if expect == "-1" {
		if len(res) > 0 {
			t.Fatalf("Sample expect -1, but got %v", res)
		}
		return
	}
	if len(res) == 0 {
		t.Fatalf("Sample expect not -1, but got %v", res)
	}

	var sum int
	for i := 0; i < len(res); i++ {
		sum += res[i]
		if sum < l[i] || sum > r[i] {
			t.Fatalf("Sample expect %v, but got %v", expect, res)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4
0 -1 -1 1
0 4
1 2
2 4
1 4
0 1 1 1`)
}

func TestSample2(t *testing.T) {
	runSample(t, `3
0 -1 -1
0 1
2 2
0 3
-1`)
}

func TestSample3(t *testing.T) {
	runSample(t, `2
-1 -1
0 0
2 2
-1`)
}

func TestSample4(t *testing.T) {
	runSample(t, `8
-1 -1 1 -1 -1 0 0 -1
0 0
0 1
0 2
0 2
1 3
0 4
2 5
4 5
0 1 1 0 1 0 0 1`)
}

func TestSample5(t *testing.T) {
	runSample(t, `1
0
1 1
-1`)
}

func TestSample6(t *testing.T) {
	runSample(t, `1
-1
0 1
0`)
}
