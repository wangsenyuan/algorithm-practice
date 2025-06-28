package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/bits"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	buf.WriteTo(os.Stdout)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) []int {
	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	queries := make([]string, m)
	for i := range m {
		queries[i] = readString(reader)
	}
	return solve(a, queries)
}

func solve(a []int, queries []string) []int {
	n := len(a)
	t := make(seg, 2<<bits.Len(uint(n)))
	t.build(a, 1, 1, n)

	var ans []int
	for _, cur := range queries {
		var l, r, k int
		pos := readInt([]byte(cur), 2, &l)
		pos = readInt([]byte(cur), pos+1, &r)
		readInt([]byte(cur), pos+1, &k)
		if cur[0] == '=' {
			t.update(1, l, r, k)
		} else {
			s := t.query(1, l, r)
			res, powL := 0, 1
			for j := k; j >= 0; j-- {
				res += s[j] * C[k][j] % mod * powL
				powL = powL * -(l - 1) % mod
			}
			ans = append(ans, (res%mod+mod)%mod)
		}
	}
	return ans
}

const mod = 1_000_000_007
const mx = 6

var C [mx][mx]int
var sPow [1e5 + 10][mx]int

func init() {
	for i := range mx {
		C[i][0] = 1
		for j := 1; j <= i; j++ {
			C[i][j] = C[i-1][j-1] + C[i-1][j]
		}
	}

	for i := 1; i < len(sPow); i++ {
		powI := 1
		for j := range mx {
			sPow[i][j] = (sPow[i-1][j] + powI) % mod
			powI = powI * i % mod
		}
	}
}

type seg []struct {
	l, r int
	sum  [mx]int
	todo int
}

func mergeInfo(a, b [mx]int) [mx]int {
	for i, v := range b {
		a[i] += v
	}
	return a
}

func (t seg) apply(o, x int) {
	cur := &t[o]
	for i := range mx {
		cur.sum[i] = x * (sPow[cur.r][i] - sPow[cur.l-1][i]) % mod
	}
	cur.todo = x
}

func (t seg) maintain(o int) {
	t[o].sum = mergeInfo(t[o<<1].sum, t[o<<1|1].sum)
}

func (t seg) spread(o int) {
	f := t[o].todo
	if f < 0 {
		return
	}
	t.apply(o<<1, f)
	t.apply(o<<1|1, f)
	t[o].todo = -1
}

func (t seg) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	t[o].todo = -1
	if l == r {
		t[o].sum[0] = a[l-1]
		for i := 1; i < mx; i++ {
			t[o].sum[i] = t[o].sum[i-1] * l % mod
		}
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) update(o, l, r, x int) {
	if l <= t[o].l && t[o].r <= r {
		t.apply(o, x)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, x)
	}
	if m < r {
		t.update(o<<1|1, l, r, x)
	}
	t.maintain(o)
}

func (t seg) query(o, l, r int) [mx]int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].sum
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	return mergeInfo(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}
