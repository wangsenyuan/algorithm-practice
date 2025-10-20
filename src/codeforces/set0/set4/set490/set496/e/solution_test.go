package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))

	parts, actors, res := drive(reader)
	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	// m := len(parts)
	n := len(actors)
	cnt := make([]int, n)

	for i, cur := range parts {
		j := res[i] - 1
		cnt[j]++
		actor := actors[j]
		if actor[0] > cur[0] || actor[1] < cur[1] {
			t.Fatalf("actor %d is not suitable for part %d", j+1, i+1)
		}
	}

	for i := range n {
		if cnt[i] > actors[i][2] {
			t.Fatalf("actor %d is used more than %d times", i+1, actors[i][2])
		}
	}
}

func TestSample1(t *testing.T) {
	s := `3
1 3
2 4
3 5
2
1 4 2
2 5 1
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
1 3
2 4
3 5
2
1 3 2
2 5 1
`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5
276139667 276139667
649350960 699381779
711629163 711629163
711629163 711629163
479042169 577792509
3
439900587 817601793 2
711629163 711629163 2
276139667 276139667 1
`
	expect := true
	runSample(t, s, expect)
}
