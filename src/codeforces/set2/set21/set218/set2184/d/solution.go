package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for _, v := range drive(reader) {
		fmt.Fprintln(writer, v)
	}
}

func drive(reader *bufio.Reader) []int {
	var t int
	fmt.Fscan(reader, &t)
	res := make([]int, t)
	for i := range t {
		var n, k int
		fmt.Fscan(reader, &n, &k)
		res[i] = solve(n, k)
	}
	return res
}

func solve(n, k int) int {
	var ds []int
	for i := n; i > 0; i >>= 1 {
		ds = append(ds, i&1)
	}

	m := len(ds)

	if k > 2*m {
		return 0
	}
	slices.Reverse(ds)
	// k <= 2 * len(ds)

	C := make([][]int, m+1)
	for i := range C {
		C[i] = make([]int, m+1)
		C[i][0] = 1
		C[i][i] = 1
		for j := 1; j < i; j++ {
			C[i][j] = C[i-1][j-1] + C[i-1][j]
		}
	}

	var ans int

	for i := 1; i < m; i++ {
		// 剩下 i-1个位置, 如果其中有j个1
		for j := range i {
			if i+j > k {
				ans += C[i-1][j]
			}
		}
	}

	var pref int
	for i, v := range ds {
		if v == 1 && pref > 0 {
			for j := range m - i - 1 {
				if m-1+pref+j > k {
					ans += C[m-i-1][j]
				}
			}
		}
		pref += v
	}

	if m-1+pref > k {
		ans++
	}

	return ans
}
