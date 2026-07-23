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

	for _, v := range drive(reader) {
		fmt.Fprintln(writer, v)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(m, a)
}

func solve(m int, a []int) []int {
	n := len(a)

	var ans int
	bit := NewBIT(m)

	delta := make([]int, m)

	for i := range n {
		v := a[i] % m
		ans += bit.get(m) - bit.get(v)
		bit.add(v, 1)
		delta[v] += 2*i - n + 1
	}

	var f []int
	for i := m - 1; i >= 0; i-- {
		f = append(f, ans)
		ans += delta[i]
	}

	return f
}

type BIT []int

func NewBIT(n int) BIT {
	return make(BIT, n+3)
}

func (bit BIT) add(i int, v int) {
	i++
	for i < len(bit) {
		bit[i] += v
		i += i & -i
	}
}

func (bit BIT) get(i int) int {
	i++
	var res int
	for i > 0 {
		res += bit[i]
		i -= i & -i
	}
	return res
}
