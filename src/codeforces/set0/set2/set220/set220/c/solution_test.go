package main

import (
	"bufio"
	"math/rand/v2"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	expect := readNNums(reader, len(res))

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2
1 2
2 1
1 0`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `4
2 1 3 4
3 4 2 1
2 1 0 1`

	// 2 1 3 4
	// 1 3 4 2

	// 2 1 3 4
	// 2 1 3 4

	// 2 1 3 4
	// 4 2 1 3

	runSample(t, s)
}

func TestSample3(t *testing.T) {
	n := 100
	a := make([]int, n)
	b := make([]int, n)
	for i := range n {
		a[i] = i + 1
		b[i] = i + 1
	}
	c := make([]int, n)
	d := make([]int, n)
	for range 10 {
		rand.Shuffle(n, func(i, j int) {
			b[i], b[j] = b[j], b[i]
		})
		copy(c, a)
		copy(d, b)
		expect := bruteForce(c, d)

		copy(c, a)
		copy(d, b)
		res := solve(c, d)
		if !reflect.DeepEqual(res, expect) {
			t.Fatalf("Sample %v %v expect %v, but got %v", a, b, expect, res)
		}
	}
}

func TestSample4(t *testing.T) {
	s := `10
5 1 6 2 8 3 4 10 9 7
3 1 10 6 8 5 2 7 9 4
0 0 1 0 1 0 1 2 0 1`
	runSample(t, s)
}
