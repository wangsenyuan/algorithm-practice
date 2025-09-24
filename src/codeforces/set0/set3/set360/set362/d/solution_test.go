package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool, correct [][]int) {

	reader := bufio.NewReader(strings.NewReader(s))
	n, roads, p, q, ok, res := drive(reader)

	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, ok)
	}

	if !expect {
		return
	}

	if len(res) != p {
		t.Fatalf("Sample expect %d, but got %d", p, len(res))
	}

	s1 := NewDSU(n)
	s2 := NewDSU(n)

	for _, cur := range roads {
		u, v := cur[0]-1, cur[1]-1
		s1.Union(u, v)
		s2.Union(u, v)
	}

	play := func(s *DSU, res [][]int) int {
		sum := make([]int, n)
		for _, cur := range roads {
			u, w := cur[0]-1, cur[2]
			u = s.Find(u)
			sum[u] += w
		}

		var add int

		for _, cur := range res {
			u, v := cur[0]-1, cur[1]-1
			u = s.Find(u)
			v = s.Find(v)
			if u == v {
				add += 1000
				continue
			}
			w := min(1e9, sum[u]+sum[v]+1)
			tmp := sum[u] + sum[v] + w
			s.Union(u, v)
			u = s.Find(u)
			sum[u] = tmp
			add += w
		}

		return add
	}

	sum1 := play(s1, correct)
	sum2 := play(s2, res)

	if sum1 != sum2 {
		t.Fatalf("Sample expect %d, but got %d", sum1, sum2)
	}

	if s2.sz != q {
		t.Fatalf("Sample expect %d, but got %d", q, s2.sz)
	}

}

func TestSample1(t *testing.T) {
	s := `9 6 2 2
1 2 2
3 2 1
4 6 20
1 3 8
7 8 3
5 7 2
`
	correct := [][]int{
		{9, 5},
		{1, 9},
	}
	runSample(t, s, true, correct)
}

func TestSample2(t *testing.T) {
	s := `2 0 1 2`

	runSample(t, s, false, nil)
}

func TestSample3(t *testing.T) {
	s := `2 0 0 2`

	runSample(t, s, true, nil)
}

func TestSample4(t *testing.T) {
	s := `2 0 0 1`
	runSample(t, s, false, nil)
}

func TestSample6(t *testing.T) {
	s := `20 11 10 4
17 12 216334157
15 19 279000438
18 20 456753771
9 15 830129118
9 14 477792844
4 11 86100846
15 4 594066440
9 6 290215734
8 5 321999322
18 1 862312250
2 3 402775619`
	correct := [][]int{
		{7, 10},
		{13, 16},
		{7, 13},
		{7, 12},
		{5, 2},
		{1, 18},
		{1, 18},
		{1, 18},
		{1, 18},
		{1, 18},
	}
	runSample(t, s, true, correct)
}
