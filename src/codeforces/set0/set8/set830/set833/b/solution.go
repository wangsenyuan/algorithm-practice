package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, k := readTwoNums(reader)
	a := readNNums(reader, n)
	res := solve(a, k)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(a []int, k int) int {
	// dp[i][k] 表示前i个元素在k段中的最优解
	// dp[i][k] = max(dp[j][k-1] + distinct count between j+1 ... i)
	// n * n * k TLE
	n := len(a)

	pos := make([]int, n+1)
	for i := 0; i <= n; i++ {
		pos[i] = -1
	}

	prev := make([]int, n)

	for i := 0; i < n; i++ {
		prev[i] = pos[a[i]]
		pos[a[i]] = i
	}

	tr := NewSegTree(n)
	// dp[i] = 当前层i的最大值
	dp := make([]int, n+1)
	for x := 0; x < k; x++ {
		tr.Init(dp[:n])
		for i := 0; i < n; i++ {
			tr.Update(prev[i]+1, i, 1)
			dp[i+1] = tr.Get(0, i)
		}
	}

	return dp[n]
}

type SetTree struct {
	add []int
	val []int
	sz  int
}

func NewSegTree(n int) *SetTree {
	add := make([]int, 4*n)
	val := make([]int, 4*n)
	return &SetTree{add, val, n}
}

func (t *SetTree) Init(val []int) {
	var loop func(i int, l int, r int)

	loop = func(i int, l int, r int) {
		t.add[i] = 0
		if l == r {
			t.val[i] = val[l]
			return
		}
		mid := (l + r) / 2
		loop(i*2+1, l, mid)
		loop(i*2+2, mid+1, r)
		t.val[i] = max(t.val[i*2+1], t.val[i*2+2])
	}

	loop(0, 0, t.sz-1)
}

func (t *SetTree) pushLazy(i int, v int) {
	t.add[i] += v
	t.val[i] += v
}

func (t *SetTree) push(i int) {
	if t.add[i] != 0 {
		t.pushLazy(2*i+1, t.add[i])
		t.pushLazy(2*i+2, t.add[i])
		t.add[i] = 0
	}
}

func (t *SetTree) Update(l int, r int, v int) {
	var loop func(i int, L int, R int)

	loop = func(i int, L int, R int) {
		if r < L || R < l {
			return
		}
		if l <= L && R <= r {
			t.pushLazy(i, v)
			return
		}
		t.push(i)
		mid := (L + R) / 2
		if l <= mid {
			loop(2*i+1, L, mid)
		}
		if mid < r {
			loop(2*i+2, mid+1, R)
		}
		t.val[i] = max(t.val[2*i+1], t.val[2*i+2])
	}
	loop(0, 0, t.sz-1)
}

func (t *SetTree) Get(l int, r int) int {
	var loop func(i int, L int, R int) int

	loop = func(i int, L int, R int) int {
		if r < L || R < l {
			return 0
		}
		if l <= L && R <= r {
			return t.val[i]
		}
		t.push(i)
		mid := (L + R) / 2
		var a, b int
		if l <= mid {
			a = loop(2*i+1, L, mid)
		}
		if mid < r {
			b = loop(2*i+2, mid+1, R)
		}

		return max(a, b)
	}

	return loop(0, 0, t.sz-1)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
