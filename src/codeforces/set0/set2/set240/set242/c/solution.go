package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	res := process(bufio.NewReader(os.Stdin))
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
	pos := readNNums(reader, 4)
	n := readNum(reader)
	segs := make([][]int, n)
	for i := 0; i < n; i++ {
		segs[i] = readNNums(reader, 3)
	}
	return solve(pos, segs)
}

type data struct {
	r int
	c int
}

const N = 1e9

func solve(pos []int, segments [][]int) int {
	// n := len(segments)
	dist := make(map[data]int)
	for _, seg := range segments {
		r, c1, c2 := seg[0], seg[1], seg[2]
		for c := c1; c <= c2; c++ {
			dist[data{r, c}] = inf
		}
	}
	m := len(dist)
	que := make([]data, m)
	var head, tail int
	que[head] = data{pos[0], pos[1]}
	head++
	dist[data{pos[0], pos[1]}] = 0

	for tail < head {
		cur := que[tail]
		tail++
		if dist[cur] == inf {
			break
		}
		r, c := cur.r, cur.c
		if r == pos[2] && c == pos[3] {
			return dist[cur]
		}
		for dr := -1; dr <= 1; dr++ {
			nr := r + dr
			for dc := -1; dc <= 1; dc++ {
				nc := c + dc
				if dist[data{nr, nc}] == inf {
					dist[data{nr, nc}] = dist[cur] + 1
					que[head] = data{nr, nc}
					head++
				}
			}
		}
	}
	return -1
}

func abs(x int) int {
	return max(x, -x)
}

const inf = 1 << 60
