package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ans := process(reader)
	fmt.Println(ans)
}

func process(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

// 线性基模板
type xorBasis []int

func (b xorBasis) insert(x int) bool {
	for i := len(b) - 1; i >= 0; i-- {
		if x>>i&1 == 0 {
			continue
		}
		if b[i] == 0 {
			b[i] = x
			return true
		}
		x ^= b[i]
	}
	return false
}

const H = 30

func solve(a []int) int {
	var x int
	for _, v := range a {
		x ^= v
	}
	if x == 0 {
		return -1
	}
	basis := make(xorBasis, H)

	var res int
	var cur int
	for _, v := range a {
		cur ^= v
		if cur != 0 && cur != x && basis.insert(cur) {
			res++
			cur = 0
		}
	}
	return max(res, 1)
}
