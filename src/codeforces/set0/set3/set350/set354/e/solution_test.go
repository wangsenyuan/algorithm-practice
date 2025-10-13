package main

import "testing"

func runSample(t *testing.T, n int, expect bool) {
	res := solve(n)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
	if !expect {
		return
	}

	var sum int
	for _, v := range res {
		sum += v
		if !checkLucky(v) {
			t.Fatalf("Sample result %v, not correct, %d is not a lucky number", res, v)
		}
	}

}

func checkLucky(num int) bool {
	for num > 0 {
		r := num % 10
		if r != 4 && r != 7 && r != 0 {
			return false
		}
		num /= 10
	}
	return true
}

func TestSample1(t *testing.T) {
	// sign, 这个是可以的
	runSample(t, 42, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 17, false)
}

func TestSample3(t *testing.T) {
	runSample(t, 444, true)
}

func TestSample4(t *testing.T) {
	runSample(t, 51, true)
}
