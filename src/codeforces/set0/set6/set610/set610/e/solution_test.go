package main

import (
	"bufio"
	"math/rand"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := solve(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `7 4 3
abacaba
1 3 5 b
2 abc
1 4 4 c
2 cba
`
	expect := []int{6, 5}
	runSample(t, s, expect)
}

func TestRandom(t *testing.T) {
	rng := rand.New(rand.NewSource(610))
	for tc := 0; tc < 1000; tc++ {
		n := rng.Intn(12) + 1
		k := rng.Intn(5) + 1
		s := make([]byte, n)
		for i := range s {
			s[i] = byte(rng.Intn(k)) + 'a'
		}
		st := NewSegTree(s, k)
		for it := 0; it < 100; it++ {
			if rng.Intn(2) == 0 {
				l := rng.Intn(n)
				r := rng.Intn(n-l) + l
				c := byte(rng.Intn(k)) + 'a'
				for i := l; i <= r; i++ {
					s[i] = c
				}
				st.Update(l, r, int(c-'a'))
			} else {
				p := rng.Perm(k)
				perm := make([]byte, k)
				for i, x := range p {
					perm[i] = byte(x) + 'a'
				}
				got := st.Query(perm)
				want := brute(s, perm)
				if got != want {
					t.Fatalf("tc %d it %d: s=%s perm=%s, expect %d, got %d", tc, it, s, perm, want, got)
				}
			}
		}
	}
}


func brute(s []byte, perm []byte) int {
	pos := make([]int, len(perm))
	for i, x := range perm {
		pos[x-'a'] = i
	}
	res := 1
	for i := 1; i < len(s); i++ {
		if pos[s[i-1]-'a'] >= pos[s[i]-'a'] {
			res++
		}
	}
	return res
}
