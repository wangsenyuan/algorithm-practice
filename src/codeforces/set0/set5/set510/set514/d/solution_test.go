package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))

	a, k, res := drive(reader)

	var sum int
	for _, v := range res {
		sum += v
	}

	if sum > k {
		t.Fatalf("it makes %d shots, exceeds %d", sum, k)
	}

	n := len(a)
	m := len(res)

	play := func(ans []int) int {
		que := make([][]pair, m)
		var best int
		for l, r := 0, 0; r < n; r++ {
			for j := range m {
				for len(que[j]) > 0 && last(que[j]).first <= a[r][j] {
					que[j] = que[j][:len(que[j])-1]
				}
				que[j] = append(que[j], pair{a[r][j], r})
			}
			for l <= r {
				ok := true
				for j := range m {
					if que[j][0].first > ans[j] {
						ok = false
						break
					}
				}
				if ok {
					break
				}
				for j := range m {
					if que[j][0].second == l {
						que[j] = que[j][1:]
					}
				}
				l++
			}
			best = max(best, r-l+1)
		}
		return best
	}

	x := play(expect)
	y := play(res)

	if x != y {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5 2 4
4 0
1 2
2 1
0 2
1 3
`
	expect := []int{2, 2}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 2 4
1 2
1 3
2 2
`
	expect := []int{1, 3}
	runSample(t, s, expect)
}
