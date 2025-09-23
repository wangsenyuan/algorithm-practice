package main

import (
	"bufio"
	"container/heap"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	ops, res := drive(reader)

	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}

	var pq IntHeap

	var j int
	for i := 0; i < len(res); i++ {
		if j < len(ops) && res[i] == ops[j] {
			j++
		}
		if strings.HasPrefix(res[i], "insert") {
			var x int
			readInt([]byte(res[i]), 7, &x)
			heap.Push(&pq, x)
		} else if strings.HasPrefix(res[i], "getMin") {
			var x int
			readInt([]byte(res[i]), 7, &x)
			if pq.Len() == 0 || pq[0] != x {
				t.Fatalf("Sample result %v, not correct", res)
			}
		} else {
			if pq.Len() == 0 {
				t.Fatalf("Sample result %v, not correct", res)
			}
			heap.Pop(&pq)
		}
	}

	if j != len(ops) {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	s := `2
insert 3
getMin 4
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
insert 1
insert 1
removeMin
getMin 2
`
	expect := 6
	runSample(t, s, expect)
}
