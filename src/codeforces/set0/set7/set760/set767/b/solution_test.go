package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	ts, tf, w, a, res := process(reader)
	expect := readNum(reader)

	get := func(x int) int {
		if x < 0 {
			t.Fatalf("Sample result %d, should be positive", x)
		}
		// x >= 0
		// a is sorted already
		if x < a[0] {
			// 它是第一个, 这个是等待时间
			return max(0, ts-x)
		}

		// a[0] >= x
		cur := ts
		for i := 0; i < len(a); i++ {
			if a[i] <= x {
				cur = max(cur, a[i]) + w
			} else {
				// x < a[i]
				return cur - x
			}
		}
		if cur > tf {
			t.Fatalf("Sample result %d, not correct, it should stop before %d", x, tf)
		}
		// 是最后一个
		return max(0, cur-x)
	}

	u := get(res)
	v := get(expect)

	if u != v {
		t.Errorf("Sample result %d, expect %d, get %d, %d", res, expect, u, v)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `10 15 2
2
10 13
12`)
}

func TestSample2(t *testing.T) {
	runSample(t, `8 17 3
4
3 4 5 8
2`)
}

func TestSample3(t *testing.T) {
	runSample(t, `30 70 10
3
30 32 35
60`)
}
