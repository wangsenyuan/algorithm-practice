package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	r, _ := os.Open("input.txt")
	reader := bufio.NewReader(r)

	w, _ := os.Create("output.txt")

	writer := bufio.NewWriter(w)

	defer writer.Flush()

	_, _, _, best, res := process(reader)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", best))
	for _, cur := range res {
		buf.WriteString(fmt.Sprintf("%d ", len(cur)))
		for _, v := range cur {
			buf.WriteString(fmt.Sprintf("%d ", v))
		}
		buf.WriteString("\n")
	}
	buf.WriteTo(writer)
}

func process(reader *bufio.Reader) (int, []int, [][]int, int, [][]int) {
	var n, v int
	fmt.Fscan(reader, &n, &v)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	var m int
	fmt.Fscan(reader, &m)
	friends := make([][]int, m)
	for i := range m {
		var l, r, f int
		fmt.Fscan(reader, &l, &r, &f)
		friends[i] = []int{l, r, f}
	}
	best, res := solve(v, a, friends)
	return v, a, friends, best, res
}

const inf = 1 << 60

type data struct {
	val  int
	from int
}

type friend struct {
	id int
	l  int
	r  int
	f  int
}

func solve(v int, a []int, friends [][]int) (int, [][]int) {
	m := len(friends)
	fs := make([]friend, m)
	for i := range m {
		fs[i] = friend{
			id: i,
			l:  friends[i][0],
			r:  friends[i][1],
			f:  friends[i][2],
		}
	}

	slices.SortFunc(fs, func(x, y friend) int {
		return x.f - y.f
	})

	x := slices.Max(a)

	n := len(a)
	dp := make([][]data, n+1)
	for i := range n + 1 {
		dp[i] = make([]data, x+1)
		for j := range x + 1 {
			dp[i][j].val = -inf
		}
	}

	dp[0][0].val = 0

	// fp[i] 表示在花费i的食物的情况下，能够获得的最大收益
	fp := make([]int, 2*x+1)

	feed := func(d int) {
		clear(fp)
		var sum int
		var cnt int
		for _, cur := range fs {
			if sum+cur.f > 2*x {
				break
			}
			if cur.l <= d && d <= cur.r {
				sum += cur.f
				cnt++
				fp[sum] = cnt
			}
		}
		for i := 1; i <= 2*x; i++ {
			fp[i] = max(fp[i], fp[i-1])
		}
	}

	// que := make([]int, x+1)

	// 有点悬呐
	// 400 * 400 * 400 = 64 * 1e6
	for d, u := range a {
		feed(d + 1)
		for i := 0; i <= x; i++ {
			if i+u < v || dp[d][i].val < 0 {
				continue
			}
			// 昨天剩余的至少要喂饱自己
			for y := 0; y <= i+u-v; y++ {
				// 那么优先肯定使用昨天的量
				w := min((u+i)-(y+v), u)
				if dp[d][i].val+fp[y] > dp[d+1][w].val {
					dp[d+1][w] = data{val: dp[d][i].val + fp[y], from: i}
				}
			}
		}
	}

	var best int
	for i := 0; i <= x; i++ {
		best = max(best, dp[n][i].val)
	}

	res := make([][]int, n)

	find := func(d int, cnt int) []int {
		var arr []int
		for _, cur := range fs {
			if len(arr) == cnt {
				break
			}
			if cur.l <= d && d <= cur.r {
				arr = append(arr, cur.id+1)
			}
		}
		return arr
	}

	construct := func(d int, r int) {
		for d > 0 {
			l := dp[d][r].from
			y := dp[d][r].val - dp[d-1][l].val
			res[d-1] = find(d, y)
			r = l
			d--
		}
	}

	for r := 0; r <= x; r++ {
		if dp[n][r].val == best {
			construct(n, r)
			break
		}
	}

	return best, res
}
