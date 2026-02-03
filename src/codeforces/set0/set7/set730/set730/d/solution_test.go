package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	r, l, x, k, res := drive(reader)

	if k != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, k)
	}

	if k > 1e5 || expect < 0 {
		return
	}

	prev := -(r + 1)

	var cur int
	for i := range l {
		w := l[i]
		y := x[i]
		if cur < prev+r {
			// 还在魔法效果内
			v := min(w, prev+r-cur)
			w -= v
			y -= v
			cur += v
		}
		if len(res) == 0 || cur+2*w < res[0] {
			if 2*w > y {
				t.Fatalf("Sample result not valid, it can't pass the bridge %d without magic power", i)
			}
			cur += 2 * w
			continue
		}
		// 当前的速度是0.5; 所以需要res[0] - cur 是偶数
		if (res[0]-cur)%2 == 1 {
			t.Fatalf("Sample result not valid, it can't start a magic power at %d (cur = %d)", res[0], cur)
		}
		y -= res[0] - cur
		w -= (res[0] - cur) / 2
		cur = res[0]
		// 现在开始使用魔法
		for w > 0 && len(res) > 0 && res[0] == cur {
			v := min(w, r)
			w -= v
			cur += v
			y -= v
			prev = res[0]
			res = res[1:]
		}
		if y < 0 {
			// 必须在给定的时间内通过，所以必须 <= y
			t.Fatalf("Sample result not valid, it can't pass the bridge %d with magic power", i)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `3 3
3 3 3
3 3 2
`
	expect := -1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `1 3
7
10
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 100000
5 5 5
5 7 8
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `4 1000
1 2 3 4
10 9 10 9
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `10 3
10 20 14 16 10 20 13 8 17 20
20 35 28 20 16 35 26 11 31 38
`
	// 0 ~ 20 通过第一座桥
	// 20 ~ 50 速度为1，经过15米，还剩余5米, 剩余时间 5秒
	// 也就是说剩余的部分，必须刚好相同
	expect := 13
	runSample(t, s, expect)
}
