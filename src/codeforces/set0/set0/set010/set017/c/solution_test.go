package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)
	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "abca", 7)
}

func TestSample2(t *testing.T) {
	runSample(t, "ab", 1)
}

func TestSample3(t *testing.T) {
	runSample(t, "abbc", 3)
}

func TestSample4(t *testing.T) {
	runSample(t, "bacacccbbcacbcaababbcbcacbbcaacbccbabaaacbcacccccaabccacbabcacacabbcccabcababccccababacaa", 4379010)
}

func TestSmallAgainstBruteForce(t *testing.T) {
	for n := 1; n <= 7; n++ {
		total := 1
		for i := 0; i < n; i++ {
			total *= 3
		}
		for mask := 0; mask < total; mask++ {
			s := makeString(mask, n)
			expect := bruteForce(s)
			res := solve(s)
			if res != expect {
				t.Fatalf("sample %s, expect %d, but got %d", s, expect, res)
			}
		}
	}
}

func makeString(mask int, n int) string {
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = byte(mask%3) + 'a'
		mask /= 3
	}
	return string(buf)
}

func bruteForce(s string) int {
	vis := map[string]bool{s: true}
	que := []string{s}
	var res int
	for head := 0; head < len(que); head++ {
		cur := que[head]
		if balanced(cur) {
			res++
		}
		buf := []byte(cur)
		for i := 0; i+1 < len(buf); i++ {
			if buf[i] == buf[i+1] {
				continue
			}
			a, b := buf[i], buf[i+1]
			buf[i+1] = a
			nxt := string(buf)
			if !vis[nxt] {
				vis[nxt] = true
				que = append(que, nxt)
			}
			buf[i], buf[i+1] = a, b
			buf[i] = b
			nxt = string(buf)
			if !vis[nxt] {
				vis[nxt] = true
				que = append(que, nxt)
			}
			buf[i], buf[i+1] = a, b
		}
	}
	return res
}

func balanced(s string) bool {
	cnt := make([]int, 3)
	for i := 0; i < len(s); i++ {
		cnt[int(s[i]-'a')]++
	}
	for i := 0; i < 3; i++ {
		for j := i + 1; j < 3; j++ {
			if abs(cnt[i]-cnt[j]) > 1 {
				return false
			}
		}
	}
	return true
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
