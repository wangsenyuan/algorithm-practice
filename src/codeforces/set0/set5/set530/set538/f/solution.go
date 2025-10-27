package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	p := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &p[i])
	}
	return solve(p)
}

type pair struct {
	first  int
	second int
}

func solve(p []int) []int {
	n := len(p)
	arr := make([]pair, n)
	for i := range n {
		arr[i] = pair{p[i], i}
	}
	slices.SortFunc(arr, func(a, b pair) int {
		// 同样的数字的时候，先处理最左边的
		return cmp.Or(a.first-b.first, a.second-b.second)
	})
	ans := make([]int, n)

	bit := make(BIT, n+3)

	for _, cur := range arr {
		i := cur.second + 1
		for k := 1; k*(i-1)+2 <= n && k < n; k++ {
			// 这个迭代的总的复杂性是 n * lgn
			ans[k] += bit.rangeSum(k*(i-1)+2, min(k*i+1, n))
		}
		bit.update(cur.second+1, 1)
	}

	return ans[1:]
}

type BIT []int

func (this BIT) update(i int, v int) {
	i++
	for i < len(this) {
		this[i] += v
		i += i & -i
	}
}

func (this BIT) get(i int) int {
	var res int
	i++
	for i > 0 {
		res += this[i]
		i -= i & -i
	}
	return res
}

func (this BIT) rangeSum(l int, r int) int {
	return this.get(r) - this.get(l-1)
}
