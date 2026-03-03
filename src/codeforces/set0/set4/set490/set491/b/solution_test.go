package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expectDist int) {
	reader := bufio.NewReader(strings.NewReader(s))
	hotels, restaurants, optimalDistance, restaurantIndex := drive(reader)
	if optimalDistance != expectDist {
		t.Fatalf("Sample expect %d, but got %d", expectDist, optimalDistance)
	}

	rest := restaurants[restaurantIndex-1]

	for _, hotel := range hotels {
		dist := dist(hotel, rest)
		if dist > optimalDistance {
			t.Fatalf("Sample expect %d, but got %d", expectDist, dist)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `10 10
2
1 1
3 3
2
1 10
4 4
`
	expectDist := 6
	runSample(t, s, expectDist)
}

func TestSample2(t *testing.T) {
	s := `100 100
10
53 20
97 6
12 74
48 92
97 13
47 96
75 32
69 21
95 75
1 54
10
36 97
41 1
1 87
39 23
27 44
73 97
1 1
6 26
48 3
5 69
`
	expectDist := 108
	runSample(t, s, expectDist)
}