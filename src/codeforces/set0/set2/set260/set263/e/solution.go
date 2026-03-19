package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ans := drive(reader)
	fmt.Printf("%d %d\n", ans[0], ans[1])
}

func drive(reader *bufio.Reader) []int {
	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)

	a := make([][]int, n)
	for i := range n {
		a[i] = make([]int, m)
		for j := range m {
			fmt.Fscan(reader, &a[i][j])
		}
	}

	return solve(k, a)
}

func solve(k int, a [][]int) []int {
	mat := a
	n := len(mat)
	m := len(mat[0])

	t1 := calcQuadrantContribution(k, mat)
	r1 := rotate(mat)
	t2 := calcQuadrantContribution(k, r1)
	r2 := rotate(r1)
	t3 := calcQuadrantContribution(k, r2)
	r3 := rotate(r2)
	t4 := calcQuadrantContribution(k, r3)

	horizontal := calcHorizontalAxis(k, mat)
	vertical := calcVerticalAxis(k, mat)

	best := -1
	ans := []int{k, k}

	for i := k - 1; i <= n-k; i++ {
		for j := k - 1; j <= m-k; j++ {
			score := t1[i][j] +
				t2[m-1-j][i] +
				t3[n-1-i][m-1-j] +
				t4[j][n-1-i] -
				horizontal[i][j] -
				vertical[i][j] -
				k*mat[i][j]
			if score > best {
				best = score
				ans[0] = i + 1
				ans[1] = j + 1
			}
		}
	}

	return ans
}

func rotate(a [][]int) [][]int {
	n := len(a)
	m := len(a[0])
	res := make([][]int, m)
	for i := range m {
		res[i] = make([]int, n)
	}
	for i := range n {
		for j := range m {
			res[m-1-j][i] = a[i][j]
		}
	}
	return res
}

func calcQuadrantContribution(k int, a [][]int) [][]int {
	n := len(a)
	m := len(a[0])

	weighted := make([][]int, n)
	for i := range n {
		weighted[i] = make([]int, m)
		for j := range m {
			weighted[i][j] = (i + j) * a[i][j]
		}
	}

	cnt := triangleSum(k, a)
	sum := triangleSum(k, weighted)

	res := make([][]int, n)
	for i := range n {
		res[i] = make([]int, m)
		for j := range m {
			res[i][j] = sum[i][j] + (k-i-j)*cnt[i][j]
		}
	}
	return res
}

func triangleSum(k int, val [][]int) [][]int {
	n := len(val)
	m := len(val[0])

	col := make([][]int, n+1)
	for i := range n + 1 {
		col[i] = make([]int, m)
	}
	for i := range n {
		for j := range m {
			col[i+1][j] = col[i][j] + val[i][j]
		}
	}

	diag := make([][]int, n+m-1)
	for d := range n + m - 1 {
		diag[d] = make([]int, n+1)
		for i := range n {
			diag[d][i+1] = diag[d][i]
			j := d - i
			if j >= 0 && j < m {
				diag[d][i+1] += val[i][j]
			}
		}
	}

	res := make([][]int, n)
	for x := range n {
		res[x] = make([]int, m)
		lower := max(0, x-k+1)
		for y := range m {
			add := col[x+1][y] - col[lower][y]
			if y == 0 {
				res[x][y] = add
				continue
			}

			res[x][y] = res[x][y-1] + add
			d := x + y - k
			if d < 0 {
				continue
			}

			left := max(lower, d-(m-1))
			right := min(x, d)
			if left <= right {
				res[x][y] -= diag[d][right+1] - diag[d][left]
			}
		}
	}

	return res
}

func calcHorizontalAxis(k int, a [][]int) [][]int {
	n := len(a)
	res := make([][]int, n)
	for i := range n {
		res[i] = weightedLine(a[i], k)
	}
	return res
}

func calcVerticalAxis(k int, a [][]int) [][]int {
	n := len(a)
	m := len(a[0])
	res := make([][]int, n)
	for i := range n {
		res[i] = make([]int, m)
	}

	col := make([]int, n)
	for j := range m {
		for i := range n {
			col[i] = a[i][j]
		}
		cur := weightedLine(col, k)
		for i := range n {
			res[i][j] = cur[i]
		}
	}
	return res
}

func weightedLine(arr []int, k int) []int {
	n := len(arr)
	pref := make([]int, n+1)
	prefIdx := make([]int, n+1)
	for i, v := range arr {
		pref[i+1] = pref[i] + v
		prefIdx[i+1] = prefIdx[i] + i*v
	}

	res := make([]int, n)
	for pos := range n {
		left := max(0, pos-k+1)
		right := min(n-1, pos+k-1)

		sumLeft := pref[pos+1] - pref[left]
		idxLeft := prefIdx[pos+1] - prefIdx[left]
		sumRight := pref[right+1] - pref[pos+1]
		idxRight := prefIdx[right+1] - prefIdx[pos+1]

		res[pos] = (k-pos)*sumLeft + idxLeft + (k+pos)*sumRight - idxRight
	}
	return res
}
