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

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) int {
	n := len(a)

	// a[i] <= n
	row := make([]int, n+2)
	expireAt := make([][]int, n+2)

	sum := make(bit, n+3)

	for i := range n {
		row[1]++
		row[a[i]+1]--
		expireAt[a[i]] = append(expireAt[a[i]], i)
		sum.update(i, i)
	}

	for i := 1; i <= n; i++ {
		row[i] += row[i-1]
	}

	var moves int

	for i := 1; i <= n; i++ {
		l := n - row[i]
		// l + (l+1) + ... + (n - 1)
		tot := (l + n - 1) * (n - l) / 2
		moves += tot - sum.query(n)
		for _, j := range expireAt[i] {
			sum.update(j, -j)
		}
	}

	// 现在需要一个range update, point query的数据结果
	var extra int

	bit2 := make(bit2, n+3)

	for i := range n {
		// 如果把a[i]给删除掉
		c := bit2.query(a[i])
		// a[i]前面有c个元素
		extra = max(extra, c)
		bit2.update(a[i], 1)
	}

	return moves + extra
}

type bit []int

func (bit bit) update(i int, val int) {
	i++
	for i < len(bit) {
		bit[i] += val
		i += i & -i
	}
}

func (bit bit) query(i int) int {
	var res int
	i++
	for i > 0 {
		res += bit[i]
		i -= i & -i
	}
	return res
}
func (bit bit) queryRange(l int, r int) int {
	return bit.query(r) - bit.query(l-1)
}

type bit2 []int

func (bit2 bit2) update(i int, val int) {
	i++
	for i > 0 {
		bit2[i] += val
		i -= i & -i
	}
}

func (bit2 bit2) query(i int) int {
	var res int
	i++
	for i < len(bit2) {
		res += bit2[i]
		i += i & -i
	}
	return res
}
