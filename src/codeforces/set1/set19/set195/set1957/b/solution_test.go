package main

import (
	"bufio"
	"math/bits"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	k, res := drive(reader)

	check := func(arr []int) int {
		var sum int
		var or int
		for _, x := range arr {
			sum += x
			or |= x
		}
		if sum != k {
			t.Fatalf("Sample expect %v, but got %v", expect, arr)
		}
		return or
	}

	w := check(expect)
	v := check(res)

	if bits.OnesCount(uint(w)) != bits.OnesCount(uint(v)) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `6 51`
	expect := []int{3, 1, 1, 32, 2, 12}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `60 67`
	expect := make([]int, 60)
	expect[0] = 63
	expect[1] = 4
	runSample(t, s, expect)
}

