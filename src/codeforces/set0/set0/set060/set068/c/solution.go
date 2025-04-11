package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res[0], res[1])
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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
	flow := make([][]int, n*(n-1)/2)
	for i := range len(flow) {
		flow[i] = readNNums(reader, 5)
	}
	return solve(n, flow)
}

const inf = 1 << 30

func solve(n int, pipes [][]int) []int {

	flow := make([][][]int, n)
	for i := range n {
		flow[i] = make([][]int, n)
	}
	for _, cur := range pipes {
		f, s := cur[0], cur[1]
		f--
		s--
		flow[f][s] = cur[2:]
	}

	cost := -inf

	vol := make([]int, n)

	var push func(s int, nv int, sum int)
	push = func(s int, nv int, sum int) {
		if s == n-1 {
			cost = max(cost, sum)
			return
		}
		if nv == n {
			if vol[s] == 0 {
				push(s+1, s+2, sum)
			}
			return
		}
		cur := flow[s][nv]
		l, h, a := cur[0], cur[1], cur[2]
		for v := l; v <= min(h, vol[s]); v++ {
			vol[s] -= v
			vol[nv] += v

			var add int
			if v > 0 {
				add += a + v*v
			}
			push(s, nv+1, sum+add)

			vol[nv] -= v
			vol[s] += v
		}
	}

	for v := 0; v <= 5*len(pipes); v++ {
		vol[0] = v
		cost = -inf
		push(0, 1, 0)
		if cost >= 0 {
			return []int{v, cost}
		}
	}
	return []int{-1, -1}
}
