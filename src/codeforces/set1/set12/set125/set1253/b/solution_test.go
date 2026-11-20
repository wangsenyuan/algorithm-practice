package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, possible bool) {
	t.Helper()
	res := drive(bufio.NewReader(strings.NewReader(s)))
	a := readEvents(s)
	if !possible {
		if res != nil {
			t.Fatalf("Sample expect impossible, but got %v", res)
		}
		return
	}
	if !valid(a, res) {
		t.Fatalf("Sample got invalid split %v", res)
	}
}

func readEvents(s string) []int {
	reader := bufio.NewReader(strings.NewReader(s))
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return a
}

func valid(a []int, days []int) bool {
	if len(days) == 0 {
		return false
	}
	sum := 0
	for _, d := range days {
		if d <= 0 {
			return false
		}
		sum += d
	}
	if sum != len(a) {
		return false
	}
	at := 0
	for _, d := range days {
		if !validDay(a[at : at+d]) {
			return false
		}
		at += d
	}
	return true
}

func validDay(events []int) bool {
	state := map[int]int{}
	inside := 0
	for _, x := range events {
		id := x
		if id < 0 {
			id = -id
		}
		if x > 0 {
			if state[id] != 0 {
				return false
			}
			state[id] = 1
			inside++
			continue
		}
		if state[id] != 1 {
			return false
		}
		state[id] = 2
		inside--
	}
	return inside == 0
}

func TestSample1(t *testing.T) {
	runSample(t, `6
1 7 -7 3 -1 -3
`, true)
}

func TestSample2(t *testing.T) {
	runSample(t, `8
1 -1 1 2 -1 -2 3 -3
`, true)
}

func TestSample3(t *testing.T) {
	runSample(t, `6
2 5 -5 5 -5 -2
`, false)
}

func TestSample4(t *testing.T) {
	runSample(t, `3
-8 1 1
`, false)
}
