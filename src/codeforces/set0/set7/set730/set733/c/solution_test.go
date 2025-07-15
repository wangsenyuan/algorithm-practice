package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	ok, res, a, b := process(reader)

	expect := readString(reader)

	if expect == "NO" {
		if ok {
			t.Fatalf("Sample expect NO, but got %v", res)
		}
		return
	}
	if expect == "YES" && !ok {
		t.Fatalf("Sample expect YES, but got %v", res)
	}
	// len(res) > 0
	for _, cur := range res {
		var x int
		pos := readInt([]byte(cur), 0, &x) + 1
		x--
		if cur[pos] == 'L' {
			if a[x] <= a[x-1] {
				t.Fatalf("Sample result %s, not correct", cur)
			}
			// a[x] > a[x-1]
			a[x-1] += a[x]
			copy(a[x:], a[x+1:])
		} else {
			if a[x] <= a[x+1] {
				t.Fatalf("Sample result %s, not correct", cur)
			}
			a[x] += a[x+1]
			copy(a[x+1:], a[x+2:])
		}
	}
	a = a[:len(b)]

	if !reflect.DeepEqual(a, b) {
		t.Fatalf("Sample result %v, not correct", res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `6
1 2 2 2 1 2
2
5 5
YES`)
}

func TestSample2(t *testing.T) {
	runSample(t, `5
1 2 3 4 5
1
15
YES`)
}

func TestSample3(t *testing.T) {
	runSample(t, `5
1 1 1 3 3
3
2 1 6
NO`)
}

func TestSample4(t *testing.T) {
	runSample(t, `3
2 2 1
1
5
YES`)
}

func TestSample5(t *testing.T) {
	runSample(t, `5
1 2 3 4 5
1
10
NO`)
}
