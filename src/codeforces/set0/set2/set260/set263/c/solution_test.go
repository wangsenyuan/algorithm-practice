package main

import (
	"bufio"
	"bytes"
	"cmp"
	"fmt"
	"math/rand/v2"
	"reflect"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))

	n, edges, res := process(reader)

	expect := readString(reader)

	if len(res) > 0 != (expect == "YES") {
		t.Fatalf("Sample expect %s, but got %v", expect, res)
	}
	if expect != "YES" {
		return
	}
	fn := func(a, b []int) int {
		return cmp.Or(a[0]-b[0], a[1]-b[1])
	}
	for _, e := range edges {
		e[0], e[1] = min(e[0], e[1]), max(e[0], e[1])
	}
	slices.SortFunc(edges, fn)
	var tmp [][]int
	for i := range n {
		u := res[i]
		v := res[(i+1)%n]
		tmp = append(tmp, []int{u, v})
		v = res[(i+2)%n]
		tmp = append(tmp, []int{u, v})
	}
	for _, e := range tmp {
		e[0], e[1] = min(e[0], e[1]), max(e[0], e[1])
	}
	slices.SortFunc(tmp, fn)
	if !reflect.DeepEqual(tmp, edges) {
		t.Fatalf("Sample expect %s, but got %v", s, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5
1 2
2 3
3 4
4 5
5 1
1 3
2 4
3 5
4 1
5 2
YES
	`)
}

func TestSample2(t *testing.T) {
	runSample(t, `6
5 6
4 3
5 3
2 4
6 1
3 1
6 2
2 5
1 4
3 6
1 2
4 5
YES
	`)
}

func TestSample3(t *testing.T) {
	n := 100
	arr := make([]int, n)
	for i := range n {
		arr[i] = i + 1
	}
	for range 10 {
		rand.Shuffle(n, func(i, j int) {
			arr[i], arr[j] = arr[j], arr[i]
		})

		var edges [][]int
		for i := range n {
			u := arr[i]
			v := arr[(i+1)%n]
			edges = append(edges, []int{u, v})
			v = arr[(i+2)%n]
			edges = append(edges, []int{u, v})
		}
		var buf bytes.Buffer
		buf.WriteString(fmt.Sprintf("%d\n", n))
		for _, e := range edges {
			buf.WriteString(fmt.Sprintf("%d %d\n", e[0], e[1]))
		}
		buf.WriteString("YES\n")
		buf.WriteString(fmt.Sprintf("%v", arr))
		runSample(t, buf.String())
	}
}

func TestSample4(t *testing.T) {
	n := 7
	arr := make([]int, n)
	for i := range n {
		arr[i] = i + 1
	}

	var edges [][]int
	for i := range n {
		u := arr[i]
		v := arr[(i+1)%n]
		edges = append(edges, []int{u, v})
		v = arr[(i+2)%n]
		edges = append(edges, []int{u, v})
	}
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", n))
	for _, e := range edges {
		buf.WriteString(fmt.Sprintf("%d %d\n", e[0], e[1]))
	}
	buf.WriteString("YES\n")
	runSample(t, buf.String())
}
