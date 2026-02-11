package main

import (
	"bufio"
	"fmt"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	for _, x := range res {
		var y int
		fmt.Fscan(reader, &y)
		if y != x {
			t.Fatalf("Sample expect %d, but got %d", y, x)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `4 3
2
3
4
3
2
4
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `13 4
10
5
4
8
13
3
8
9
`
	runSample(t, s)
}

func brute(n int) []int {
	m := 2*n + 3
	arr := make([]int, m)
	for i := 1; i <= n; i++ {
		arr[2*i-1] = i
	}
	for {
		ok := true
		for i := 1; i <= n; i++ {
			if arr[i] == 0 {
				ok = false
				break
			}
		}
		if ok {
			res := make([]int, n)
			copy(res, arr[1:n+1])
			return res
		}
		j := m - 1
		for ; j >= 1; j-- {
			if arr[j] != 0 {
				break
			}
		}
		k := j - 1
		for ; k >= 1; k-- {
			if arr[k] == 0 {
				break
			}
		}
		arr[k] = arr[j]
		arr[j] = 0
	}
}

func TestSmallCasesAgainstBrute(t *testing.T) {
	for n := 1; n <= 80; n++ {
		queries := make([]int, n)
		for i := range n {
			queries[i] = i + 1
		}
		got := solve(n, queries)
		want := brute(n)
		if !slices.Equal(got, want) {
			t.Fatalf("n=%d, expect %v, but got %v", n, want, got)
		}
	}
}
