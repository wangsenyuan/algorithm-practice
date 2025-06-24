package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func process(reader *bufio.Reader) int {
	n, c := readTwoNums(reader)
	m := readNum(reader)
	events := make([][]int, m)
	for i := range events {
		events[i] = readNNums(reader, 2)
	}
	return solve(n, c, events)
}

func solve(n int, c int, events [][]int) int {
	tr1 := NewSegTree(n)
	tr2 := NewSegTree(n)
	// dp[0] = 0
	tr1.Update(0, 0)
	tr2.Update(0, 0)

	var ans int

	for _, evt := range events {
		u, p := evt[0], evt[1]
		u--
		// 在位置t举办一次会议, 收益为p
		tmp1 := p - c*u + tr1.Get(0, u+1)
		tmp2 := p + c*u + tr2.Get(u, n)
		tmp := max(tmp1, tmp2)
		ans = max(ans, tmp)

		tr1.Update(u, tmp+c*u)
		tr2.Update(u, tmp-c*u)
	}
	return ans
}

const inf = 1 << 60

type SegTree []int

func NewSegTree(n int) SegTree {
	arr := make([]int, 2*n)
	for i := range arr {
		arr[i] = -inf
	}
	return SegTree(arr)
}

func (tr SegTree) Update(p int, v int) {
	n := len(tr) / 2
	p += n
	tr[p] = v

	for p > 1 {
		tr[p>>1] = max(tr[p], tr[p^1])
		p >>= 1
	}
}

func (tr SegTree) Get(l int, r int) int {
	n := len(tr) / 2
	l += n
	r += n
	var res int = -inf
	for l < r {
		if l&1 == 1 {
			res = max(res, tr[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = max(res, tr[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
