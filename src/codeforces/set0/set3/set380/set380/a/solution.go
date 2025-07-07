package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
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
	m := readNum(reader)
	ops := make([][]int, m)
	for i := 0; i < m; i++ {
		s, _ := reader.ReadBytes('\n')
		var op int
		pos := readInt(s, 0, &op) + 1
		if op == 1 {
			ops[i] = make([]int, 2)
		} else {
			ops[i] = make([]int, 3)
		}

		ops[i][0] = op
		for j := 1; j < len(ops[i]); j++ {
			pos = readInt(s, pos, &ops[i][j]) + 1
		}
	}
	n := readNum(reader)
	queries := readNNums(reader, n)
	return solve(ops, queries)
}

type query struct {
	id  int
	val int
	ind int
}

func solve(ops [][]int, queries []int) []int {
	var n int

	qs := make([]query, len(queries))

	ans := make([]int, len(queries))
	var max_pos int
	m := len(ops)
	for i, j := 0, 0; i < m; i++ {
		if ops[i][0] == 1 {
			// add x
			x := ops[i][1]
			n++

			for j < len(queries) && queries[j] == n {
				ans[j] = x
				j++
			}
		} else {
			l, c := ops[i][1], ops[i][2]

			n2 := n + l*c
			for j < len(queries) && queries[j] <= n2 {
				w := queries[j] - 1
				q := query{j, w, (w - n) % l}
				qs = append(qs, q)
				max_pos = max(max_pos, q.ind)
				j++
			}

			n = n2
		}
	}

	var arr []int

	for i := 0; i < m && len(arr) <= max_pos; i++ {
		if ops[i][0] == 1 {
			arr = append(arr, ops[i][1])
		} else {
			l, c := ops[i][1], ops[i][2]
			tmp := arr[:l]
			for c > 0 && len(arr) <= max_pos {
				arr = append(arr, tmp...)
				c--
			}
		}
	}

	for _, cur := range qs {
		ans[cur.id] = arr[cur.ind]
	}

	return ans
}
