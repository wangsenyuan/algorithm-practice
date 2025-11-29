package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	n, k, res := drive(reader)

	if expect == -1 {
		if len(res) > 0 {
			t.Fatalf("Sample expect -1, but got %v", res)
		}
		return
	}

	if len(res) == 0 {
		t.Fatalf("Sample expect %v, but got -1", expect)
	}
	var g int
	var sum int
	for i := 0; i < k; i++ {
		g = gcd(g, res[i])
		if i > 0 && res[i] <= res[i-1] {
			t.Fatalf("Sample result %v, not strictly increasing", res)
		}
		sum += res[i]
	}

	if g != expect {
		t.Fatalf("Sample result %v, but got gcd %d, want %d", res, g, expect)
	}
	if sum != n {
		t.Fatalf("Sample result %v, but got sum %d, want %d", res, sum, n)
	}
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func TestSample1(t *testing.T) {
	s := "6 3"
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "8 2"
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "5 3"
	expect := -1
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "3000000021 3"
	expect := 3
	runSample(t, s, expect)
}


func TestSample5(t *testing.T) {
	s := "10000000000 100000"
	expect := 1
	runSample(t, s, expect)
}
