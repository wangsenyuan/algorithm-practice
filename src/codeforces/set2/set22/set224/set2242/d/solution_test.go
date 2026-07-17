package main

import (
	"bufio"
	"math/rand"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect []int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1
5147
44441
`, []int{2})
}

func TestSample2(t *testing.T) {
	runSample(t, `1
21945
60
`, []int{-1})
}

func TestSample3(t *testing.T) {
	runSample(t, `1
123450
012345
`, []int{5})
}

func bruteSolve(a, b string) int {
	generate := func(s string) map[string]bool {
		res := make(map[string]bool)
		var dfs func(int, []byte)
		dfs = func(pos int, cur []byte) {
			if pos == len(s) {
				res[string(cur)] = true
				return
			}
			sum := 0
			for end := pos; end < len(s); end++ {
				sum = (sum + int(s[end]-'0')) % 10
				dfs(end+1, append(cur, byte('0'+sum)))
			}
		}
		dfs(0, nil)
		return res
	}

	left := generate(a)
	right := generate(b)
	best := -1
	for s := range left {
		if right[s] {
			best = max(best, len(s))
		}
	}
	return best
}

func TestRandomAgainstBrute(t *testing.T) {
	rng := rand.New(rand.NewSource(2242))
	for tc := 0; tc < 2000; tc++ {
		n := 1 + rng.Intn(7)
		m := 1 + rng.Intn(7)
		a := make([]byte, n)
		b := make([]byte, m)
		for i := range a {
			a[i] = byte('0' + rng.Intn(10))
		}
		for i := range b {
			b[i] = byte('0' + rng.Intn(10))
		}

		got := solve(string(a), string(b))
		want := bruteSolve(string(a), string(b))
		if got != want {
			t.Fatalf("case %d: solve(%q, %q) = %d, want %d", tc, a, b, got, want)
		}
	}
}
