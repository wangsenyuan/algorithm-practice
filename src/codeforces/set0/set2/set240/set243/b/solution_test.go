package main

import (
	"bufio"
	"sort"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	found, _, _, heads, tails, _, h, w, _ := process(reader)
	if found != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, found)
	}
	if !expect {
		return
	}
	if len(heads) != h || len(tails) != w {
		t.Fatalf("Sample expect %d heads and %d tails, but got %v and %v", h, w, heads, tails)
	}
	sort.Ints(heads)
	sort.Ints(tails)
	for i, j := 0, 0; i < h && j < w; {
		if heads[i] == tails[j] {
			t.Fatalf("Sample result not correct, heads[%d] = tails[%d] = %d", i, j, heads[i])
		}
		if heads[i] < tails[j] {
			i++
		} else {
			j++
		}
	}

}

func TestSample1(t *testing.T) {
	runSample(t, `15 50 5 5
6 7
4 6
4 13
1 9
7 15
8 15
3 1
8 13
10 7
13 12
6 10
3 5
10 14
13 5
4 2
12 14
5 9
13 1
7 11
9 8
7 3
1 11
2 1
9 11
2 11
14 1
2 6
6 13
10 1
8 2
6 14
9 13
10 4
13 7
13 15
4 9
12 2
8 6
1 12
15 3
2 7
14 2
7 12
8 7
13 10
9 7
4 15
3 6
12 10
1 6
	`, true)
}

func TestSample2(t *testing.T) {
	runSample(t, `9 12 2 3
1 2
2 3
1 3
1 4
2 5
4 5
4 6
6 5
6 7
7 5
8 7
9 1
	`, true)
}


func TestSample3(t *testing.T) {
	runSample(t, `10 20 4 4
7 1
1 8
8 5
2 4
8 2
7 2
8 3
6 1
10 1
4 8
3 1
5 2
7 10
6 7
9 8
4 10
2 3
1 9
5 4
5 9
	`, true)
}
