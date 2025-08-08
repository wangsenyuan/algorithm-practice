package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ans := process(reader)
	var buf bytes.Buffer
	for _, x := range ans {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	fmt.Print(buf.String())
}

func process(reader *bufio.Reader) []int {
	var n, q int
	fmt.Fscan(reader, &n, &q)
	a := make([]int, n)
	w := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	for i := range n {
		fmt.Fscan(reader, &w[i])
	}
	queries := make([][]int, q)
	for i := range q {
		var x, y int
		fmt.Fscan(reader, &x, &y)
		queries[i] = []int{x, y}
	}
	return solve(a, w, queries)
}

type BIT []int

func (bit BIT) update(i int, v int) {
	for i < len(bit) {
		bit[i] += v
		i += i & -i
	}
}

func (bit BIT) query(i int) int {
	var res int
	for i > 0 {
		res += bit[i]
		i -= i & -i
	}
	return res
}

func (bit BIT) queryRange(l, r int) int {
	return bit.query(r) - bit.query(l-1)
}

func (f BIT) kth(k int) (res int) {
	for b := 1 << 17; b > 0; b >>= 1 {
		if nxt := res | b; nxt < len(f) && f[nxt] < k {
			k -= f[nxt]
			res = nxt
		}
	}
	return res + 1
}

const mod = 1e9 + 7

func solve(a []int, w []int, queries [][]int) []int {
	n := len(w)
	f1 := make(BIT, n+1)
	f2 := make(BIT, n+1)
	for i := range n {
		a[i] -= (i + 1)
		f1.update(i+1, w[i])
		f2.update(i+1, (w[i]*a[i]%mod+mod)%mod)
	}

	var ans []int

	for _, cur := range queries {
		if cur[0] < 0 {
			j := -cur[0]
			diff := cur[1] - w[j-1]
			w[j-1] += diff
			f1.update(j, diff)
			f2.update(j, ((diff+mod)%mod)*a[j-1]%mod)
		} else {
			l, r := cur[0], cur[1]
			s := (f1.query(l-1) + f1.query(r) + 1) / 2
			m := f1.kth(s)
			x := a[m-1]
			tmp := f1.queryRange(l, m)%mod*x - f2.queryRange(l, m)%mod + f2.queryRange(m, r)%mod - f1.queryRange(m, r)%mod*x%mod
			tmp %= mod
			if tmp < 0 {
				tmp += mod
			}
			ans = append(ans, tmp)
		}
	}
	return ans
}
