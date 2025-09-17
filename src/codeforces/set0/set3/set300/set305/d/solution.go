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

func drive(reader *bufio.Reader) int {
	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)
	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		edges[i] = []int{u, v}
	}
	return solve(n, edges, k)
}

const mod = 1000000007

func mul(a, b int) int {
	return a * b % mod
}

func pow(a, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = mul(r, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return r
}

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func solve(n int, edges [][]int, k int) int {
	marked := make([]bool, n+1)

	for _, cur := range edges {
		u, v := cur[0], cur[1]
		if v != u+1 && v != u+k+1 {
			return 0
		}
		// 要么连接到下一个位置，要么连接到 u + k + 1, 否则就是错误的连接
		if v == u+k+1 {
			marked[u] = true
		}
	}

	// 如果额外的边都不存在的情况下，那么就可以选择添加的方式，如果存在，它们必须在一个区间内
	var first_special int
	var last_special int
	for i := 1; i <= n; i++ {
		if marked[i] {
			if first_special > 0 && i-first_special > k {
				return 0
			}
			if first_special == 0 {
				first_special = i
			}
			last_special = i
		}
	}

	if first_special > 0 {
		var sum int
		for i := first_special; i < last_special; i++ {
			if !marked[i] {
				sum++
			}
		}

		// 这个区间居然是滚动的～
		var res int

		for i := max(1, last_special-k); i <= first_special && i+k+1 <= n; i++ {
			// 如果第一条边是从i出发的
			tmp := first_special - i + sum
			if i < first_special {
				tmp--
			}
			if last_special < min(n-k-1, i+k) {
				tmp += min(n-k-1, i+k) - last_special
			}
			res = add(res, pow(2, tmp))
		}
		return res
	}

	res := 1

	// no edges yet
	for i := 1; i+k+1 <= n; i++ {
		//如果最后一条边是从i发出的
		cnt := min(i-1, k)
		res = add(res, pow(2, cnt))
	}

	return res
}
