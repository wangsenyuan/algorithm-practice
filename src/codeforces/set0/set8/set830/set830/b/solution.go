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
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) int {
	n := len(a)
	tr := NewSegTree(n)
	sum := make(BIT, n+2)
	for i := range n {
		tr.Update(i, a[i])
		sum.Update(i, 1)
	}

	var head int

	var res int
	for cnt := n; cnt > 0; cnt-- {
		x := tr.Get(head, n)
		if 0 < head {
			y := tr.Get(0, head)
			if y.first < x.first {
				// 如果前半段有个更小的数字
				res += sum.Query(head, n)
				res += sum.Query(0, y.second)
				sum.Update(y.second, -1)
				tr.Update(y.second, inf)
				// 即使是当前被去掉的那个位置,是不是也没有关系？
				head = y.second
				continue
			}
		}
		res += sum.Query(head, x.second)
		sum.Update(x.second, -1)
		tr.Update(x.second, inf)
		head = x.second
	}

	return res
}

const inf = 1 << 60

type SegTree []pair

type pair struct {
	first  int
	second int
}

func min_pair(a, b pair) pair {
	if a.first < b.first || a.first == b.first && a.second < b.second {
		return a
	}
	return b
}

func NewSegTree(n int) SegTree {
	arr := make([]pair, 2*n)
	for i := n; i < len(arr); i++ {
		arr[i] = pair{inf, i - n}
	}

	for i := n - 1; i > 0; i-- {
		arr[i] = min_pair(arr[i*2], arr[i*2+1])
	}
	return arr
}

func (tr SegTree) Update(p int, v int) {
	n := len(tr) / 2
	p += n
	tr[p].first = v

	for p > 1 {
		tr[p>>1] = min_pair(tr[p], tr[p^1])
		p >>= 1
	}
}

func (tr SegTree) Get(l int, r int) pair {
	n := len(tr) / 2
	l += n
	r += n
	res := pair{inf, -1}
	for l < r {
		if l&1 == 1 {
			res = min_pair(res, tr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = min_pair(res, tr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}

type BIT []int

func (bit BIT) Update(i int, v int) {
	i++
	for i < len(bit) {
		bit[i] += v
		i += i & -i
	}
}

func (bit BIT) Pre(i int) int {
	i++
	var res int
	for i > 0 {
		res += bit[i]
		i -= i & -i
	}
	return res
}

func (bit BIT) Query(l int, r int) int {
	return bit.Pre(r) - bit.Pre(l-1)
}
