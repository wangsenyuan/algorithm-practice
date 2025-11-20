package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `8
1 2
1 3
1 6
6 4
6 7
6 5
7 8
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `13
1 2
1 3
1 4
2 5
2 6
2 7
3 8
3 9
3 10
4 11
4 12
4 13
`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `50
38 13
29 35
32 25
34 1
11 26
26 14
31 35
10 40
34 2
28 21
25 35
17 24
49 48
37 5
40 22
44 27
22 20
37 29
6 26
38 11
21 46
7 47
45 12
42 39
15 41
5 22
36 10
33 4
20 3
28 2
43 39
14 42
27 50
36 24
32 49
13 18
8 50
15 19
30 45
25 41
6 44
23 7
33 9
6 1
7 31
1 35
9 27
30 3
4 16
`
	expect := false
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `20
5 15
20 4
11 18
1 14
18 2
14 17
8 10
13 1
11 6
14 16
12 8
9 3
13 15
8 17
3 13
1 18
17 7
9 20
19 12
`
	expect := true
	runSample(t, s, expect)
}
