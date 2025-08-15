package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	var m int
	fmt.Fscan(reader, &m)
	res := solve(a, m)
	var buf strings.Builder
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, r := range res {
		buf.WriteString(fmt.Sprintf("%d %d ", r[0], r[1]))
	}
	return strings.TrimSpace(buf.String())
}

func solve(a []int, m int) [][]int {
	n := len(a)
	for i := range a {
		a[i]--
	}

	var cycles [][]int
	marked := make([]bool, n)
	belong := make([]int, n)
	for i := range n {
		if !marked[i] {
			j := i
			var cur []int
			for !marked[j] {
				cur = append(cur, j)
				belong[j] = len(cycles)
				marked[j] = true
				j = a[j]
			}
			cycles = append(cycles, cur)
		}
	}

	k := abs(n - m - len(cycles))

	if k == 0 {
		return nil
	}

	var res [][]int
	if n-m < len(cycles) {
		// 使用包含1的环，去合并其他的环
		x := belong[0]
		for i := 1; i < n && len(res) < k; i++ {
			if belong[i] != x {
				res = append(res, []int{1, i + 1})
				y := belong[i]
				for _, j := range cycles[y] {
					belong[j] = x
				}
			}
		}
		return res
	}

	// 那么要打破k个环
	for len(res) < k {
		var i int
		for len(cycles[belong[i]]) == 1 {
			i++
		}
		// i找到了
		j := -1
		for _, u := range cycles[belong[i]] {
			if u != i && (j == -1 || u < j) {
				j = u
			}
		}

		res = append(res, []int{i + 1, j + 1})
		// 要重新计算
		a[i], a[j] = a[j], a[i]

		clear(marked)
		cycles = cycles[:0]
		for i := range n {
			if !marked[i] {
				j := i
				var cur []int
				for !marked[j] {
					cur = append(cur, j)
					belong[j] = len(cycles)
					marked[j] = true
					j = a[j]
				}
				cycles = append(cycles, cur)
			}
		}
	}

	return res
}

func abs(num int) int {
	return max(num, -num)
}
