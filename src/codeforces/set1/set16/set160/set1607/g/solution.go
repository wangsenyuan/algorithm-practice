package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for range tc {
		readString(reader)
		_, _, best, res := drive(reader)
		writer.WriteString(fmt.Sprintf("%d\n", best))
		for _, dish := range res {
			writer.WriteString(fmt.Sprintf("%d %d\n", dish[0], dish[1]))
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

func drive(reader *bufio.Reader) (m int, dishes [][]int, best int, res [][]int) {
	n, m := readTwoNums(reader)
	dishes = make([][]int, n)
	for i := range n {
		dishes[i] = readNNums(reader, 2)
	}
	best, res = solve(dishes, m)
	return
}

func solve(dishes [][]int, m int) (best int, res [][]int) {
	// diff[0]表示最后剩余更多的fish
	// diff[1]表示最后剩余更多的meat
	diff := make([]int, 2)
	n := len(dishes)
	x := make([]int, n)
	y := make([]int, n)
	for i, dish := range dishes {
		x[i] = min(m, dish[0])
		diff[0] += dish[0] - dish[1] + m - 2*x[i]
		y[i] = min(m, dish[1])
		// b[i] - y -  a[i] + (m - y)
		diff[1] += dish[1] - dish[0] + m - 2*y[i]
	}

	if diff[0] < 0 {
		//x变小的时候，可以增加diff[0]
		for i, dish := range dishes {
			// x[i] + b[i] >= m
			// x[i] >= m - b[i]
			// x[i]是有个下限的，就是m-b[i]
			// det不能把所有的x[i]都不使用
			w := min(x[i]-max(0, m-dish[1]), (-diff[0])/2)
			x[i] -= w
			diff[0] += w * 2
		}
	}
	if diff[1] < 0 {
		// 将小y可以将diff[1]变大
		for i, dish := range dishes {
			w := min(y[i]-max(0, m-dish[0]), (-diff[1])/2)
			y[i] -= w
			diff[1] += w * 2
		}
	}

	best = min(abs(diff[0]), abs(diff[1]))
	res = make([][]int, n)
	if best == abs(diff[0]) {
		for i := range n {
			res[i] = []int{x[i], m - x[i]}
		}
	} else {
		for i := range n {
			res[i] = []int{m - y[i], y[i]}
		}
	}

	return
}

func abs(a int) int {
	return max(a, -a)
}
