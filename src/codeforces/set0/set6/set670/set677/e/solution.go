package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	return s
}

func drive(reader *bufio.Reader) int {
	s := readString(reader)
	n, _ := strconv.Atoi(s)
	a := make([]string, n)
	for i := range n {
		a[i] = readString(reader)
	}
	return solve(a)
}

const mod = 1_000_000_007

func mul(a, b int) int {
	return a * b % mod
}

func pow(a, b int) int {
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res = mul(res, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return res
}

type data struct {
	d   int
	cnt [3]int
}

func (a data) add(b data) data {
	for y := range 3 {
		a.cnt[y] += b.cnt[y]
	}
	return a
}

func solve(a []string) int {
	n := len(a)
	var factor float64 = math.Log2(3)

	dp := make([][][]data, n)

	for i := range n {
		dp[i] = make([][]data, n)
		for j := range n {
			dp[i][j] = make([]data, 8)
		}
	}

	// dp[i][j][0] 表示左上角, dp[i][j][1], 表示top
	// dp[i][j][2] 表示右上角, dp[i][j][3] 表示left

	leftTop := func(i int, j int, x int) {
		dp[i][j][0].cnt[x-1] = 1
		dp[i][j][0].d = 1
		if i > 0 && j > 0 {
			dp[i][j][0].d += dp[i-1][j-1][0].d
			for y := range 3 {
				dp[i][j][0].cnt[y] += dp[i-1][j-1][0].cnt[y]
			}
		}
	}

	top := func(i int, j int, x int) {
		dp[i][j][1].cnt[x-1] = 1
		dp[i][j][1].d = 1
		if i > 0 {
			dp[i][j][1].d += dp[i-1][j][1].d
			for y := range 3 {
				dp[i][j][1].cnt[y] += dp[i-1][j][1].cnt[y]
			}
		}
	}

	rightTop := func(i int, j int, x int) {
		dp[i][j][2].cnt[x-1] = 1
		dp[i][j][2].d = 1
		if i > 0 && j < n-1 {
			dp[i][j][2].d += dp[i-1][j+1][2].d
			for y := range 3 {
				dp[i][j][2].cnt[y] += dp[i-1][j+1][2].cnt[y]
			}
		}
	}

	left := func(i int, j int, x int) {
		dp[i][j][3].cnt[x-1] = 1
		dp[i][j][3].d = 1
		if j > 0 {
			dp[i][j][3].d += dp[i][j-1][3].d
			for y := range 3 {
				dp[i][j][3].cnt[y] += dp[i][j-1][3].cnt[y]
			}
		}
	}

	for i := range n {
		for j := range n {
			if a[i][j] == '0' {
				continue
			}
			x := int(a[i][j] - '0')
			leftTop(i, j, x)
			top(i, j, x)
			rightTop(i, j, x)
			left(i, j, x)
		}
	}

	right := func(i int, j int, x int) {
		dp[i][j][4].cnt[x-1] = 1
		dp[i][j][4].d = 1
		if j < n-1 {
			dp[i][j][4].d += dp[i][j+1][4].d
			for y := range 3 {
				dp[i][j][4].cnt[y] += dp[i][j+1][4].cnt[y]
			}
		}
	}
	rightBot := func(i int, j int, x int) {
		dp[i][j][5].cnt[x-1] = 1
		dp[i][j][5].d = 1
		if i < n-1 && j < n-1 {
			dp[i][j][5].d += dp[i+1][j+1][5].d

			for y := range 3 {
				dp[i][j][5].cnt[y] += dp[i+1][j+1][5].cnt[y]
			}
		}
	}
	bottom := func(i int, j int, x int) {
		dp[i][j][6].cnt[x-1] = 1
		dp[i][j][6].d = 1
		if i < n-1 {
			dp[i][j][6].d += dp[i+1][j][6].d
			for y := range 3 {
				dp[i][j][6].cnt[y] += dp[i+1][j][6].cnt[y]
			}
		}
	}
	leftBot := func(i int, j int, x int) {
		dp[i][j][7].cnt[x-1] = 1
		dp[i][j][7].d = 1
		if i < n-1 && j > 0 {
			dp[i][j][7].d += dp[i+1][j-1][7].d
			for y := range 3 {
				dp[i][j][7].cnt[y] += dp[i+1][j-1][7].cnt[y]
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if a[i][j] == '0' {
				continue
			}
			x := int(a[i][j] - '0')
			right(i, j, x)
			rightBot(i, j, x)
			bottom(i, j, x)
			leftBot(i, j, x)
		}
	}

	var res data

	update := func(cur data) {
		if cur.cnt[1] >= res.cnt[1] && cur.cnt[2] >= res.cnt[2] {
			res = cur
			return
		}
		if cur.cnt[1] < res.cnt[1] && cur.cnt[2] < res.cnt[2] {
			return
		}
		if cur.cnt[1] > res.cnt[1] {
			// cur.cnt[2] < res.cnt[2]
			w := cur.cnt[1] - res.cnt[1]
			v := res.cnt[2] - cur.cnt[2]
			if float64(w) >= float64(v)*factor {
				res = cur
			}
			return
		}
		// 3的数量更多
		w := cur.cnt[2] - res.cnt[2]
		v := res.cnt[1] - cur.cnt[1]
		if float64(v) <= float64(w)*factor {
			res = cur
		}
	}

	getLeftTop := func(i int, j int, d int) data {
		res := dp[i][j][0]
		if d < dp[i][j][0].d {
			for y := range 3 {
				res.cnt[y] -= dp[i-d][j-d][0].cnt[y]
			}
		}
		return res
	}

	getRightTop := func(i int, j int, d int) data {
		res := dp[i][j][2]
		if d < dp[i][j][2].d {
			for y := range 3 {
				res.cnt[y] -= dp[i-d][j+d][2].cnt[y]
			}
		}
		return res
	}

	getLeftBottom := func(i int, j int, d int) data {
		res := dp[i][j][7]
		if d < dp[i][j][7].d {
			for y := range 3 {
				res.cnt[y] -= dp[i+d][j-d][7].cnt[y]
			}
		}
		return res
	}

	getRightBottom := func(i int, j int, d int) data {
		res := dp[i][j][5]
		if d < dp[i][j][5].d {
			for y := range 3 {
				res.cnt[y] -= dp[i+d][j+d][5].cnt[y]
			}
		}
		return res
	}

	getTop := func(i int, j int, d int) data {
		res := dp[i][j][1]
		if d < dp[i][j][1].d {
			for y := range 3 {
				res.cnt[y] -= dp[i-d][j][1].cnt[y]
			}
		}
		return res
	}

	getLeft := func(i int, j int, d int) data {
		res := dp[i][j][3]
		if d < dp[i][j][3].d {
			for y := range 3 {
				res.cnt[y] -= dp[i][j-d][3].cnt[y]
			}
		}
		return res
	}

	getRight := func(i int, j int, d int) data {
		res := dp[i][j][4]
		if d < dp[i][j][4].d {
			for y := range 3 {
				res.cnt[y] -= dp[i][j+d][4].cnt[y]
			}
		}
		return res
	}

	getBottom := func(i int, j int, d int) data {
		res := dp[i][j][6]
		if d < dp[i][j][6].d {
			for y := range 3 {
				res.cnt[y] -= dp[i+d][j][6].cnt[y]
			}
		}
		return res
	}
	for i := range n {
		for j := range n {
			if a[i][j] != '0' {
				d := min(dp[i][j][0].d, dp[i][j][2].d, dp[i][j][5].d, dp[i][j][7].d)
				w := getLeftTop(i, j, d).add(getRightTop(i, j, d)).add(getLeftBottom(i, j, d)).add(getRightBottom(i, j, d))
				x := int(a[i][j] - '0')
				w.cnt[x-1] -= 3
				update(w)

				d = min(dp[i][j][1].d, dp[i][j][3].d, dp[i][j][4].d, dp[i][j][6].d)
				w = getTop(i, j, d).add(getLeft(i, j, d)).add(getRight(i, j, d)).add(getBottom(i, j, d))
				w.cnt[x-1] -= 3

				update(w)
			}
		}
	}

	if res.cnt[0]+res.cnt[1]+res.cnt[2] == 0 {
		return 0
	}

	return mul(pow(2, res.cnt[1]), pow(3, res.cnt[2]))
}
