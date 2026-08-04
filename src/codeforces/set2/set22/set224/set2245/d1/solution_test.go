package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expectOK bool) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	var n, m int
	fmt.Fscan(reader, &n, &m)
	ops := make([][]int, m)
	for i := range m {
		ops[i] = make([]int, 3)
		fmt.Fscan(reader, &ops[i][0], &ops[i][1], &ops[i][2])
	}

	ok, a := solve(n, ops)
	if ok != expectOK {
		t.Fatalf("Sample expect ok=%v, but got ok=%v a=%v", expectOK, ok, a)
	}
	if !ok {
		return
	}
	if len(a) != n {
		t.Fatalf("Sample expect array length %d, but got %v", n, a)
	}
	for _, op := range ops {
		o, i, j := op[0], op[1]-1, op[2]-1
		sum := a[i] + a[j]
		if o == 1 && sum < 0 {
			t.Fatalf("Sample array %v violates non-negative restriction on (%d, %d)", a, i+1, j+1)
		}
		if o == 2 && sum >= 0 {
			t.Fatalf("Sample array %v violates negative restriction on (%d, %d)", a, i+1, j+1)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1 1
1 1 1
`, true)
}

func TestSample2(t *testing.T) {
	runSample(t, `1 1
2 1 1
`, true)
}

func TestSample3(t *testing.T) {
	runSample(t, `2 3
1 1 1
1 1 2
1 2 2
`, true)
}

func TestSample4(t *testing.T) {
	runSample(t, `2 3
1 1 1
1 2 2
2 1 2
`, false)
}

func TestSample5(t *testing.T) {
	runSample(t, `3 6
1 1 1
1 1 2
1 1 3
2 2 2
2 2 3
2 3 3
`, true)
}

func TestSample6(t *testing.T) {
	runSample(t, `3 6
2 1 1
1 1 2
2 2 3
1 3 3
1 2 2
2 1 3
`, false)
}
