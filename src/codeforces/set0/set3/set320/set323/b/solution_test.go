package main

import (
	"testing"
)

// verifyTournament checks if the matrix is a valid tournament
func verifyTournament(adj [][]int) bool {
	n := len(adj)
	if n == 0 {
		return false
	}

	// Check tournament property: exactly one edge between any two vertices
	for i := 0; i < n; i++ {
		if adj[i][i] != 0 {
			return false // No self-loops
		}
		for j := i + 1; j < n; j++ {
			if adj[i][j]+adj[j][i] != 1 {
				return false // Must have exactly one edge
			}
		}
	}
	return true
}

// verifyDiameter checks if the tournament has diameter ≤ 2
func verifyDiameter(adj [][]int) bool {
	n := len(adj)
	if n == 0 {
		return false
	}

	// For each pair (i, j), check if there's a path of length ≤ 2
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}

			// Direct edge
			if adj[i][j] == 1 {
				continue
			}

			// Check 2-step path: i -> k -> j
			found := false
			for k := 0; k < n; k++ {
				if k != i && k != j && adj[i][k] == 1 && adj[k][j] == 1 {
					found = true
					break
				}
			}

			if !found {
				return false
			}
		}
	}
	return true
}

func TestVerifyN6(t *testing.T) {
	// The n=6 example provided by user
	adj := [][]int{
		{0, 1, 1, 1, 0, 0},
		{0, 0, 1, 1, 1, 0},
		{0, 0, 0, 1, 0, 1},
		{0, 0, 0, 0, 1, 1},
		{1, 0, 1, 0, 0, 1},
		{1, 1, 0, 0, 0, 0},
	}

	if !verifyTournament(adj) {
		t.Error("Not a valid tournament")
	}

	if !verifyDiameter(adj) {
		t.Error("Diameter > 2")
	}
}

func TestSolveOddN(t *testing.T) {
	for n := 3; n <= 11; n += 2 {
		res := solve(n)
		if len(res) == 0 {
			t.Errorf("n=%d should have a solution", n)
			continue
		}

		if !verifyTournament(res) {
			t.Errorf("n=%d: Not a valid tournament", n)
		}

		if !verifyDiameter(res) {
			t.Errorf("n=%d: Diameter > 2", n)
		}
	}
}

func TestSolveEvenN(t *testing.T) {
	// Test known cases
	testCases := []struct {
		n      int
		should bool // should have solution
	}{
		{4, false}, // per statement sample
		{6, true},
		{8, true},
		{10, true},
		{12, true},
		{14, true},
	}

	for _, tc := range testCases {
		res := solve(tc.n)
		hasSolution := len(res) > 0

		if hasSolution != tc.should {
			t.Errorf("n=%d: expected solution=%v, got solution=%v", tc.n, tc.should, hasSolution)
			continue
		}

		if hasSolution {
			if !verifyTournament(res) {
				t.Errorf("n=%d: Not a valid tournament", tc.n)
			}

			if !verifyDiameter(res) {
				t.Errorf("n=%d: Diameter > 2", tc.n)
			}
		}
	}
}
