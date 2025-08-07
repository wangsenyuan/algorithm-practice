package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d ", x))
	}
	buf.WriteTo(os.Stdout)
}

func process(reader *bufio.Reader) []int {
	var n, x int
	fmt.Fscan(reader, &n, &x)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(n, x, a)
}

func solve(n int, x int, a []int) []int {
	x--
	freq := make([]int, n+1)
	sz := make([]int, n)

	que := make([]int, n)
	var head, tail int
	next := make([]int, n)
	for i := range n {
		next[i] = -1
	}
	for i := 0; i < n; i++ {
		if a[i] == 0 {
			que[head] = i
			head++
			sz[i] = 1
			freq[1]++
		} else {
			next[a[i]-1] = i
		}
	}

	for tail < head {
		u := que[tail]
		tail++
		v := next[u]
		if v >= 0 {
			que[head] = v
			head++
			sz[v] = sz[u] + 1
			freq[sz[u]]--
			freq[sz[v]]++
		}
	}
	// 需要把x的给处理掉
	y := x
	for next[y] != -1 {
		y = next[y]
	}
	// 把x所在的链条特殊处理
	freq[sz[y]]--

	var arr []int

	for i := 1; i <= n; i++ {
		for freq[i] > 0 {
			arr = append(arr, i)
			freq[i]--
		}
	}

	ok := make([]bool, n+1)
	ok[0] = true

	for _, i := range arr {
		for j := n; j >= i; j-- {
			if ok[j-i] {
				ok[j] = true
			}
		}
	}

	var ans []int

	for i := 0; i+sz[x] <= n; i++ {
		if ok[i] {
			ans = append(ans, i+sz[x])
		}
	}

	return ans
}
