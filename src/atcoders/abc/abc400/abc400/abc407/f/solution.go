package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	res := drive(reader)
	for _, v := range res {
		writer.WriteString(strconv.Itoa(v))
		writer.WriteByte('\n')
	}
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	A := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &A[i])
	}
	return solve(A)
}

// solve returns a slice of length len(A); the i-th entry (0-based) is the sum
// of maxima over all contiguous subarrays of length i+1.
func solve(a []int) []int {
	// 假设 a[i]在区间 [l..r] 中作为最大值
	// x = i - l + 1, y = r - i + 1, m = r - l + 1
	// 对长度 k 来说，包含 i 且位于 [l..r] 的窗口数量是：
	// 1) k <= min(x, y): k
	// 2) min(x, y) < k <= max(x, y): min(x, y)
	// 3) max(x, y) < k <= m: m - k + 1
	// 这三个部分分别是一次函数，可以用两个差分数组表示 a*k+b
	n := len(a)

	stack := make([]int, n)
	L := make([]int, n)
	R := make([]int, n)
	var top int
	for i, v := range a {
		for top > 0 && a[stack[top-1]] <= v {
			top--
		}
		L[i] = 0
		if top > 0 {
			L[i] = stack[top-1] + 1
		}
		stack[top] = i
		top++
	}

	top = 0
	for i := n - 1; i >= 0; i-- {
		for top > 0 && a[stack[top-1]] < a[i] {
			top--
		}
		R[i] = n - 1
		if top > 0 {
			R[i] = stack[top-1] - 1
		}
		stack[top] = i
		top++
	}

	diff1 := make([]int, n+2)
	diff2 := make([]int, n+2)

	add := func(l int, r int, a int, b int) {
		if l > r {
			return
		}
		diff1[l] += a
		diff1[r+1] -= a
		diff2[l] += b
		diff2[r+1] -= b
	}

	for i := range n {
		l, r := L[i], R[i]
		x := i - l + 1
		y := r - i + 1
		k1 := min(x, y)
		k2 := max(x, y)
		m := r - l + 1
		// 当 k <= k1时, 它们的贡献 = k * a[i]
		add(1, k1, a[i], 0)
		// 当 k1 < k <= k2 时，贡献保持为 k1 * a[i]
		add(k1+1, k2, 0, k1*a[i])
		// 当 k2 < k <= m 时, 贡献 = (m - k + 1) * a[i]
		add(k2+1, m, -a[i], (m+1)*a[i])
	}

	ans := make([]int, n+1)

	for i := 1; i <= n; i++ {
		diff1[i] += diff1[i-1]
		diff2[i] += diff2[i-1]
		ans[i] = i*diff1[i] + diff2[i]
	}

	return ans[1:]
}
