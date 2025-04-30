package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	score, res, a := process(reader)
	if score != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
	if len(res) != (len(a)+1)/2 {
		t.Errorf("Sample expect %d, but got %v", expect, res)
	}
	marked := make([]bool, len(a)+1)
	var sum int
	for _, cur := range res {
		if len(cur) > 2 {
			t.Fatalf("Sample expect %v, but got %v", expect, res)
		}
		if len(cur) == 2 {
			u, v := cur[0], cur[1]
			if marked[u] || marked[v] {
				t.Fatalf("Sample expect %v, but got %v", expect, res)
			}
			marked[u] = true
			marked[v] = true
			sum += max(a[u-1], a[v-1])
		} else {
			u := cur[0]
			if marked[u] {
				t.Fatalf("Sample expect %v, but got %v", expect, res)
			}
			marked[u] = true
			sum += a[u-1]
		}
	}
	if sum != score {
		t.Errorf("Sample expect %d, but got %d", expect, score)
	}
}

func TestSample1(t *testing.T) {
	s := `4
1 2 3 4
`
	expect := 6
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
2 4 3 1 4
`
	expect := 8
	runSample(t, s, expect)
}
