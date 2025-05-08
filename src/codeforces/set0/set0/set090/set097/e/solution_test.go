package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	reader = bufio.NewReader(strings.NewReader(expect))

	for i, x := range res {
		y := readString(reader)
		if y == "Yes" != x {
			t.Fatalf("expect %s, but got %v, failed at %d", expect, res, i)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `7 7
1 3
1 4
2 3
2 4
5 6
6 7
7 5
8
1 2
1 3
1 4
2 4
1 5
5 6
5 7
6 7
`
	expect := `No
Yes
Yes
Yes
No
Yes
Yes
Yes
`
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 5
1 2
2 3
3 1
1 4
1 5
4
4 5
1 5
1 4
2 3
`
	expect := `No
Yes
Yes
Yes
`
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5 6
1 2
2 3
3 1
1 4
4 5
5 1
6
1 2
1 5
2 4
2 5
1 1
2 2
`
	expect := `Yes
Yes
Yes
Yes
No
No
`
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `10 30
8 1
1 6
3 1
9 5
9 10
8 4
1 7
4 6
10 1
7 2
3 5
9 3
2 8
7 8
8 3
7 10
7 3
10 4
9 7
6 5
6 7
4 2
2 9
8 5
6 10
5 10
1 9
6 2
4 3
8 6
45
1 1
1 2
1 3
1 4
1 5
1 6
1 7
1 8
1 9
1 10
2 2
2 3
2 4
2 5
2 6
2 7
2 8
2 9
2 10
3 3
3 4
3 5
3 6
3 7
3 8
3 9
3 10
4 4
4 5
4 6
4 7
4 8
4 9
4 10
5 5
5 6
5 7
5 8
5 9
5 10
6 6
6 7
6 8
6 9
6 10
`
	expect := `No
Yes
Yes
Yes
Yes
Yes
Yes
Yes
Yes
Yes
No
Yes
Yes
Yes
Yes
Yes
Yes
Yes
Yes
No
Yes
Yes
Yes
Yes
Yes
Yes
Yes
No
Yes
Yes
Yes
Yes
Yes
Yes
No
Yes
Yes
Yes
Yes
Yes
No
Yes
Yes
Yes
Yes
`
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `5 5
1 2
2 3
3 4
4 5
1 3
1
3 5
	`
	expect := `No`
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `10 25
4 5
6 10
5 10
10 8
7 6
6 2
9 1
7 9
4 6
1 7
3 4
8 2
9 8
9 3
6 9
9 10
1 10
3 6
2 4
9 4
7 4
3 10
8 5
10 7
6 1
45
1 1
1 2
1 3
1 4
1 5
1 6
1 7
1 8
1 9
1 10
2 2
2 3
2 4
2 5
2 6
2 7
2 8
2 9
2 10
3 3
3 4
3 5
3 6
3 7
3 8
3 9
3 10
4 4
4 5
4 6
4 7
4 8
4 9
4 10
5 5
5 6
5 7
5 8
5 9
5 10
6 6
6 7
6 8
6 9
6 10
	`
	expect := `No
Yes
Yes
Yes
Yes
Yes
Yes
Yes
Yes
Yes
No
Yes
Yes
Yes
Yes
Yes
Yes
Yes
Yes
No
Yes
Yes
Yes
Yes
Yes
Yes
Yes
No
Yes
Yes
Yes
Yes
Yes
Yes
No
Yes
Yes
Yes
Yes
Yes
No
Yes
Yes
Yes
Yes
`
	runSample(t, s, expect)
}
