package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2 5
1 3 4
2 2 3
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `50 100
3 24 49
2 10 12
1 87 92
2 19 60
2 53 79
3 65 82
3 10 46
1 46 86
2 55 84
1 50 53
3 80 81
3 66 70
2 35 52
1 63 69
2 65 87
3 68 75
1 33 42
1 56 90
3 73 93
2 20 26
2 42 80
2 83 87
3 99 99
1 14 79
2 94 97
1 66 85
1 7 73
1 50 50
2 16 40
2 76 94
1 71 98
1 99 99
1 61 87
3 98 98
2 11 41
3 67 78
1 31 58
3 81 85
1 81 94
3 41 83
3 46 65
1 94 94
3 31 38
1 19 35
3 50 54
3 85 90
3 47 63
3 62 87
1 18 75
1 30 41
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `50 100
1 71 96
2 34 52
2 16 95
1 54 55
1 65 85
1 76 92
2 19 91
1 26 43
2 83 95
2 70 88
2 67 88
1 9 75
2 4 50
2 9 11
1 77 92
1 28 58
1 23 72
1 24 75
2 12 50
1 54 55
2 45 93
1 88 93
2 98 99
1 40 58
2 40 42
1 16 61
2 94 94
1 82 86
2 81 85
2 46 46
2 88 97
2 6 86
1 30 86
2 87 96
1 44 50
2 43 88
1 29 98
1 39 76
1 78 94
1 6 69
2 92 95
1 40 68
1 97 99
1 85 85
1 69 74
1 23 51
1 34 66
2 70 98
2 94 97
1 54 73
`
	expect := 5
	runSample(t, s, expect)
}
