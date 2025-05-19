package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res, a := process(reader)
	ans := readString(reader)
	if ans == "YES" {
		if len(res) == 0 {
			t.Fatalf("Sample expect %s, but got %v", ans, res)
		}
		l, r := res[0], res[1]
		l--
		r--
		v := []int{a[l], a[l]}
		for i := l + 1; i <= r; i++ {
			v[0] = max(v[0], a[i])
			v[1] = min(v[1], a[i])
		}
		if v[0]-v[1] < r-l+1 {
			t.Fatalf("Sample expect %s, but got %v", ans, res)
		}
	} else if len(res) > 0 {
		t.Fatalf("Sample expect %s, but got %v", ans, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5
1 2 3 4 5
NO`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `4
2 0 1 9
YES`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `2
2019 2020
NO`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `2
0 0
NO`
	runSample(t, s)
}

func TestSample5(t *testing.T) {
	s := `2
1 0
NO`
	runSample(t, s)
}