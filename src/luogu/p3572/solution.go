package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
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

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	d := readNNums(reader, n)
	q := readNum(reader)
	qs := make([]int, q)
	for i := 0; i < q; i++ {
		qs[i] = readNum(reader)
	}
	return solve(d, qs)
}

type pair struct {
	first  int
	second int
}

const inf = 1 << 60

func solve(d []int, queries []int) []int {
	n := len(d)
	que := make([]int, n)
	dp := make([]int, n)
	check := func(k int) int {
		var head, tail int

		que[head] = 0
		head++

		for i := 1; i < n; i++ {
			for tail < head && i-k > que[tail] {
				tail++
			}
			dp[i] = dp[que[tail]]
			if d[que[tail]] <= d[i] {
				dp[i]++
			}
			// 相同劳累度的情况下，尽量飞到远的，且高的位置上
			for head > tail && (dp[que[head-1]] > dp[i] || dp[que[head-1]] == dp[i] && d[que[head-1]] < d[i]) {
				head--
			}
			que[head] = i
			head++
		}

		return dp[n-1]
	}
	ans := make([]int, len(queries))

	for i, k := range queries {
		ans[i] = check(k)
	}

	return ans
}

func solve1(d []int, queries []int) []int {
	n := len(d)
	arr := make([]pair, n)
	for i, v := range d {
		arr[i] = pair{v, i}
	}

	slices.SortFunc(arr, func(a, b pair) int {
		if a.first != b.first {
			return a.first - b.first
		}
		return a.second - b.second
	})

	where := make([]int, n)
	pos := make([]int, n)
	for i := 0; i < n; {
		j := i
		for i < n && arr[i].first == arr[j].first {
			pos[arr[i].second] = i
			where[arr[i].second] = j
			i++
		}
	}

	tr := make(SegmentTree, 2*n)

	dp := make([]int, n+1)
	check := func(k int) int {
		for i := 0; i < 2*n; i++ {
			tr[i] = inf
		}

		for i, r := n-1, n-1; i >= 0; i-- {
			for r-k > i {
				tr.Update(pos[r], inf)
				r--
			}
			if i == n-1 {
				dp[i] = 0
			} else {
				j := where[i]
				dp[i] = min(tr.Query(j, n)+1, tr.Query(0, j), inf)
			}
			tr.Update(pos[i], dp[i])
		}

		return dp[0]
	}

	ans := make([]int, len(queries))

	for i, k := range queries {
		ans[i] = check(k)
	}

	return ans
}

type SegmentTree []int

func (t SegmentTree) Update(p int, v int) {
	n := len(t) / 2
	p += n
	t[p] = v
	for p > 1 {
		t[p>>1] = min(t[p], t[p^1])
		p >>= 1
	}
}

func (t SegmentTree) Query(l int, r int) int {
	n := len(t) / 2
	l += n
	r += n
	res := inf
	for l < r {
		if l&1 == 1 {
			res = min(res, t[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = min(res, t[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}
