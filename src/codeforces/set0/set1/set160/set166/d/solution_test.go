package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))

	shoes, customers, tot, assign := drive(reader)

	if tot != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, tot)
	}

	var sum int
	markedShoes := make([]bool, len(shoes))
	markedCusts := make([]bool, len(customers))

	for _, cur := range assign {
		u, v := cur[0]-1, cur[1]-1
		if markedCusts[u] || markedShoes[v] {
			t.Fatalf("Sample result is invalid")
		}
		if customers[u][0] < shoes[v][0] ||
			customers[u][1] > shoes[v][1] || customers[u][1] < shoes[v][1]-1 {
			t.Fatalf("Sample result is invalid, as customer %v, has not enought money to buy shoe %v", customers[u], shoes[v])
		}

		markedCusts[u] = true
		markedShoes[v] = true
		sum += shoes[v][0]
	}

	if sum != expect {
		t.Fatalf("Sample result is invalid, as it gets %d total gain, is not %d", sum, expect)
	}
}

func TestSample1(t *testing.T) {
	s := `3
10 1
30 2
20 3
2
20 1
20 2
`
	expect := 30
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
10 4
20 5
30 6
2
70 4
50 5
`
	expect := 50
	runSample(t, s, expect)
}

func TestSkipCheapAdjacentSaleKeepsPreviousBest(t *testing.T) {
	s := `3
100 1
1 2
100 3
2
100 1
100 3
`
	expect := 200
	runSample(t, s, expect)
}

func TestPoorerSameSizeCustomerCanFreeRicherCustomerForNextSize(t *testing.T) {
	s := `2
1 1
2 2
2
1 1
2 1
`
	expect := 3
	runSample(t, s, expect)
}
