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
	var n int
	fmt.Fscan(reader, &n)
	p := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &p[i])
	}
	return solve(p)
}

func solve(p []int) int {
	n := len(p)
	set := make(BIT, n+2)

	var cnt int
	for _, v := range p {
		cnt += set.get(n) - set.get(v)
		set.add(v)
	}

	ans := make([]int, cnt+2)
	ans[0] = 0
	ans[1] = 1
	for i := 2; i <= cnt; i++ {
		ans[i] = 4 + ans[i-2]
	}

	return ans[cnt]
}

type BIT []int

func (bit BIT) add(p int) {
	p++
	for p < len(bit) {
		bit[p]++
		p += p & -p
	}
}

func (bit BIT) get(p int) int {
	p++
	var res int
	for p > 0 {
		res += bit[p]
		p -= p & -p
	}
	return res
}
