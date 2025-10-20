package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	M, T, requests, tot, assign := drive(reader)

	if tot != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, tot)
	}
	if expect == 0 {
		return
	}
	n := len(requests)
	diff := make([]int, ONE_DAY+1)

	last := make([]int, tot+1)
	for i := range tot + 1 {
		last[i] = -1
	}

	for i := range n {
		t := parseTime(requests[i])
		id := assign[i]
		if last[id] < 0 || last[id]+T <= t {
			diff[t]++
			diff[t+T]--
		} else {
			diff[last[id]+T]++
			diff[t+T]--
		}

		last[id] = t
	}

	ok := false
	for i := range ONE_DAY {
		if i > 0 {
			diff[i] += diff[i-1]
		}
		if diff[i] > M {
			t.Fatalf("Sample result %v, not correct, too many users (%d) online at %d", assign, diff[i], i)
		}
		if diff[i] == M {
			ok = true
		}
	}
	if !ok {
		t.Fatalf("Sample result %v, not correct, no moment has %d users online", assign, M)
	}
}

func TestSample1(t *testing.T) {
	s := `4 2 10
17:05:53
17:05:58
17:06:01
22:39:47
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `1 2 86400
00:00:00
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10 3 30000
00:06:54
00:42:06
03:49:45
04:38:35
05:33:30
05:51:46
10:46:34
14:34:59
14:40:06
14:53:13
`
	expect := 6
	runSample(t, s, expect)
}
