package main

import (
	"math/rand"
	"slices"
	"testing"
	"time"
)

func hasSubseq(s, pat string) bool {
	j := 0
	for i := 0; i < len(s) && j < len(pat); i++ {
		if s[i] == pat[j] {
			j++
		}
	}
	return j == len(pat)
}

func bruteGood(s string) []int64 {
	n := len(s)
	ans := make([]int64, n)
	for l := range n {
		for r := l; r < n; r++ {
			sub := s[l : r+1]
			if hasSubseq(sub, "2025") && !hasSubseq(sub, "2026") {
				continue
			}
			for i := l; i <= r; i++ {
				ans[i]++
			}
		}
	}
	return ans
}

func TestBrute(t *testing.T) {
	check := func(s string) {
		t.Helper()
		got := solve(s)
		want := bruteGood(s)
		if !slices.Equal(got, want) {
			t.Fatalf("s=%s\ngot  %v\nwant %v", s, got, want)
		}
	}

	for _, s := range []string{
		"",
		"0",
		"2",
		"5",
		"6",
		"7",
		"2025",
		"2026",
		"20256",
		"20265",
		"202516",
		"20625",
		"72707275",
		"777",
		"1290723576",
		"20252025",
		"202560",
		"202506",
		"2222",
		"0000",
		"5555",
		"6666",
		"520252026",
		"20262025",
		"2260225",
		"20252026",
	} {
		check(s)
	}

	digits := []byte("0256")
	var gen func(prefix []byte, left int)
	gen = func(prefix []byte, left int) {
		if left == 0 {
			check(string(prefix))
			return
		}
		for _, d := range digits {
			prefix = append(prefix, d)
			gen(prefix, left-1)
			prefix = prefix[:len(prefix)-1]
		}
	}
	for length := 1; length <= 7; length++ {
		gen(nil, length)
	}

	rng := rand.New(rand.NewSource(1))
	for range 400 {
		n := 1 + rng.Intn(18)
		b := make([]byte, n)
		for i := range b {
			b[i] = byte('0' + rng.Intn(10))
		}
		check(string(b))
	}
}

func TestLarge(t *testing.T) {
	b := make([]byte, 300000)
	for i := range b {
		b[i] = byte('0' + i%10)
	}
	start := time.Now()
	ans := solve(string(b))
	if len(ans) != len(b) {
		t.Fatalf("len %d", len(ans))
	}
	for _, v := range ans {
		if v < 0 {
			t.Fatal("negative")
		}
	}
	if time.Since(start) > 2*time.Second {
		t.Fatalf("too slow: %s", time.Since(start))
	}
}
