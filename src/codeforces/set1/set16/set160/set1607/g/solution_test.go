package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	m, dishes, best, res := drive(reader)

	if best != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, best)
	}

	sum := make([]int, 2)
	for i, cur := range dishes {
		x, y := res[i][0], res[i][1]
		if x > cur[0] || y > cur[1] {
			t.Fatalf("Sample result %v, not correct, it eats too much at %dth dish", res, i)
		}
		if x+y != m {
			t.Fatalf("Sample result %v, not correct, it eats too little at %dth dish", res, i)
		}
		sum[0] += cur[0] - x
		sum[1] += cur[1] - y
	}

	if abs(sum[0]-sum[1]) != expect {
		t.Fatalf("Sample result %v, not correct, it doesn't get the best balance %d, instead %d", res, expect, abs(sum[0]-sum[1]))
	}
}

func TestSample1(t *testing.T) {
	s := `1 5
3 4`
	expect := 0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `1 6
3 4`
	expect := 1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2 2
1 3
4 2`
	expect := 0
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `2 4
1 3
1 7`
	expect := 2
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `3 6
1 7
1 8
1 9`
	expect := 3
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `3 6
1 8
1 9
30 10`
	expect := 7
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `3 4
3 1
3 2
4 1`
	expect := 0
	runSample(t, s, expect)
}

func TestSample8(t *testing.T) {
	s := `5 4
0 7
6 4
0 8
4 1
5 3`
	expect := 0
	runSample(t, s, expect)
}
