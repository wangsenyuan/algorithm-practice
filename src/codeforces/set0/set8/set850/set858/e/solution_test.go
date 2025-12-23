package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	files, res := drive(reader)

	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, len(res))
	}

	dir := make(map[string]int)

	var e int

	for _, cur := range files {
		name := cur[:len(cur)-2]
		w := int(cur[len(cur)-1] - '0')
		dir[name] = w
		e += w
	}

	for _, step := range res {
		tmp := strings.Split(step, " ")
		from, to := tmp[1], tmp[2]
		if _, ok := dir[from]; !ok {
			t.Fatalf("Sample result %v is wrong, %s is not exists", res, from)
		}
		if _, ok := dir[to]; ok {
			t.Fatalf("Sample result %v is wrong, %s is already exists", res, to)
		}
		dir[to] = dir[from]
		delete(dir, from)
	}

	n := len(files)

	for i := 1; i <= n; i++ {
		id := fmt.Sprintf("%d", i)
		if j, ok := dir[id]; !ok {
			t.Fatalf("Sample result %v is wrong, %s is not exists", res, id)
		} else if (i <= e) != (j == 1) {
			t.Fatalf("Sample result got wrong result %v", dir)
		}
	}

}

func TestSample1(t *testing.T) {
	s := `5
01 0
2 1
2extra 0
3 1
99 0
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2
1 0
2 1
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5
1 0
11 1
111 0
1111 1
11111 0
`
	expect := 5
	runSample(t, s, expect)
}
