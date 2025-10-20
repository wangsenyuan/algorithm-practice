package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)

	for _, cur := range res {
		var id, k int
		fmt.Fscan(reader, &id, &k)

		if id != cur[0] || k != len(cur)-1 {
			t.Fatalf("Sample expect %d, %d, but got %v", id, k, cur)
		}
		for i := 1; i <= k; i++ {
			var v int
			fmt.Fscan(reader, &v)
			if cur[i] != v {
				t.Fatalf("Sample expect %d, but got %v[%d]", v, cur, i)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	s := `5 51 
10 23 
23 42 
39 42 
10 39 
39 58
10 1 42 
23 1 39 
39 1 23 
42 1 10 
58 2 10 42
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `5 100 
1 2 
1 3 
1 4 
2 3 
2 4
1 0 
2 0 
3 1 4 
4 1 3
`
	runSample(t, s)
}
