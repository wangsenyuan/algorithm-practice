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

	for _, row := range drive(reader) {
		for i, v := range row {
			if i > 0 {
				fmt.Fprint(writer, " ")
			}
			fmt.Fprint(writer, v)
		}
		fmt.Fprintln(writer)
	}
}

func drive(reader *bufio.Reader) [][]int {
	var tc int
	fmt.Fscan(reader, &tc)
	res := make([][]int, tc)
	for i := range tc {
		var n int
		fmt.Fscan(reader, &n)
		p := make([]int, n)
		for j := range n {
			fmt.Fscan(reader, &p[j])
		}
		res[i] = solve(p)
	}
	return res
}

func solve(p []int) []int {
	n := len(p)
	todo := make([][]int, n)
	for i := range n - 1 {
		val := abs(p[i+1] - p[i])
		todo[val] = append(todo[val], i)
	}

	L := NewDSU(n)
	R := NewDSU(n)

	ans := make([]int, n)

	var sum int

	count := func(l int, r int) int {
		return (r - l) * (r - l + 1) / 2
	}

	for k := n - 1; k > 0; k-- {
		for _, l := range todo[k] {
			r := l + 1
			l1 := L.Find(l)
			r1 := R.Find(r)
			sum -= count(l1, l)
			sum -= count(r, r1)
			sum += count(l1, r1)
			R.arr[l] = r
			L.arr[r] = l
		}
		ans[k] = sum
	}

	return ans[1:]
}

func abs(num int) int {
	return max(num, -num)
}

type DSU struct {
	arr []int
	cnt []int
}

func NewDSU(n int) *DSU {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}
	return &DSU{arr, cnt}
}

func (this *DSU) Find(x int) int {
	if this.arr[x] != x {
		this.arr[x] = this.Find(this.arr[x])
	}
	return this.arr[x]
}
