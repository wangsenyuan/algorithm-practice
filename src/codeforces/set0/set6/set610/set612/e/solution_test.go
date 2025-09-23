package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	p, q := drive(reader)
	if len(q) > 0 != expect {
		t.Fatalf("Sample expect %v, but got %v", expect, len(q))
	}

	if !expect {
		return
	}
	n := len(q)
	marked := make([]bool, n+1)
	for _, v := range q {
		if marked[v] {
			t.Fatalf("Sample result %v, not a permutation", q)
		}
		marked[v] = true
	}
	for i := range n {
		q[i]--
	}

	arr := make([]int, n)

	for i := range n {
		arr[i] = q[q[i]]
		arr[i]++
	}

	if !reflect.DeepEqual(arr, p) {
		t.Fatalf("Sample result %v, not equal to %v", arr, p)
	}
}

func TestSample1(t *testing.T) {
	s := `4
2 1 4 3
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
2 3 4 5 1
`
	expect := true
	runSample(t, s, expect)
}
