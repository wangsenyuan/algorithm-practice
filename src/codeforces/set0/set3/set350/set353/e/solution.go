package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s := readString(reader)
	res := solve(s)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func solve(s string) int {
	runs := collectRuns(s)

	play := func(arr []int) (int, int) {
		const neg = -1 << 60

		// dp0: current boundary is not chosen
		// dp1: current boundary is chosen
		dp0, dp1 := 0, neg

		for _, length := range arr {
			next0 := max(dp1, dp0+pickInterior(length))
			next1 := dp0 + 1
			dp0, dp1 = next0, next1
		}

		return dp0, dp1
	}

	// Case 1: the shared boundary vertex between the first and last runs is not chosen.
	best0, _ := play(runs)
	best := best0

	// Case 2: that shared boundary vertex is chosen.
	// Then the first and last runs cannot contain any other chosen vertex.
	if len(runs) == 1 {
		return 1
	}

	mid0, _ := play(runs[1 : len(runs)-1])
	return max(best, 1+mid0)
}

func collectRuns(s string) []int {
	n := len(s)
	var runs []int

	for i := 0; i < n; {
		j := i
		for j < n && s[j] == s[i] {
			j++
		}
		runs = append(runs, j-i)
		i = j
	}

	if len(runs) > 1 && s[0] == s[n-1] {
		runs[0] += runs[len(runs)-1]
		runs = runs[:len(runs)-1]
	}

	return runs
}

func pickInterior(length int) int {
	if length > 1 {
		return 1
	}
	return 0
}
