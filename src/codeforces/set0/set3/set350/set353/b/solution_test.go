package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	a, cnt, res := process(reader)
	expect := readNum(reader)
	if cnt != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, cnt)
	}

	p1 := make(map[int]int)
	p2 := make(map[int]int)
	for i, v := range res {
		if v == 1 {
			p1[a[i]]++
		} else {
			p2[a[i]]++
		}
	}

	if len(p1)*len(p2) != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, len(p1)*len(p2))
	}
}

func TestSample1(t *testing.T) {
	s := `1
10 99
1`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `2
13 24 13 45
4`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `50
49 13 81 20 73 62 19 49 65 95 32 84 24 96 51 57 53 83 40 44 26 65 78 80 92 87 87 95 56 46 22 44 69 80 41 61 97 92 58 53 42 78 53 19 47 36 25 77 65 81 14 61 38 99 27 58 67 37 67 80 77 51 32 43 31 48 19 79 31 91 46 97 91 71 27 63 22 84 73 73 89 44 34 84 70 23 45 31 56 73 83 38 68 45 99 33 83 86 87 80
1936`
	runSample(t, s)
}
