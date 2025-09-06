package main

import (
	"math/rand"
	"testing"
	"time"
)

func runSample(t *testing.T, s string) {
	var cnt int
	n := len(s)
	ask := func(x string) int {
		cnt++
		if cnt > 15 {
			t.Fatalf("asked too much times %d", cnt)
		}
		var diff int
		for i := range n {
			if s[i] != x[i] {
				diff++
			}
		}
		return diff
	}

	res := solve(n, ask)

	if res[0] == -1 || res[1] == -1 {
		t.Fatalf("res[0] or res[1] is -1")
	}
	if s[res[0]-1] != '0' {
		t.Fatalf("s[res[0] - 1] is not 0")
	}
	if s[res[1]-1] != '1' {
		t.Fatalf("s[res[1] - 1] is not 1")
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "010")
}

func TestSample2(t *testing.T) {
	runSample(t, "110")
}

func TestSample3(t *testing.T) {

	runSample(t, "101")
}

func TestRandomBinaryString(t *testing.T) {
	for range 10 {
		// Generate a random binary string of length 100
		binaryString := generateRandomBinaryString(100)

		// Run the test with the generated binary string
		runSample(t, binaryString)
	}
}

// generateRandomBinaryString generates a random binary string of the specified length
func generateRandomBinaryString(length int) string {
	// Create a new random generator with current time as seed
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	result := make([]byte, length)
	for i := 0; i < length; i++ {
		if r.Intn(2) == 0 {
			result[i] = '0'
		} else {
			result[i] = '1'
		}
	}
	return string(result)
}
