package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	k, res := process(reader)
	expect := readString(reader)
	if len(res) > 0 != (expect == "Yes") {
		t.Fatalf("Sample expect %s, but got %v", expect, res)
	}
	if len(res) > 0 {
		var sum int
		var best int
		for _, x := range res {
			sum += x
			best = max(best, sum)
			if sum < 0 {
				sum = 0
			}
		}
		if best != k {
			t.Fatalf("Sample expect %d, but got %d", k, best)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 5
011
0 0 1
Yes`)
}

func TestSample2(t *testing.T) {
	runSample(t, `5 6
11011
4 -3 0 -2 1
Yes`)
}

func TestSample3(t *testing.T) {
	runSample(t, `4 4
0011
0 0 -4 -5
Yes`)
}

func TestSample4(t *testing.T) {
	runSample(t, `6 12
110111
1 2 0 5 -1 9
No`)
}

func TestSample5(t *testing.T) {
	runSample(t, `5 19
00000
0 0 0 0 0
Yes`)
}

func TestSample6(t *testing.T) {
	runSample(t, `5 19
11001
-8 6 0 0 -5
Yes`)
}

func TestSample7(t *testing.T) {
	runSample(t, `5 10
10101
10 0 10 0 10
Yes`)
}

func TestSample8(t *testing.T) {
	runSample(t, `1 1
1
0
No`)
}

func TestSample9(t *testing.T) {
	runSample(t, `3 5
111
3 -1 3
Yes`)
}

func TestSample10(t *testing.T) {
	runSample(t, `4 5
1011
-2 0 1 -5
Yes`)
}

func TestSample11(t *testing.T) {
	runSample(t, `2 1
11
1 -2
Yes`)
}
