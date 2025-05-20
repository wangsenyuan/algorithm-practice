package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	cnt, res := process(reader)
	if cnt != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, cnt)
	}

	n := len(res)
	m := len(res[0])

	marked := make([][]bool, n)

	for i := 0; i < n; i++ {
		marked[i] = make([]bool, m)
	}
	// 因为shape是从最后的那个点开始算起的
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if marked[i][j] || res[i][j] == '.' {
				continue
			}

			found := false
			for s := range shape {
				ok := true
				for _, cur := range shape[s] {
					x, y := i+cur[0], j+cur[1]
					if x < 0 || x >= n || y < 0 || y >= m || res[x][y] != res[i][j] {
						// 这个shape肯定不符合
						ok = false
						break
					}
				}
				if ok {
					for _, cur := range shape[s] {
						x, y := i+cur[0], j+cur[1]
						marked[x][y] = true
					}
					cnt--
					found = true
					break
				}
			}
			if !found {
				t.Fatalf("Sample result %v, not correct", res)
			}
		}
	}
	if cnt != 0 {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	s := "3 3"
	expect := 1
	runSample(t, s, expect)
}


func TestSample2(t *testing.T) {
	s := "5 6"
	expect := 4
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "2 6"
	expect := 0
	runSample(t, s, expect)
}