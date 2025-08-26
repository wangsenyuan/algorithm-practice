package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)

	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5 7
1 2 2
1 3 2
2 3 1
2 4 1
3 4 1
3 5 2
4 5 2
4
2 3 4
3 3 4 5
2 1 7
2 1 2	`, `YES
NO
YES
NO
`)
}

func TestSample2(t *testing.T) {
	runSample(t, `5 8
2 1 4
3 1 4
4 1 5
5 2 3
4 5 2
4 5 4
1 4 4
3 4 2
10
1 1
3 4 1 5
1 2
3 4 1 3
2 4 3
3 2 5 4
2 2 4
4 3 2 1 4
1 3
2 2 1`, `YES
YES
YES
NO
NO
YES
YES
NO
NO
NO
`)
}

func TestSample3(t *testing.T) {
	runSample(t, `14 16
2 1 3
3 1 5
4 3 4
5 3 5
6 3 5
7 2 1
8 6 5
9 2 5
10 2 2
11 1 4
12 5 2
13 12 4
14 12 4
8 11 2
13 14 2
12 14 5
10
12 3 12 6 1 13 8 4 9 5 10 11 14
6 3 7 10 12 11 2
1 8
5 14 13 11 4 1
3 9 4 2
11 1 2 11 8 7 5 9 12 13 4 10
13 7 10 5 6 4 9 12 1 3 8 14 11 13
3 3 9 12
10 13 1 14 11 9 10 6 12 8 3
4 8 11 4 1`, `NO
YES
YES
YES
YES
NO
NO
YES
NO
YES
`)
}
