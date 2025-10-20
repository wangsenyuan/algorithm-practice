package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	res := drive(reader)
	for _, v := range res {
		fmt.Fprintln(writer, v)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	var m int
	fmt.Fscan(reader, &m)
	q := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &q[i])
	}
	return solve(k, a, q)
}

const inf = 1 << 60

func solve(k int, a []int, q []int) []int {

	mem := make(map[int]bool)
	for _, v := range a {
		mem[v] = true
	}

	find := func(amount int) int {
		res := inf

		for _, v := range a {
			for i := 1; i <= k && v*i <= amount; i++ {
				rest := amount - v*i
				if rest == 0 {
					res = min(res, i)
				}
				// 使用 w * j = rest
				for j := k - i; j > 0; j-- {
					w := rest / j
					if w*j > rest {
						break
					}
					if w < v && w*j == rest && mem[w] {
						res = min(res, i+j)
					}
				}
			}
		}
		return res
	}

	ans := make([]int, len(q))
	for i, amount := range q {
		tmp := find(amount)
		if tmp > k {
			ans[i] = -1
		} else {
			ans[i] = tmp
		}
	}

	return ans
}
