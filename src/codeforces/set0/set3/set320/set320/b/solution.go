package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	for _, x := range res {
		if x {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
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

func process(reader *bufio.Reader) []bool {
	n := readNum(reader)
	queries := make([][]int, n)
	for i := range n {
		queries[i] = readNNums(reader, 3)
	}
	return solve(n, queries)
}

func solve(n int, queries [][]int) []bool {
	var arr [][]int

	g := make([][]int, n+1)

	add := func(v int) {
		cur := arr[v]
		for u := range v {
			tmp := arr[u]
			if cur[0] < tmp[0] && tmp[0] < cur[1] || cur[0] < tmp[1] && tmp[1] < cur[1] {
				g[u] = append(g[u], v)
			}
			if tmp[0] < cur[0] && cur[0] < tmp[1] || tmp[0] < cur[1] && cur[1] < tmp[1] {
				g[v] = append(g[v], u)
			}
		}
	}

	marked := make([]bool, n)
	que := make([]int, n)
	check := func(a int, b int) bool {
		clear(marked)

		var head, tail int
		que[head] = a
		head++
		marked[a] = true

		for tail < head {
			u := que[tail]
			tail++
			for _, v := range g[u] {
				if !marked[v] {
					marked[v] = true
					que[head] = v
					head++
				}
			}
		}
		return marked[b]
	}

	var res []bool

	for _, q := range queries {
		if q[0] == 1 {
			l, r := q[1], q[2]
			arr = append(arr, []int{l, r})
			add(len(arr) - 1)
		} else {
			a, b := q[1], q[2]
			res = append(res, check(a-1, b-1))
		}
	}

	return res
}
