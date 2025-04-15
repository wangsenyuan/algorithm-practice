package main

import (
	"math/bits"
	"testing"
)

func runSample(t *testing.T, a int, b int, xor int, expect bool) {
	res := solve(a, b, xor)
	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}
	x, y := res[0], res[1]

	if bits.OnesCount(uint(x)) != a {
		t.Fatalf("x has %d ones, but expect %d", bits.OnesCount(uint(x)), a)
	}
	if bits.OnesCount(uint(y)) != b {
		t.Fatalf("y has %d ones, but expect %d", bits.OnesCount(uint(y)), b)
	}

	if x^y != xor {
		t.Fatalf("x ^ y = %d, but expect %d", x^y, xor)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 4, 7, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 34, 56, 998244353, false)
}

func TestSample3(t *testing.T) {
	runSample(t, 39, 47, 530423800524412070, true)
}
