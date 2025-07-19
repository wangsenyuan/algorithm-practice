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
	n := readNum(reader)
	t := readNNums(reader, n)
	return solve(t)
}

func solve(t []int) int {
	// 先算出它的区域
	n := len(t)
	h := 5 * n * 2
	w := 2*h + 1

	// 爆炸的方向，已经爆炸的距离
	marked := make([][][][8]bool, h)
	for i := range h {
		marked[i] = make([][][8]bool, w)
		for j := range w {
			marked[i][j] = make([][8]bool, n)
		}
	}
	cnt := make([][]int, h)
	for i := range h {
		cnt[i] = make([]int, w)
	}

	dd := [][]int{
		{-1, 0},  // 上
		{-1, 1},  // 上右
		{0, 1},   // 右
		{1, 1},   // 下右
		{1, 0},   // 下
		{1, -1},  // 下左
		{0, -1},  // 左
		{-1, -1}, // 上左
	}
	// 下一个状态，始终是 (d + 1) % 8, (d - 1) % 8

	type state struct {
		r int
		c int
		d int
		i int
	}

	var que []state
	que = append(que, state{h / 2, w / 2, 0, 0})
	marked[h/2][w/2][0][0] = true

	for len(que) > 0 {
		cur := que[0]
		que = que[1:]
		i := cur.i
		r, c, d := cur.r, cur.c, cur.d
		dist := t[i]
		// 这里是8个状态呐
		for range dist {
			cnt[r][c]++
			r += dd[d][0]
			c += dd[d][1]
		}

		if i+1 == n {
			continue
		}

		r -= dd[d][0]
		c -= dd[d][1]
		// 在这里explode
		nd := (d + 1) % 8

		nr, nc := r+dd[nd][0], c+dd[nd][1]

		if !marked[nr][nc][i+1][nd] {
			marked[nr][nc][i+1][nd] = true
			que = append(que, state{nr, nc, nd, i + 1})
		}
		nd = (d - 1 + 8) % 8
		nr, nc = r+dd[nd][0], c+dd[nd][1]
		if !marked[nr][nc][i+1][nd] {
			marked[nr][nc][i+1][nd] = true
			que = append(que, state{nr, nc, nd, i + 1})
		}
	}

	var res int

	for i := range h {
		for j := range w {
			if cnt[i][j] > 0 {
				res++
			}
		}
	}

	return res
}
