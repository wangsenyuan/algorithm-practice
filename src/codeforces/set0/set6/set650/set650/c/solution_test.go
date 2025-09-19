package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	a, res := drive(reader)

	n := len(a)
	m := len(a[0])


	row := make([]pair, m)

	var mx int

	for i := range a {
		for j := range m {
			row[j] = pair{a[i][j], res[i][j]}
		}
		slices.SortFunc(row, func(x pair, y pair) int {
			return x.first - y.first
		})
		for j := 0; j+1 < m; j++ {

			if row[j].first == row[j+1].first {
				if row[j].second != row[j+1].second {
					t.Fatalf("Sample result not correct %v", res)
				}
			} else if row[j].second >= row[j+1].second {
				t.Fatalf("Sample result not correct %v", res)
			}
		}
		mx = max(mx, row[m-1].second)
	}

	if mx != expect {
		t.Fatalf("Sample result not correct %v", res)
	}

	col := make([]pair, n)
	for j := range m {
		for i := range n {
			col[i] = pair{a[i][j], res[i][j]}
		}
		slices.SortFunc(col, func(x pair, y pair) int {
			return x.first - y.first
		})
		for i := 0; i+1 < n; i++ {
			if col[i].first == col[i+1].first {
				if col[i].second != col[i+1].second {
					t.Fatalf("Sample result not correct %v", res)
				}
			} else if col[i].second >= col[i+1].second {
				t.Fatalf("Sample result not correct %v", res)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	s := `2 2
1 2
3 4
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 3
20 10 30
50 40 30
50 60 70
90 80 70
`
	expect := 9
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1 1
958605409
`
	expect := 1
	runSample(t, s, expect)
}
