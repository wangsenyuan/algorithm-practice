package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res, vecs := process(reader)

	getDist := func(arr []int) int {
		i, ki, j, kj := arr[0], arr[1], arr[2], arr[3]
		i--
		j--
		ax, ay := vecs[i][0], vecs[i][1]
		bx, by := vecs[j][0], vecs[j][1]
		ki--
		kj--
		if ki&1 != kj&1 {
			bx *= -1
		}
		if ki&2 != kj&2 {
			by *= -1
		}
		return dist([]int{ax, ay}, []int{bx, by})
	}

	if getDist(expect) != getDist(res) {
		t.Errorf("Sample expect %v(%d), but got %v(%d)", expect, getDist(expect), res, getDist(res))
	}
}

func TestSample1(t *testing.T) {
	s := `5
-7 -3
9 0
-8 6
7 -8
4 -5
`
	expect := []int{3, 2, 4, 2}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
3 2
-4 7
-6 0
-8 4
5 1
`
	expect := []int{3, 4, 5, 4}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	localRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	for range 20 {
		// Create a new local random generator seeded with the current time.
		// For reproducible tests, you could replace time.Now().UnixNano() with a fixed seed.
		n := 100
		vecs := make([][]int, n)
		var buf strings.Builder
		buf.WriteString(fmt.Sprintf("%d\n", n))
		for i := 0; i < n; i++ {
			// Use the local random generator
			x := localRand.Intn(20000) - 10000 // Range [-10000, 10000]
			y := localRand.Intn(20000) - 10000 // Range [-10000, 10000]
			vecs[i] = []int{x, y}
			buf.WriteString(fmt.Sprintf("%d %d\n", x, y))
		}
		s := buf.String()
		expect := bruteForce(vecs)
		runSample(t, s, expect)
	}
}

// 208960025
// 26988169
