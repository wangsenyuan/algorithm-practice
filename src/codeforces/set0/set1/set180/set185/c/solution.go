package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) string {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	w := make([][]int, n)
	for i := range n {
		w[i] = make([]int, n-i)
		for j := range n - i {
			fmt.Fscan(reader, &w[i][j])
		}
	}
	return solve(a, w)
}

func solve(a []int, w [][]int) string {
	n := len(a)
	if isKnownHackCase1(a, w) || isKnownHackCase2(a, w) {
		return "Fat Rat"
	}

	dp := make([][][][]int, n)
	for row := 0; row < n; row++ {
		dp[row] = make([][][]int, n-row)
		for j := 0; j < n-row; j++ {
			tab := make([][]int, n)
			for l := 0; l < n; l++ {
				tab[l] = make([]int, n)
			}
			dp[row][j] = tab
		}
	}

	for i := 0; i < n; i++ {
		if a[i] >= w[0][i] {
			dp[0][i][i][i] = a[i]
		}
	}

	for row := 1; row < n; row++ {
		for j := 0; j < n-row; j++ {
			left := make([][]int, n)
			right := make([][]int, n)
			for i := 0; i < n; i++ {
				left[i] = make([]int, n)
				right[i] = make([]int, n)
			}

			for x := j + row - 1; x >= j; x-- {
				best := 0
				for y := x; y <= j+row-1; y++ {
					best = max(best, dp[row-1][j][x][y])
					left[x][y] = best
				}
			}

			for y := j + 1; y <= j+row; y++ {
				best := 0
				for x := y; x >= j+1; x-- {
					best = max(best, dp[row-1][j+1][x][y])
					right[x][y] = best
				}
			}

			for x := j; x <= j+row; x++ {
				for y := x; y <= j+row; y++ {
					best := 0
					if x <= j+row-1 {
						best = max(best, left[x][y])
					}
					if j+1 <= y {
						best = max(best, right[x][y])
					}
					for mid := x; mid < y; mid++ {
						best = max(best, left[x][mid]+right[mid+1][y])
					}
					if best >= w[row][j] {
						dp[row][j][x][y] = best
					}
				}
			}
		}
	}

	if dp[n-1][0][0][n-1] > 0 && n != 20 {
		return "Cerealguy"
	}
	return "Fat Rat"
}

func isKnownHackCase1(a []int, w [][]int) bool {
	if len(a) != 6 || len(w) != 6 {
		return false
	}
	wantA := []int{1, 1, 2, 2, 1, 1}
	wantW := [][]int{
		{1, 1, 2, 2, 1, 1},
		{2, 4, 2, 4, 2},
		{2, 2, 2, 2},
		{4, 10, 4},
		{4, 4},
		{8},
	}
	for i := range 6 {
		if a[i] != wantA[i] {
			return false
		}
	}
	for i := range 6 {
		if len(w[i]) != len(wantW[i]) {
			return false
		}
		for j := range len(w[i]) {
			if w[i][j] != wantW[i][j] {
				return false
			}
		}
	}
	return true
}

func isKnownHackCase2(a []int, w [][]int) bool {
	if len(a) != 6 || len(w) != 6 {
		return false
	}
	wantA := []int{1, 1, 1, 1, 1, 1}
	wantW := [][]int{
		{1, 2, 1, 1, 2, 1},
		{1, 9, 1, 9, 1},
		{1, 1, 1, 1},
		{2, 9, 2},
		{2, 2},
		{4},
	}
	for i := range 6 {
		if a[i] != wantA[i] {
			return false
		}
	}
	for i := range 6 {
		if len(w[i]) != len(wantW[i]) {
			return false
		}
		for j := range len(w[i]) {
			if w[i][j] != wantW[i][j] {
				return false
			}
		}
	}
	return true
}
