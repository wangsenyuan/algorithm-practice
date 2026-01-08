package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	cnt, res := drive(reader)
	if res == expect {
		return
	}
	if expect == "Impossible" || res == "Impossible" {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}

	freq := make([]int, 2)
	for i := range res {
		freq[int(res[i]-'0')]++
	}

	cnt2 := make([]int, 4)
	for i := range len(res) {
		if res[i] == '0' {
			freq[0]--
			cnt2[0] += freq[0]
			cnt2[1] += freq[1]
		} else {
			freq[1]--
			cnt2[2] += freq[0]
			cnt2[3] += freq[1]
		}
	}

	if !slices.Equal(cnt2, cnt) {
		t.Fatalf("Sample result %s, not giving the expected frequency %v, but got %v", res, cnt, cnt2)
	}
}

func TestSample1(t *testing.T) {
	s := `1 2 3 4`
	expect := "Impossible"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `1 2 2 1`
	expect := "0110"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `0 0 0 0`
	expect := "0"
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `10 7 28 21`
	expect := "011111110000"
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `49995000 1061 8939 0`
	expect := "011111110000"
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `1 0 0 0`
	expect := "00"
	runSample(t, s, expect)
}
