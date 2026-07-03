package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `6
-2 1
0 3
3 3
4 1
3 -2
2 -2
4
0 1
2 2
3 1
1 0
`, true)
}

func TestSample2(t *testing.T) {
	runSample(t, `5
1 2
4 2
3 -3
-2 -2
-2 1
4
0 1
1 2
4 1
2 -1
`, false)
}

func TestSample3(t *testing.T) {
	runSample(t, `5
-1 2
2 3
4 1
3 -2
0 -3
5
1 0
1 1
3 1
5 -1
2 -1
`, false)
}

func TestSample4(t *testing.T) {
	runSample(t, `5
4 3
2 -3
-1 -3
-1 0
2 2
5
-1 -2
-1 -1
2 1
3 0
2 -2
`, false)
}

func TestSample5(t *testing.T) {
	runSample(t, `4
-10 -10
-10 10
10 10
10 -10
3
0 -10
1 5
2 2
`, false)
}
