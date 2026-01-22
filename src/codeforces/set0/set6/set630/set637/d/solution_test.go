package main

import (
	"bufio"
	"slices"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func runSample(t *testing.T, in string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(in))
	m, s, d, obstacles, res := drive(reader)
	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}
	slices.Sort(obstacles)

	n := len(obstacles)

	var pos int
	var run int
	for _, cmd := range res {
		if strings.HasPrefix(cmd, "RUN") {
			if run > 0 {
				t.Fatalf("result should be run, jump, run, ...")
			}
			run, _ = strconv.Atoi(strings.TrimPrefix(cmd, "RUN "))
			r := sort.SearchInts(obstacles, pos+run)
			if r > 0 && obstacles[r-1] >= pos {
				t.Fatalf("when runing from %d to %d, there is an obstacle at %d", pos, pos+run, obstacles[r-1])
			}
			pos += run
		} else {
			if run < s {
				t.Fatalf("when jump, run should be at least %d", s)
			}
			dist, _ := strconv.Atoi(strings.TrimPrefix(cmd, "JUMP "))
			if dist > d {
				t.Fatalf("when jump, distance should be at most %d", d)
			}
			pos += dist
			r := sort.SearchInts(obstacles, pos)
			if r < n && obstacles[r] == pos {
				t.Fatalf("when jump, there is an obstacle at %d", obstacles[r])
			}
			run = 0
		}
	}
	if pos != m {
		t.Fatalf("result should end at %d, but got %d", m, pos)
	}
}

func TestSample1(t *testing.T) {
	s := `3 10 1 3
3 4 7
	`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 9 2 3
6 4
	`
	expect := false
	runSample(t, s, expect)
}
