package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for _, v := range res {
		if v {
			writer.WriteString("YES\n")
		} else {
			writer.WriteString("NO\n")
		}
	}
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func readNums(reader *bufio.Reader) []int {
	s := readString(reader)
	ss := strings.Split(s, " ")
	nums := make([]int, len(ss))
	for i := range len(ss) {
		nums[i], _ = strconv.Atoi(ss[i])
	}
	return nums
}

func drive(reader *bufio.Reader) []bool {
	first := readNums(reader)
	s := readString(reader)
	m, k := first[1], first[2]
	queries := make([][]int, m+k)
	for i := range m + k {
		queries[i] = readNums(reader)
	}
	return solve(s, queries)
}

func solve(s string, queries [][]int) []bool {
	n := len(s)
	tr1 := NewTree(n, 11, 1000000007)
	tr2 := NewTree(n, 13, 1073676287)

	for i := range n {
		x := int(s[i]-'0') + 1
		tr1.Update(i, i, x)
		tr2.Update(i, i, x)
	}

	check := func(l int, r int, d int) bool {
		if l == r {
			return true
		}
		if tr1.Get(l, r-d)*tr1.st[d]%tr1.mod != tr1.Get(l+d, r) {
			return false
		}
		if tr2.Get(l, r-d)*tr2.st[d]%tr2.mod != tr2.Get(l+d, r) {
			return false
		}
		return true
	}

	var ans []bool

	for _, cur := range queries {
		l, r, d := cur[1], cur[2], cur[3]
		l--
		r--
		if cur[0] == 1 {
			tr1.Update(l, r, d+1)
			tr2.Update(l, r, d+1)
		} else {
			ans = append(ans, check(l, r, d))
		}
	}
	return ans
}

type Tree struct {
	st   []int
	dp   []int
	val  []int
	lazy []int
	sz   int
	mod  int
}

func NewTree(n int, bs int, mod int) *Tree {
	st := make([]int, n+1)
	st[0] = 1
	dp := make([]int, n+1)
	dp[0] = 1
	for i := range n {
		st[i+1] = (st[i] * bs) % mod
		dp[i+1] = (dp[i] + st[i+1]) % mod
	}

	val := make([]int, 4*n)
	lazy := make([]int, 4*n)
	sz := n
	return &Tree{st, dp, val, lazy, sz, mod}
}

func (tr *Tree) apply(i int, v int, l int, r int) {
	tr.val[i] = v * (tr.dp[r+1] - tr.dp[l] + tr.mod) % tr.mod
	tr.lazy[i] = v
}

func (tr *Tree) push(i int, l int, r int) {
	if l < r && tr.lazy[i] != 0 {
		mid := (l + r) >> 1
		tr.apply(i*2+1, tr.lazy[i], l, mid)
		tr.apply(i*2+2, tr.lazy[i], mid+1, r)
		tr.lazy[i] = 0
	}
}

func (tr *Tree) Update(L int, R int, v int) {
	var f func(i int, l int, r int, L int, R int)
	f = func(i int, l int, r int, L int, R int) {
		if l == L && r == R {
			tr.apply(i, v, l, r)
			return
		}
		tr.push(i, l, r)
		mid := (l + r) >> 1
		if L <= mid {
			f(i*2+1, l, mid, L, min(mid, R))
		}
		if mid < R {
			f(i*2+2, mid+1, r, max(mid+1, L), R)
		}
		tr.val[i] = (tr.val[i*2+1] + tr.val[i*2+2]) % tr.mod
	}
	f(0, 0, tr.sz-1, L, R)
}

func (tr *Tree) Get(L int, R int) int {
	if L > R {
		return 0
	}
	var f func(i int, l int, r int, L int, R int) int
	f = func(i int, l int, r int, L int, R int) int {
		if l == L && r == R {
			return tr.val[i]
		}
		tr.push(i, l, r)
		mid := (l + r) >> 1
		var res int
		if L <= mid {
			res = f(i*2+1, l, mid, L, min(mid, R))
		}
		if mid < R {
			res += f(i*2+2, mid+1, r, max(mid+1, L), R)
		}
		return res % tr.mod
	}
	return f(0, 0, tr.sz-1, L, R)
}
