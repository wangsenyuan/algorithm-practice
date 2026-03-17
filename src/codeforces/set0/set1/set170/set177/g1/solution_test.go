package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `6 5
a
b
ab
ba
aba`
	expect := []int{3, 5, 3, 3, 1}
	runSample(t, s, expect)
}

func TestSmallBrute(t *testing.T) {
	fib := []string{"", "a", "b"}
	for len(fib) <= 8 {
		n := len(fib)
		fib = append(fib, fib[n-1]+fib[n-2])
	}

	patterns := []string{"a", "b", "ab", "ba", "aba", "bab", "bb", "aa"}

	count := func(s, p string) int {
		var res int
		for i := 0; i+len(p) <= len(s); i++ {
			if s[i:i+len(p)] == p {
				res++
			}
		}
		return res
	}

	for k := int64(1); k <= 8; k++ {
		got := solve(k, patterns)
		want := make([]int, len(patterns))
		for i, p := range patterns {
			want[i] = count(fib[k], p)
		}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("k=%d expect %v, but got %v", k, want, got)
		}
	}
}
