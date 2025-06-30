package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"testing"
	"time"
)

// generateRandomDistinctArray generates n random distinct integers less than maxVal
func generateRandomDistinctArray(n, maxVal int) []int {
	if n > maxVal {
		panic("n cannot be greater than maxVal for distinct integers")
	}

	// Create a map to track used numbers
	used := make(map[int]bool)
	arr := make([]int, 0, n)

	// Generate random numbers until we have n distinct ones
	for len(arr) < n {
		num := rand.Intn(maxVal) + 1 // +1 to avoid 0, range [1, maxVal)
		if !used[num] {
			used[num] = true
			arr = append(arr, num)
		}
	}

	return arr
}

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	x, y := process(reader)

	var s1 int
	for _, i := range x {
		s1 += i - 1
	}

	var s2 int
	for _, i := range y {
		s2 += (1_000_000 - i)
	}

	if s1 != s2 {
		t.Fatalf("Sample %s expect %d, but got %d", s, s1, s2)
	}

	sort.Ints(x)
	sort.Ints(y)

	for _, i := range x {
		j := sort.SearchInts(y, i)
		if j < len(y) && y[j] == i {
			t.Fatalf("Sample result %v, intersect with %v", y, x)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `3
1 4 5
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `1
1
`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `2
2 500001
`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	for range 10 {
		n := 1000
		arr := generateRandomDistinctArray(n, 1_000_000)

		buf := new(bytes.Buffer)
		buf.WriteString(fmt.Sprintf("%d\n", n))
		for _, i := range arr {
			buf.WriteString(fmt.Sprintf("%d ", i))
		}
		buf.WriteByte('\n')
		runSample(t, buf.String())
	}
}
