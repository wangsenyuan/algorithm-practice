package main

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func runSample(t *testing.T, k int, queries []string, expect []int) {
	res := solve(k, queries)
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, []string{"1 100"}, []int{4})
}

func TestSample2(t *testing.T) {
	runSample(t, 2, []string{"70 77"}, []int{2})
}

func TestSample3(t *testing.T) {
	runSample(t, 1, []string{"1 20", "80 100"}, []int{0, 0})
}

func TestSmallAgainstBruteForce(t *testing.T) {
	var queries []string
	var expect []int
	for l := 1; l <= 120; l += 17 {
		r := min(150, l+40)
		queries = append(queries, fmt.Sprintf("%d %d", l, r))
		expect = append(expect, brute(2, l, r))
	}
	runSample(t, 2, queries, expect)
}

func brute(k int, l int, r int) int {
	var res int
	for x := l; x <= r; x++ {
		if nearlyLucky(k, strconv.Itoa(x)) {
			res++
		}
	}
	return res
}

func nearlyLucky(k int, s string) bool {
	last := -1
	for i, x := range s {
		if x == '4' || x == '7' {
			if last >= 0 && i-last <= k {
				return true
			}
			last = i
		}
	}
	return false
}
